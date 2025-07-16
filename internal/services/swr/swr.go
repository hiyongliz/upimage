package swr

import (
	"github.com/hiyongliz/upimage/pkg/swrapi"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
)

type SWRService struct {
	swrapi *swrapi.SWRAPI
}

func New(swrapi *swrapi.SWRAPI) *SWRService {
	return &SWRService{
		swrapi: swrapi,
	}
}

func (s *SWRService) CreateNamespace(namespace string) error {
	return s.swrapi.CreateNamespace(namespace)
}

func (s *SWRService) UpdateRepoPublic(namespace, repo string) error {
	return s.swrapi.UpdateRepo(namespace, repo, model.UpdateRepoRequestBody{
		IsPublic: true,
	})
}
