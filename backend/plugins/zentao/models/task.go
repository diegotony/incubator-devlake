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

package models

import (
	"github.com/apache/incubator-devlake/core/models/common"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

type ZentaoTaskRes struct {
	Id             int64               `json:"id"`
	Project        int64               `json:"project"`
	Parent         int64               `json:"parent"`
	Execution      int64               `json:"execution"`
	Module         int                 `json:"module"`
	Design         int                 `json:"design"`
	Story          int64               `json:"story"`
	StoryVersion   int                 `json:"storyVersion"`
	DesignVersion  int                 `json:"designVersion"`
	FromBug        int                 `json:"fromBug"`
	Feedback       int                 `json:"feedback"`
	FromIssue      int                 `json:"fromIssue"`
	Name           string              `json:"name"`
	Type           string              `json:"type"`
	Mode           string              `json:"mode"`
	Pri            int                 `json:"pri"`
	Estimate       float64             `json:"estimate"`
	Consumed       float64             `json:"consumed"`
	Left           float64             `json:"left"`
	Deadline       string              `json:"deadline"`
	Status         string              `json:"status"`
	SubStatus      string              `json:"subStatus"`
	Color          string              `json:"color"`
	Description    string              `json:"desc"`
	Version        int                 `json:"version"`
	OpenedBy       *ApiAccount         `json:"openedBy"`
	OpenedDate     *helper.Iso8601Time `json:"openedDate"`
	AssignedTo     *ApiAccount         `json:"assignedTo"`
	AssignedDate   *helper.Iso8601Time `json:"assignedDate"`
	EstStarted     string              `json:"estStarted"`
	RealStarted    *helper.Iso8601Time `json:"realStarted"`
	FinishedBy     *ApiAccount         `json:"finishedBy"`
	FinishedDate   *helper.Iso8601Time `json:"finishedDate"`
	FinishedList   string              `json:"finishedList"`
	CanceledBy     *ApiAccount         `json:"canceledBy"`
	CanceledDate   *helper.Iso8601Time `json:"canceledDate"`
	ClosedBy       *ApiAccount         `json:"closedBy"`
	ClosedDate     *helper.Iso8601Time `json:"closedDate"`
	PlanDuration   int                 `json:"planDuration"`
	RealDuration   int                 `json:"realDuration"`
	ClosedReason   string              `json:"closedReason"`
	LastEditedBy   *ApiAccount         `json:"lastEditedBy"`
	LastEditedDate *helper.Iso8601Time `json:"lastEditedDate"`
	ActivatedDate  *helper.Iso8601Time `json:"activatedDate"`
	OrderIn        int                 `json:"order"`
	Repo           int                 `json:"repo"`
	Mr             int                 `json:"mr"`
	Entry          string              `json:"entry"`
	NumOfLine      string              `json:"lines"`
	V1             string              `json:"v1"`
	V2             string              `json:"v2"`
	Vision         string              `json:"vision"`
	StoryID        int64               `json:"storyID"`
	StoryTitle     string              `json:"storyTitle"`
	Branch         interface {
	} `json:"branch"`
	LatestStoryVersion interface {
	} `json:"latestStoryVersion"`
	StoryStatus interface {
	} `json:"storyStatus"`
	AssignedToRealName string           `json:"assignedToRealName"`
	PriOrder           string           `json:"priOrder"`
	Children           []*ZentaoTaskRes `json:"children"`
	Delay              int              `json:"delay"`
	NeedConfirm        bool             `json:"needConfirm"`
	Progress           float64          `json:"progress"`
}

type ZentaoTask struct {
	common.NoPKModel
	ConnectionId       uint64  `gorm:"primaryKey;type:BIGINT  NOT NULL"`
	ID                 int64   `json:"id" gorm:"primaryKey;type:BIGINT  NOT NULL;autoIncrement:false"`
	Project            int64   `json:"project"`
	Parent             int64   `json:"parent"`
	Execution          int64   `json:"execution"`
	Module             int     `json:"module"`
	Design             int     `json:"design"`
	Story              int64   `json:"story"`
	StoryVersion       int     `json:"storyVersion"`
	DesignVersion      int     `json:"designVersion"`
	FromBug            int     `json:"fromBug"`
	Feedback           int     `json:"feedback"`
	FromIssue          int     `json:"fromIssue"`
	Name               string  `json:"name"`
	Type               string  `json:"type"`
	Mode               string  `json:"mode"`
	Pri                int     `json:"pri"`
	Estimate           float64 `json:"estimate"`
	Consumed           float64 `json:"consumed"`
	Left               float64 `json:"left" gorm:"column:db_left"`
	Deadline           string  `json:"deadline"`
	Status             string  `json:"status"`
	SubStatus          string  `json:"subStatus"`
	Color              string  `json:"color"`
	Description        string  `json:"desc"`
	Version            int     `json:"version"`
	OpenedById         int64
	OpenedByName       string
	OpenedDate         *helper.Iso8601Time `json:"openedDate"`
	AssignedToId       int64
	AssignedToName     string
	AssignedDate       *helper.Iso8601Time `json:"assignedDate"`
	EstStarted         string              `json:"estStarted"`
	RealStarted        *helper.Iso8601Time `json:"realStarted"`
	FinishedId         int64
	FinishedDate       *helper.Iso8601Time `json:"finishedDate"`
	FinishedList       string              `json:"finishedList"`
	CanceledId         int64
	CanceledDate       *helper.Iso8601Time `json:"canceledDate"`
	ClosedById         int64
	ClosedDate         *helper.Iso8601Time `json:"closedDate"`
	PlanDuration       int                 `json:"planDuration"`
	RealDuration       int                 `json:"realDuration"`
	ClosedReason       string              `json:"closedReason"`
	LastEditedId       int64
	LastEditedDate     *helper.Iso8601Time `json:"lastEditedDate"`
	ActivatedDate      *helper.Iso8601Time `json:"activatedDate"`
	OrderIn            int                 `json:"order"`
	Repo               int                 `json:"repo"`
	Mr                 int                 `json:"mr"`
	Entry              string              `json:"entry"`
	NumOfLine          string              `json:"lines"`
	V1                 string              `json:"v1"`
	V2                 string              `json:"v2"`
	Deleted            bool                `json:"deleted"`
	Vision             string              `json:"vision"`
	StoryID            int64               `json:"storyID"`
	StoryTitle         string              `json:"storyTitle"`
	Branch             int                 `json:"branch"`
	LatestStoryVersion int                 `json:"latestStoryVersion"`
	StoryStatus        string              `json:"storyStatus"`
	AssignedToRealName string              `json:"assignedToRealName"`
	PriOrder           string              `json:"priOrder"`
	NeedConfirm        bool                `json:"needConfirm"`
	Progress           float64             `json:"progress"`
	Url                string              `json:"url"`
	StdStatus          string              `json:"stdStatus" gorm:"type:varchar(20)"`
	StdType            string              `json:"stdType" gorm:"type:varchar(20)"`
}

func (ZentaoTask) TableName() string {
	return "_tool_zentao_tasks"
}
