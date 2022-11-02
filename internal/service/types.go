/*
 *  Copyright (c) 2022, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package service

import (
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/cache"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
)

type ListServicesResponse struct {
	Count      int              `json:"count"`
	List       []*gen.Service   `json:"list"`
	Pagination cache.Pagination `json:"pagination"`
}

type ServiceUsageResponse struct {
	Count      int               `json:"count"`
	List       []ServiceUsageAPI `json:"list"`
	Pagination cache.Pagination  `json:"pagination"`
}

type ServiceUsageAPI struct {
	Name        string `json:"name"`
	Context     string `json:"context"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}
