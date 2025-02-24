/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/jira/models"
	"github.com/apache/incubator-devlake/plugins/jira/tasks/apiv2models"
)

var _ plugin.SubTaskEntryPoint = ExtractIssues

var ExtractIssuesMeta = plugin.SubTaskMeta{
	Name:             "extractIssues",
	EntryPoint:       ExtractIssues,
	EnabledByDefault: true,
	Description:      "extract Jira issues",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_TICKET, plugin.DOMAIN_TYPE_CROSS},
}

type typeMappings struct {
	typeIdMappings         map[string]string
	stdTypeMappings        map[string]string
	standardStatusMappings map[string]models.StatusMappings
}

func ExtractIssues(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*JiraTaskData)
	db := taskCtx.GetDal()
	connectionId := data.Options.ConnectionId
	boardId := data.Options.BoardId
	logger := taskCtx.GetLogger()
	logger.Info("extract Issues, connection_id=%d, board_id=%d", connectionId, boardId)
	mappings, err := getTypeMappings(data, db)
	if err != nil {
		return err
	}
	extractor, err := api.NewApiExtractor(api.ApiExtractorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx: taskCtx,
			/*
				This struct will be JSONEncoded and stored into database along with raw data itself, to identity minimal
				set of data to be process, for example, we process JiraIssues by Board
			*/
			Params: JiraApiParams{
				ConnectionId: data.Options.ConnectionId,
				BoardId:      data.Options.BoardId,
			},
			/*
				Table store raw data
			*/
			Table: RAW_ISSUE_TABLE,
		},
		Extract: func(row *api.RawData) ([]interface{}, errors.Error) {
			return extractIssues(data, mappings, row)
		},
	})
	if err != nil {
		return err
	}
	return extractor.Execute()
}

func extractIssues(data *JiraTaskData, mappings *typeMappings, row *api.RawData) ([]interface{}, errors.Error) {
	var apiIssue apiv2models.Issue
	err := errors.Convert(json.Unmarshal(row.Data, &apiIssue))
	if err != nil {
		return nil, err
	}
	err = apiIssue.SetAllFields(row.Data)
	if err != nil {
		return nil, err
	}
	var results []interface{}
	// if the field `created` is nil, ignore it
	if apiIssue.Fields.Created == nil {
		return results, nil
	}
	sprints, issue, comments, worklogs, changelogs, changelogItems, users := apiIssue.ExtractEntities(data.Options.ConnectionId)
	for _, sprintId := range sprints {
		sprintIssue := &models.JiraSprintIssue{
			ConnectionId:     data.Options.ConnectionId,
			SprintId:         sprintId,
			IssueId:          issue.IssueId,
			IssueCreatedDate: &issue.Created,
			ResolutionDate:   issue.ResolutionDate,
		}
		results = append(results, sprintIssue)
	}
	if issue.ResolutionDate != nil {
		issue.LeadTimeMinutes = uint(issue.ResolutionDate.Unix()-issue.Created.Unix()) / 60
	}
	if data.Options.ScopeConfig != nil && data.Options.ScopeConfig.StoryPointField != "" {
		unknownStoryPoint := apiIssue.Fields.AllFields[data.Options.ScopeConfig.StoryPointField]
		switch sp := unknownStoryPoint.(type) {
		case string:
			// string, try to parse
			issue.StoryPoint, _ = strconv.ParseFloat(sp, 32)
		case nil:
		default:
			// not string, convert to float64, ignore it if failed
			issue.StoryPoint, _ = unknownStoryPoint.(float64)
		}

	}

	// code in next line will set issue.Type to issueType.Name
	issue.Type = mappings.typeIdMappings[issue.Type]
	issue.StdType = mappings.stdTypeMappings[issue.Type]
	if issue.StdType == "" {
		issue.StdType = strings.ToUpper(issue.Type)
	}
	issue.StdStatus = getStdStatus(issue.StatusKey)
	if value, ok := mappings.standardStatusMappings[issue.Type][issue.StatusKey]; ok {
		issue.StdStatus = value.StandardStatus
	}
	results = append(results, issue)
	for _, comment := range comments {
		results = append(results, comment)
	}
	for _, worklog := range worklogs {
		results = append(results, worklog)
	}
	var issueUpdated *time.Time
	// likely this issue has more changelogs to be collected
	if len(changelogs) == 100 {
		issueUpdated = nil
	} else {
		issueUpdated = &issue.Updated
	}
	for _, changelog := range changelogs {
		changelog.IssueUpdated = issueUpdated
		results = append(results, changelog)
	}
	for _, changelogItem := range changelogItems {
		results = append(results, changelogItem)
	}
	for _, user := range users {
		if user.AccountId != "" {
			results = append(results, user)
		}
	}
	results = append(results, &models.JiraBoardIssue{
		ConnectionId: data.Options.ConnectionId,
		BoardId:      data.Options.BoardId,
		IssueId:      issue.IssueId,
	})
	labels := apiIssue.Fields.Labels
	for _, v := range labels {
		issueLabel := &models.JiraIssueLabel{
			IssueId:      issue.IssueId,
			LabelName:    v,
			ConnectionId: data.Options.ConnectionId,
		}
		results = append(results, issueLabel)
	}
	return results, nil
}

func getTypeMappings(data *JiraTaskData, db dal.Dal) (*typeMappings, errors.Error) {
	typeIdMapping := make(map[string]string)
	issueTypes := make([]models.JiraIssueType, 0)
	clauses := []dal.Clause{
		dal.From(&models.JiraIssueType{}),
		dal.Where("connection_id = ?", data.Options.ConnectionId),
	}
	err := db.All(&issueTypes, clauses...)
	if err != nil {
		return nil, err
	}
	for _, issueType := range issueTypes {
		typeIdMapping[issueType.Id] = issueType.Name
	}
	stdTypeMappings := make(map[string]string)
	standardStatusMappings := make(map[string]models.StatusMappings)
	if data.Options.ScopeConfig != nil {
		for userType, stdType := range data.Options.ScopeConfig.TypeMappings {
			stdTypeMappings[userType] = strings.ToUpper(stdType.StandardType)
			standardStatusMappings[userType] = stdType.StatusMappings
		}
	}
	return &typeMappings{
		typeIdMappings:         typeIdMapping,
		stdTypeMappings:        stdTypeMappings,
		standardStatusMappings: standardStatusMappings,
	}, nil
}
