package runtime

import (
	"encoding/json"
	"fmt"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/cache"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/service"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/httproute/gateway/clientset/v1alpha2"
	"k8s.io/client-go/kubernetes"
	"log"
	"net/http"
)

type Server struct {
	ClientSetK8s      *kubernetes.Clientset
	ClientSetV1alpha1 *v1alpha2.HttpRouteV1Alpha1Client
	ApiCache          *cache.APILocalCache
	ServiceCache      *cache.ServiceLocalCache
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
	s := ""
	a, err := r2.ApiCache.APIsSearch(params.Offset, params.Limit, &s, &s, (*string)(params.SortBy), (*string)(params.SortOrder))
	b, err := json.Marshal(a)
	if err != nil {
		log.Printf("Error marshalling apis : %v", err)
	}
	w.Write(b)
}

func (r2 Server) CreateAPI(w http.ResponseWriter, r *http.Request) {
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

func (r2 Server) ValidateAPI(w http.ResponseWriter, r *http.Request) {
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
	s, err := r2.ServiceCache.ServicesSearch(params.Offset, params.Limit, params.Name, params.Namespace,
		(*string)(params.SortBy), (*string)(params.SortOrder))
	b, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling Services : %v\n", err)
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
