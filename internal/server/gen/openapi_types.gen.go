// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package gen

import (
	"encoding/json"
	"fmt"
)

// Defines values for APIType.
const (
	ASYNC   APIType = "ASYNC"
	GRAPHQL APIType = "GRAPHQL"
	HTTP    APIType = "HTTP"
	SSE     APIType = "SSE"
	WEBHOOK APIType = "WEBHOOK"
	WEBSUB  APIType = "WEBSUB"
	WS      APIType = "WS"
)

// API defines model for API.
type API struct {
	Context     string  `json:"context"`
	CreatedTime *string `json:"createdTime,omitempty"`

	// Endpoint configuration of the API. This can be used to provide different types of endpoints including Simple REST Endpoints, Loadbalanced and Failover.
	//
	// `Simple REST Endpoint`
	//   {
	//     "endpoint_type": "http",
	//     "sandbox_endpoints":       {
	//        "url": "https://pizzashack-service:8080/am/sample/pizzashack/v3/api/"
	//     },
	//     "production_endpoints":       {
	//        "url": "https://pizzashack-service:8080/am/sample/pizzashack/v3/api/"
	//     }
	//   }
	EndpointConfig    *map[string]interface{} `json:"endpointConfig,omitempty"`
	LastUpdatedTime   *string            `json:"lastUpdatedTime,omitempty"`
	MediationPolicies *[]MediationPolicy `json:"mediationPolicies,omitempty"`
	Name              string             `json:"name"`
	Operations        *[]APIOperations   `json:"operations,omitempty"`
	ServiceInfo       *APIServiceInfo    `json:"serviceInfo,omitempty"`

	// The api creation type to be used. Accepted values are HTTP, WS, GRAPHQL, WEBSUB, SSE, WEBHOOK, ASYNC
	Type    *APIType `json:"type,omitempty"`
	Version string   `json:"version"`
}

// The api creation type to be used. Accepted values are HTTP, WS, GRAPHQL, WEBSUB, SSE, WEBHOOK, ASYNC
type APIType string

