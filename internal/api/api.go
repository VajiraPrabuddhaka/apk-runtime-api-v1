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

package api

import (
	"fmt"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/getkin/kin-openapi/openapi3"
	"k8s.io/client-go/kubernetes"
)

func GetAPI(namespace string, id string, clientSet *kubernetes.Clientset) (gen.API, error) {
	//ToDo write logic to get API from k8s API server

	//mocking the response
	return gen.API{
		Context:         "test",
		CreatedTime:     nil,
		EndpointConfig:  nil,
		LastUpdatedTime: nil,
		Name:            "TestAPI",
		Operations:      nil,
		ServiceInfo:     nil,
		Type:            nil,
		Version:         "1.0.0",
	}, nil
}

func GetAPIs(namespace string, offset int, limit int, clientSet *kubernetes.Clientset) (gen.APIList, error) {
	//ToDo write logic to get API from k8s API server

	return gen.APIList{
		Count: nil,
		List:  nil,
		Pagination: &gen.Pagination{
			Limit:    nil,
			Next:     nil,
			Offset:   nil,
			Previous: nil,
			Total:    nil,
		},
	}, nil
}

func CreateAPI(api gen.API, oasVersion string) (gen.API, error) {
	//todo: add logic to create API
	// 1. Generate swagger
	// 2. create k8s artifacts
	return gen.API{
		Context:         "",
		CreatedTime:     nil,
		EndpointConfig:  nil,
		LastUpdatedTime: nil,
		Name:            "",
		Operations:      nil,
		ServiceInfo:     nil,
		Type:            nil,
		Version:         "",
	}, nil
}

func genSwagger(api gen.API) {
	c := openapi3.NewSchema()
	fmt.Print(c)
}
