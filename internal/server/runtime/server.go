package runtime

import (
	"encoding/json"
	"fmt"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/service"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/httproute/gateway/clientset/v1alpha2"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type Server struct {
	ClientSetK8s      *kubernetes.Clientset
	ClientSetV1alpha1 *v1alpha2.HttpRouteV1Alpha1Client
}

func (r2 Server) GetAllAPIs(w http.ResponseWriter, r *http.Request, params gen.GetAllAPIsParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) CreateAPI(w http.ResponseWriter, r *http.Request, params gen.CreateAPIParams) {
	//TODO implement me
	panic("implement me")
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

func (r2 Server) ImportServiceFromCatalog(w http.ResponseWriter, r *http.Request, params gen.ImportServiceFromCatalogParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ValidateAPI(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ValidateAsyncAPISpecification(w http.ResponseWriter, r *http.Request, params gen.ValidateAsyncAPISpecificationParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ValidateGraphQLSchema(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) ValidateOpenAPIDefinition(w http.ResponseWriter, r *http.Request, params gen.ValidateOpenAPIDefinitionParams) {
	//TODO implement me
	panic("implement me")
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

func (r2 Server) GetAPIGraphQLSchema(w http.ResponseWriter, r *http.Request, apiId string, params gen.GetAPIGraphQLSchemaParams) {
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

func (r2 Server) GetAllGateways(w http.ResponseWriter, r *http.Request, params gen.GetAllGatewaysParams) {
	//TODO implement me
	panic("implement me")
}

func (r2 Server) CreateGateway(w http.ResponseWriter, r *http.Request, params gen.CreateGatewayParams) {
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

func (r2 Server) Search(w http.ResponseWriter, r *http.Request, params gen.SearchParams) {
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

func (r2 Server) GetServiceById(w http.ResponseWriter, r *http.Request, serviceId string) {
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

func (r2 Server) GetServiceUsage(w http.ResponseWriter, r *http.Request, serviceId string) {
	s := service.GetServiceUsage("default", serviceId, r2.ClientSetK8s)
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
