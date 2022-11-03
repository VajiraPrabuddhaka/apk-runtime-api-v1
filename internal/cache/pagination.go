/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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

package cache

import (
	"errors"
	"sort"
)

type Pagination struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Total    int    `json:"total"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func GenResultWithPagination(list interface{}, item string, offset *int, limit *int, sortBy *string,
	sortOrder string) (interface{}, int, *Pagination, error) {
	switch item {
	case "API":
		var apis []API
		i, _ := getSliceForPagination(list, item, offset, limit)
		apis = i.([]API)
		pagination := &Pagination{
			Offset: *offset,
			Limit:  *limit,
			Total:  len(apis),
			//Todo: check the values for Next & Previous and set properly
			Next:     "",
			Previous: "",
		}
		switch *sortBy {
		case "apiName":
			if sortOrder == "desc" {
				sort.Slice(apis, func(i, j int) bool {
					return apis[i].Name < apis[j].Name
				})
				return apis, len(apis), pagination, nil
			} else {
				sort.Slice(apis, func(i, j int) bool {
					return apis[i].Name > apis[j].Name
				})
				return apis, len(apis), pagination, nil
			}
		case "version":
			if sortOrder == "desc" {
				sort.Slice(apis, func(i, j int) bool {
					return apis[i].Version < apis[j].Version
				})
				return apis, len(apis), pagination, nil
			} else {
				sort.Slice(apis, func(i, j int) bool {
					return apis[i].Version > apis[j].Version
				})
				return apis, len(apis), pagination, nil
			}
		default:
			if sortOrder == "desc" {
				sort.Slice(apis, func(i, j int) bool {
					return apis[i].CreatedTime < apis[j].CreatedTime
				})
				return apis, len(apis), pagination, nil
			} else {
				sort.Slice(apis, func(i, j int) bool {
					return apis[i].CreatedTime > apis[j].CreatedTime
				})
				return apis, len(apis), pagination, nil
			}
		}
	case "Service":
		var services []Service
		i, _ := getSliceForPagination(list, item, offset, limit)
		services = i.([]Service)
		pagination := &Pagination{
			Offset:   *offset,
			Limit:    *limit,
			Total:    len(services),
			Next:     "",
			Previous: "",
		}
		switch *sortBy {
		case "Name":
			if sortOrder == "desc" {
				sort.Slice(services, func(i, j int) bool {
					return services[i].Name < services[j].Name
				})
				return services, len(services), pagination, nil
			} else {
				sort.Slice(services, func(i, j int) bool {
					return services[i].Name > services[j].Name
				})
				return services, len(services), pagination, nil
			}
		case "Type":
			if sortOrder == "desc" {
				sort.Slice(services, func(i, j int) bool {
					return services[i].Type < services[j].Type
				})
				return services, len(services), pagination, nil
			} else {
				sort.Slice(services, func(i, j int) bool {
					return services[i].Type > services[j].Type
				})
				return services, len(services), pagination, nil
			}
		default:
			if sortOrder == "desc" {
				sort.Slice(services, func(i, j int) bool {
					return services[i].Name < services[j].Name
				})
				return services, len(services), pagination, nil
			} else {
				sort.Slice(services, func(i, j int) bool {
					return services[i].Name > services[j].Name
				})
				return services, len(services), pagination, nil
			}
		}
	}
	return nil, 0, nil, errors.New("pagination error")
}

func getSliceForPagination(list interface{}, item string, offset *int, limit *int) (interface{}, error) {
	switch item {
	case "Service":
		var services []Service
		services = list.([]Service)
		if len(services) >= *offset && len(services) >= *limit {
			return services[*offset:*limit], nil
		} else if len(services) >= *offset && len(services) < *limit {
			return services[*offset:], nil
		} else {
			return []Service{}, nil
		}
	case "API":
		var apis []API
		apis = list.([]API)
		if len(apis) >= *offset && len(apis) >= *limit {
			return apis[*offset:*limit], nil
		} else if len(apis) >= *offset && len(apis) < *limit {
			return apis[*offset:], nil
		} else {
			return []API{}, nil
		}
	default:
		return nil, errors.New("invalid type to generate pagination")
	}
}
