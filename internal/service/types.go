package service

import (
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/common"
	"github.com/VajiraPrabuddhaka/apk-runtime-api-v1/internal/server/gen"
)

type ListServicesResponse struct {
	Count      int               `json:"count"`
	List       []*gen.Service    `json:"list"`
	Pagination common.Pagination `json:"pagination"`
}

type ServiceUsageResponse struct {
	Count      int               `json:"count"`
	List       []ServiceUsageAPI `json:"list"`
	Pagination common.Pagination `json:"pagination"`
}

type ServiceUsageAPI struct {
	Name        string `json:"name"`
	Context     string `json:"context"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}
