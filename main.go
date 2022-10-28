package main

import (
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/runtime"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/pkg/k8s/client"
	"net/http"
)

func main() {

	s := &runtime.Server{
		ClientSetK8s:      client.GetOutClusterClientSetKubernetes(),
		ClientSetV1alpha1: client.GetOutClusterClientSetV1alpha1(),
	}
	h := gen.Handler(s)

	http.ListenAndServe(":3000", h)
}
