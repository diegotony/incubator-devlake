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

package archived

import (
	"github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

type TeambitionTaskTag struct {
	ConnectionId   uint64 `gorm:"primaryKey;type:BIGINT"`
	Id             string `gorm:"primaryKey;type:varchar(100)"`
	CreatorId      string `gorm:"type:varchar(100)"`
	ProjectId      string `gorm:"type:varchar(100)"`
	OrganizationId string `gorm:"type:varchar(100)"`
	Name           string `gorm:"type:varchar(100)"`
	Color          string `gorm:"type:varchar(100)"`
	IsArchived     bool
	Created        *api.Iso8601Time
	Updated        *api.Iso8601Time

	archived.NoPKModel
}

func (TeambitionTaskTag) TableName() string {
	return "_tool_teambition_task_tags"
}
