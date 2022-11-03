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
	"context"
	"encoding/json"
	"fmt"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/getkin/kin-openapi/openapi3"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

func ValidateOpenAPIDefinition(oasDefinition []byte) gen.APIDefinitionValidationResponse {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, _ := loader.LoadFromData(oasDefinition)
	err := doc.Validate(context.Background())
	if err != nil {
		fmt.Printf("Error", err)
		return gen.APIDefinitionValidationResponse{IsValid: false}
	}
	return gen.APIDefinitionValidationResponse{IsValid: true}
}

func GenerateAPIDefinition(api gen.API) openapi3.T {
	var paths openapi3.Paths = make(map[string]*openapi3.PathItem)
	pathItem := openapi3.PathItem{}
	for _, operations := range *api.Operations {
		convertResourceToSwaggerOperation(&pathItem, operations)
	}
	data, _ := json.Marshal(pathItem)
	println(data)
	paths["/*"] = &pathItem

	doc := openapi3.T{
		OpenAPI: "3.0.1",
		Info: &openapi3.Info{
			Title:   api.Name,
			Version: api.Version,
		},
		Servers: openapi3.Servers{
			{
				URL: "http://example.com/api/",
			},
		},
		Paths: paths,
	}

	return doc
}

func GetDefaultUriTemplates(apiType string) []gen.APIOperations {
	supportedMethods := [5]string{"get", "post", "put", "delete", "patch"}
	defaultUriMapping := "/*"
	var apiOperations []gen.APIOperations
	for _, method := range supportedMethods {
		methodStr := method
		apiOperations = append(apiOperations, gen.APIOperations{
			Target: &defaultUriMapping,
			Verb:   &methodStr,
		})
	}
	return apiOperations
}

func convertResourceToSwaggerOperation(pathItem *openapi3.PathItem, operation gen.APIOperations) {
	defaultOKDescription := "OK"
	defaultResponse := &openapi3.Operation{Responses: openapi3.Responses{
		"200": &openapi3.ResponseRef{Value: &openapi3.Response{Description: &defaultOKDescription}},
	}}
	switch *operation.Verb {
	case "get":
		pathItem.Get = defaultResponse
	case "post":
		pathItem.Post = defaultResponse
	case "put":
		pathItem.Put = defaultResponse
	case "delete":
		pathItem.Delete = defaultResponse
	case "patch":
		pathItem.Patch = defaultResponse
	default:
		pathItem.Get = defaultResponse
	}
}

func CreateSwaggerDefinitionConfigMap(apiName string, definition []byte, clientSet *kubernetes.Clientset) {
	swaggerMap := make(map[string]string)
	swaggerMap["swagger.yaml"] = string(definition)
	cm := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "swagger-definition-" + strings.ToLower(apiName),
			Namespace: "default",
		},
		Data: swaggerMap,
	}
	//TODO get proper namespace
	_, err := clientSet.CoreV1().ConfigMaps("default").Create(context.Background(), &cm, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Unable to store API definition for API:"+apiName+", ", err)
	}
}
