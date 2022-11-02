package runtime

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/service"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/httproute/gateway/clientset/v1alpha2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type Server struct {
	ClientSetK8s      *kubernetes.Clientset
	ClientSetV1alpha1 *v1alpha2.HttpRouteV1Alpha1Client
}

func (r2 Server) ImportAPIDefinition(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ImportService(w http.ResponseWriter, r *http.Request, params gen.ImportServiceParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetAPIDefinition(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) UpdateAPIDefinition(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetAllAPIs(w http.ResponseWriter, r *http.Request, params gen.GetAllAPIsParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) CreateAPI(w http.ResponseWriter, r *http.Request) {
	// Declare a new API struct.
	var api gen.API

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&api)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if api.Operations == nil {
		defaultOperations := getDefaultUriTemplates("HTTP")
		api.Operations = &defaultOperations
	}

	openApiDef := generateAPIDefinition(api)

	var data []byte
	if data, err = yaml.Marshal(openApiDef); err != nil {
		panic(err)
	}

	createSwaggerDefintionConfigMap(api.Name, data, r2.ClientSetK8s)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (r2 Server) ExportAPI(w http.ResponseWriter, r *http.Request, params gen.ExportAPIParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ImportAPI(w http.ResponseWriter, r *http.Request, params gen.ImportAPIParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ImportGraphQLSchema(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ImportOpenAPIDefinition(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ValidateAPI(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ValidateAPIDefinition(w http.ResponseWriter, r *http.Request, params gen.ValidateAPIDefinitionParams) {
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f, _, _ := r.FormFile("file")
	oasdefinition, _ := ioutil.ReadAll(f)
	b, err := json.Marshal(validateOpenAPIDefinition(oasdefinition))
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (r2 Server) DeleteAPI(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetAPI(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) UpdateAPI(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetApisApiIdAsyncapi(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) PutApisApiIdAsyncapi(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) UpdateAPIGraphQLSchema(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetAPISwagger(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) UpdateAPISwagger(w http.ResponseWriter, r *http.Request, apiId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) DeleteGateway(w http.ResponseWriter, r *http.Request, gatewayId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetGateway(w http.ResponseWriter, r *http.Request, gatewayId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) UpdateGateway(w http.ResponseWriter, r *http.Request, gatewayId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetAllPolicies(w http.ResponseWriter, r *http.Request, params gen.GetAllPoliciesParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) CreatePolicy(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) DeletePolicy(w http.ResponseWriter, r *http.Request, mediationPolicyId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) GetPolicy(w http.ResponseWriter, r *http.Request, mediationPolicyId string) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) SearchServices(w http.ResponseWriter, r *http.Request, params gen.SearchServicesParams) {
	//ToDO handle errors
	s := service.GetServices("default", 0, 4, r2.ClientSetK8s)
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (r2 Server) GetServiceById(w http.ResponseWriter, r *http.Request, serviceId string, params gen.GetServiceByIdParams) {
	//ToDO handle errors
	s := service.GetService("default", serviceId, r2.ClientSetK8s)
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (r2 Server) GetServiceUsage(w http.ResponseWriter, r *http.Request, serviceId string, params gen.GetServiceUsageParams) {
	s := service.GetServiceUsage("default", serviceId, r2.ClientSetK8s)
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func validateOpenAPIDefinition(oasDefinition []byte) gen.APIDefinitionValidationResponse {
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

func generateAPIDefinition(api gen.API) openapi3.T {
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

func getDefaultUriTemplates(apiType string) []gen.APIOperations {
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

func createSwaggerDefintionConfigMap(apiName string, definition []byte, clientSet *kubernetes.Clientset) {
	swaggerMap := make(map[string]string)
	swaggerMap["swagger.yaml"] = string(definition)
	cm := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "swagger-definition-" + apiName,
			Namespace: "my-namespace",
		},
		Data: swaggerMap,
	}
	clientSet.CoreV1().ConfigMaps("my-namespace").Create(context.Background(), &cm, metav1.CreateOptions{})
}