// APIInfo defines model for APIInfo.
type APIInfo struct {
	Context     *string `json:"context,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	UpdatedTime *string `json:"updatedTime,omitempty"`
	Version     *string `json:"version,omitempty"`
}

// APIList defines model for APIList.
type APIList struct {
	// Number of APIs returned.
	Count      *int        `json:"count,omitempty"`
	List       *[]APIInfo  `json:"list,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// APIOperationPolicies defines model for APIOperationPolicies.
type APIOperationPolicies struct {
	Fault    *[]OperationPolicy `json:"fault,omitempty"`
	Request  *[]OperationPolicy `json:"request,omitempty"`
	Response *[]OperationPolicy `json:"response,omitempty"`
}

// APIOperations defines model for APIOperations.
type APIOperations struct {
	Id                *string               `json:"id,omitempty"`
	OperationPolicies *APIOperationPolicies `json:"operationPolicies,omitempty"`
	PayloadSchema     *string               `json:"payloadSchema,omitempty"`
	Scopes            *[]string             `json:"scopes,omitempty"`
	Target            *string               `json:"target,omitempty"`
	UriMapping        *string               `json:"uriMapping,omitempty"`
	Verb              *string               `json:"verb,omitempty"`
}

// APIServiceInfo defines model for API_serviceInfo.
type APIServiceInfo struct {
	Name *string `json:"name,omitempty"`
}

// AsyncAPISpecificationValidationResponse defines model for AsyncAPISpecificationValidationResponse.
type AsyncAPISpecificationValidationResponse struct {
	// AsyncAPI specification content
	Content *string `json:"content,omitempty"`

	// If there are more than one error list them out. For example, list out validation error by each field.
	Errors *[]ErrorListItem `json:"errors,omitempty"`

	// API definition information
	Info *AsyncAPISpecificationValidationResponseInfo `json:"info,omitempty"`

	// This attribute declares whether this definition is valid or not.
	IsValid bool `json:"isValid"`
}

// API definition information
type AsyncAPISpecificationValidationResponseInfo struct {
	AsyncAPIVersion *string `json:"asyncAPIVersion,omitempty"`

	// contains available transports for an async API
	AsyncTransportProtocols *[]string `json:"asyncTransportProtocols,omitempty"`
	Context                 *string   `json:"context,omitempty"`
	Description             *string   `json:"description,omitempty"`

	// contains host/servers specified in the AsyncAPI file/URL
	Endpoints     *[]string `json:"endpoints,omitempty"`
	GatewayVendor *string   `json:"gatewayVendor,omitempty"`
	Name          *string   `json:"name,omitempty"`
	Protocol      *string   `json:"protocol,omitempty"`
	Version       *string   `json:"version,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code int64 `json:"code"`

	// A detail description about the error message.
	Description *string `json:"description,omitempty"`

	// If there are more than one error list them out.
	// For example, list out validation errors by each field.
	Error *[]ErrorListItem `json:"error,omitempty"`

	// Error message.
	Message string `json:"message"`

	// Preferably an url with more details about the error.
	MoreInfo *string `json:"moreInfo,omitempty"`
}

// ErrorListItem defines model for ErrorListItem.
type ErrorListItem struct {
	Code string `json:"code"`

	// A detail description about the error message.
	Description *string `json:"description,omitempty"`

	// Description about individual errors occurred
	Message string `json:"message"`
}

// Gateway defines model for Gateway.
type Gateway struct {
	// Name of the Gateway
	Name string `json:"name"`
}

// GatewayList defines model for GatewayList.
type GatewayList struct {
	List       *[]Gateway  `json:"list,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// GraphQLSchema defines model for GraphQLSchema.
type GraphQLSchema struct {
	Name             string  `json:"name"`
	SchemaDefinition *string `json:"schemaDefinition,omitempty"`
}

// GraphQLValidationResponse defines model for GraphQLValidationResponse.
type GraphQLValidationResponse struct {
	ErrorMessage string `json:"errorMessage"`

	// This attribute declares whether this definition is valid or not.
	IsValid bool `json:"isValid"`
}

// MediationPolicy defines model for MediationPolicy.
type MediationPolicy struct {
	Name   string  `json:"name"`
	Shared *bool   `json:"shared,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// MediationPolicyList defines model for MediationPolicyList.
type MediationPolicyList struct {
	List       *[]MediationPolicy `json:"list,omitempty"`
	Pagination *Pagination        `json:"pagination,omitempty"`
}

// OpenAPIDefinitionValidationResponse defines model for OpenAPIDefinitionValidationResponse.
type OpenAPIDefinitionValidationResponse struct {
	// OpenAPI definition content.
	Content *string `json:"content,omitempty"`

	// If there are more than one error list them out.
	// For example, list out validation errors by each field.
	Errors *[]ErrorListItem `json:"errors,omitempty"`

	// API definition information
	Info *OpenAPIDefinitionValidationResponseInfo `json:"info,omitempty"`

	// This attribute declares whether this definition is valid or not.
	IsValid bool `json:"isValid"`
}

// API definition information
type OpenAPIDefinitionValidationResponseInfo struct {
	Context     *string `json:"context,omitempty"`
	Description *string `json:"description,omitempty"`

	// contains host/servers specified in the OpenAPI file/URL
	Endpoints      *[]string `json:"endpoints,omitempty"`
	Name           *string   `json:"name,omitempty"`
	OpenAPIVersion *string   `json:"openAPIVersion,omitempty"`
	Version        *string   `json:"version,omitempty"`
}

// OperationPolicy defines model for OperationPolicy.
type OperationPolicy struct {
	Parameters    *OperationPolicy_Parameters `json:"parameters,omitempty"`
	PolicyId      *string                     `json:"policyId,omitempty"`
	PolicyName    string                      `json:"policyName"`
	PolicyVersion *string                     `json:"policyVersion,omitempty"`
}

// OperationPolicy_Parameters defines model for OperationPolicy.Parameters.
type OperationPolicy_Parameters struct {
	AdditionalProperties map[string]map[string]interface{} `json:"-"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	Limit *int `json:"limit,omitempty"`

	// Link to the next subset of resources qualified.
	// Empty if no more resources are to be returned.
	Next   *string `json:"next,omitempty"`
	Offset *int    `json:"offset,omitempty"`

	// Link to the previous subset of resources qualified.
	// Empty if current subset is the first subset returned.
	Previous *string `json:"previous,omitempty"`
	Total    *int    `json:"total,omitempty"`
}

// PortMapping defines model for PortMapping.
type PortMapping struct {
	Name       string  `json:"name"`
	Port       int32 `json:"port"`
	Protocol   *string  `json:"protocol"`
	Targetport int32 `json:"targetport"`
}

// SearchResultList defines model for SearchResultList.
type SearchResultList struct {
	// Number of results returned.
	Count      *int                      `json:"count,omitempty"`
	List       *[]map[string]interface{} `json:"list,omitempty"`
	Pagination *Pagination               `json:"pagination,omitempty"`
}

// Service defines model for Service.
type Service struct {
	Name        string        `json:"name"`
	Namespace   string        `json:"namespace"`
	Portmapping []*PortMapping `json:"portmapping"`
	Type        string        `json:"type"`
}

// ServiceList defines model for ServiceList.
type ServiceList struct {
	List       *[]Service  `json:"list,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// ApiIdAsyncapiBody defines model for apiId_asyncapi_body.
type ApiIdAsyncapiBody struct {
	// AsyncAPI definition of the API
	ApiDefinition *string `json:"apiDefinition,omitempty"`

	// AsyncAPI definition as a file
	File *string `json:"file,omitempty"`

	// AsyncAPI definition URL of the API
	Url *string `json:"url,omitempty"`
}

// ApiIdGraphqlschemaBody defines model for apiId_graphqlschema_body.
type ApiIdGraphqlschemaBody struct {
	// schema definition of the GraphQL API
	SchemaDefinition string `json:"schemaDefinition"`
}

// ApiIdSwaggerBody defines model for apiId_swagger_body.
type ApiIdSwaggerBody struct {
	// Swagger definition of the API
	ApiDefinition *string `json:"apiDefinition,omitempty"`

	// Swagger definitio as a file
	File *string `json:"file,omitempty"`

	// Swagger definition URL of the API
	Url *string `json:"url,omitempty"`
}

// ApisImportBody defines model for apis_import_body.
type ApisImportBody struct {
	// Zip archive consisting on exported API configuration
	File string `json:"file"`
}

// ApisImportgraphqlschemaBody defines model for apis_importgraphqlschema_body.
type ApisImportgraphqlschemaBody struct {
	// Additional attributes specified as a stringified JSON with API's schema
	AdditionalProperties *string `json:"additionalProperties,omitempty"`

	// Definition to uploads a file
	File *string `json:"file,omitempty"`

	// Definition type to upload
	Type *string `json:"type,omitempty"`
}

// ApisImportopenapiBody defines model for apis_importopenapi_body.
type ApisImportopenapiBody struct {
	// Additional attributes specified as a stringified JSON with API's schema
	AdditionalProperties *string `json:"additionalProperties,omitempty"`

	// Definition to upload as a file
	File *string `json:"file,omitempty"`

	// Inline content of the OpenAPI definition
	InlineAPIDefinition *string `json:"inlineAPIDefinition,omitempty"`

	// Definition url
	Url *string `json:"url,omitempty"`
}

// ApisValidateasyncapiBody defines model for apis_validateasyncapi_body.
type ApisValidateasyncapiBody struct {
	// AsyncAPI definition as a file
	File *string `json:"file,omitempty"`

	// AsyncAPI definition url
	Url *string `json:"url,omitempty"`
}

// ApisValidategraphqlschemaBody defines model for apis_validategraphqlschema_body.
type ApisValidategraphqlschemaBody struct {
	// Definition to upload as a file
	File string `json:"file"`
}

// ApisValidateopenapiBody defines model for apis_validateopenapi_body.
type ApisValidateopenapiBody struct {
	// OpenAPI definition as a file
	File *string `json:"file,omitempty"`

	// Inline content of the OpenAPI definition
	InlineAPIDefinition *string `json:"inlineAPIDefinition,omitempty"`

	// OpenAPI definition url
	Url *string `json:"url,omitempty"`
}

// GetAllAPIsParams defines parameters for GetAllAPIs.
type GetAllAPIsParams struct {
	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Criteria for sorting.
	SortBy *GetAllAPIsParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Order of sorting(ascending/descending).
	SortOrder *GetAllAPIsParamsSortOrder `form:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Query     *string                    `form:"query,omitempty" json:"query,omitempty"`

	// Media types acceptable for the response. Default is application/json.
	Accept *string `json:"Accept,omitempty"`
}

// GetAllAPIsParamsSortBy defines parameters for GetAllAPIs.
type GetAllAPIsParamsSortBy string

// GetAllAPIsParamsSortOrder defines parameters for GetAllAPIs.
type GetAllAPIsParamsSortOrder string

// CreateAPIJSONBody defines parameters for CreateAPI.
type CreateAPIJSONBody = API

// CreateAPIParams defines parameters for CreateAPI.
type CreateAPIParams struct {
	// Open api version
	OpenAPIVersion *CreateAPIParamsOpenAPIVersion `form:"openAPIVersion,omitempty" json:"openAPIVersion,omitempty"`
}

// CreateAPIParamsOpenAPIVersion defines parameters for CreateAPI.
type CreateAPIParamsOpenAPIVersion string

// ExportAPIParams defines parameters for ExportAPI.
type ExportAPIParams struct {
	// Name of the API
	ApiId *string `form:"apiId,omitempty" json:"apiId,omitempty"`

	// API Name
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Version of the API
	Version *string `form:"version,omitempty" json:"version,omitempty"`

	// Format of output documents. Can be YAML or JSON.
	Format *ExportAPIParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// ExportAPIParamsFormat defines parameters for ExportAPI.
type ExportAPIParamsFormat string

// ImportAPIParams defines parameters for ImportAPI.
type ImportAPIParams struct {
	// Whether to update the API or not. This is used when updating already existing APIs
	Overwrite *bool `form:"overwrite,omitempty" json:"overwrite,omitempty"`
}

// ImportServiceFromCatalogJSONBody defines parameters for ImportServiceFromCatalog.
type ImportServiceFromCatalogJSONBody = API

// ImportServiceFromCatalogParams defines parameters for ImportServiceFromCatalog.
type ImportServiceFromCatalogParams struct {
	// ID of service that should be imported from Service Catalog
	ServiceKey string `form:"serviceKey" json:"serviceKey"`
}

// ValidateAsyncAPISpecificationParams defines parameters for ValidateAsyncAPISpecification.
type ValidateAsyncAPISpecificationParams struct {
	// Specify whether to return the full content of the AsyncAPI specification in the response. This is only applicable when using url based validation
	ReturnContent *bool `form:"returnContent,omitempty" json:"returnContent,omitempty"`
}

// ValidateOpenAPIDefinitionParams defines parameters for ValidateOpenAPIDefinition.
type ValidateOpenAPIDefinitionParams struct {
	// Specify whether to return the full content of the OpenAPI definition in the response. This is only
	// applicable when using url based validation
	ReturnContent *bool `form:"returnContent,omitempty" json:"returnContent,omitempty"`
}

// UpdateAPIJSONBody defines parameters for UpdateAPI.
type UpdateAPIJSONBody = API

// GetAPIGraphQLSchemaParams defines parameters for GetAPIGraphQLSchema.
type GetAPIGraphQLSchemaParams struct {
	// Media types acceptable for the response. Default is application/json.
	Accept *string `json:"Accept,omitempty"`
}

// GetAllGatewaysParams defines parameters for GetAllGateways.
type GetAllGatewaysParams struct {
	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Criteria for sorting.
	SortBy *GetAllGatewaysParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Order of sorting(ascending/descending).
	SortOrder *GetAllGatewaysParamsSortOrder `form:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Query     *string                        `form:"query,omitempty" json:"query,omitempty"`

	// Media types acceptable for the response. Default is application/json.
	Accept *string `json:"Accept,omitempty"`
}

// GetAllGatewaysParamsSortBy defines parameters for GetAllGateways.
type GetAllGatewaysParamsSortBy string

// GetAllGatewaysParamsSortOrder defines parameters for GetAllGateways.
type GetAllGatewaysParamsSortOrder string

// CreateGatewayJSONBody defines parameters for CreateGateway.
type CreateGatewayJSONBody = Gateway

// CreateGatewayParams defines parameters for CreateGateway.
type CreateGatewayParams struct {
	// Open api version
	OpenAPIVersion *CreateGatewayParamsOpenAPIVersion `form:"openAPIVersion,omitempty" json:"openAPIVersion,omitempty"`
}

// CreateGatewayParamsOpenAPIVersion defines parameters for CreateGateway.
type CreateGatewayParamsOpenAPIVersion string

// UpdateGatewayJSONBody defines parameters for UpdateGateway.
type UpdateGatewayJSONBody = Gateway

// GetAllPoliciesParams defines parameters for GetAllPolicies.
type GetAllPoliciesParams struct {
	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Criteria for sorting.
	SortBy *GetAllPoliciesParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Order of sorting(ascending/descending).
	SortOrder *GetAllPoliciesParamsSortOrder `form:"sortOrder,omitempty" json:"sortOrder,omitempty"`
	Query     *string                        `form:"query,omitempty" json:"query,omitempty"`

	// Media types acceptable for the response. Default is application/json.
	Accept *string `json:"Accept,omitempty"`
}

// GetAllPoliciesParamsSortBy defines parameters for GetAllPolicies.
type GetAllPoliciesParamsSortBy string

// GetAllPoliciesParamsSortOrder defines parameters for GetAllPolicies.
type GetAllPoliciesParamsSortOrder string

// CreatePolicyJSONBody defines parameters for CreatePolicy.
type CreatePolicyJSONBody = MediationPolicy

// SearchParams defines parameters for Search.
type SearchParams struct {
	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// **Search**.
	//
	// You can search by proving a keyword.
	Query *string `form:"query,omitempty" json:"query,omitempty"`
}

// SearchServicesParams defines parameters for SearchServices.
type SearchServicesParams struct {
	// Filter services by the name of the service
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Criteria for sorting.
	SortBy *SearchServicesParamsSortBy `form:"sortBy,omitempty" json:"sortBy,omitempty"`

	// Order of sorting(ascending/descending).
	SortOrder *SearchServicesParamsSortOrder `form:"sortOrder,omitempty" json:"sortOrder,omitempty"`

	// Maximum size of resource array to return.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// Starting point within the complete list of items qualified.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`
}

// SearchServicesParamsSortBy defines parameters for SearchServices.
type SearchServicesParamsSortBy string

// SearchServicesParamsSortOrder defines parameters for SearchServices.
type SearchServicesParamsSortOrder string

// CreateAPIJSONRequestBody defines body for CreateAPI for application/json ContentType.
type CreateAPIJSONRequestBody = CreateAPIJSONBody

// ImportServiceFromCatalogJSONRequestBody defines body for ImportServiceFromCatalog for application/json ContentType.
type ImportServiceFromCatalogJSONRequestBody = ImportServiceFromCatalogJSONBody

// UpdateAPIJSONRequestBody defines body for UpdateAPI for application/json ContentType.
type UpdateAPIJSONRequestBody = UpdateAPIJSONBody

// CreateGatewayJSONRequestBody defines body for CreateGateway for application/json ContentType.
type CreateGatewayJSONRequestBody = CreateGatewayJSONBody

// UpdateGatewayJSONRequestBody defines body for UpdateGateway for application/json ContentType.
type UpdateGatewayJSONRequestBody = UpdateGatewayJSONBody

// CreatePolicyJSONRequestBody defines body for CreatePolicy for application/json ContentType.
type CreatePolicyJSONRequestBody = CreatePolicyJSONBody

// Getter for additional properties for OperationPolicy_Parameters. Returns the specified
// element and whether it was found
func (a OperationPolicy_Parameters) Get(fieldName string) (value map[string]interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OperationPolicy_Parameters
func (a *OperationPolicy_Parameters) Set(fieldName string, value map[string]interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OperationPolicy_Parameters to handle AdditionalProperties
func (a *OperationPolicy_Parameters) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal map[string]interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OperationPolicy_Parameters to handle AdditionalProperties
func (a OperationPolicy_Parameters) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
