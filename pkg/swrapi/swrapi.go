package swrapi

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hiyongliz/upimage/pkg/swrapi/models"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	swr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
	swrregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/region"
)

type SWRAPI struct {
	client *swr.SwrClient
}

func New(region string) (*SWRAPI, error) {
	ak := os.Getenv("HUAWEICLOUD_SDK_AK")
	sk := os.Getenv("HUAWEICLOUD_SDK_SK")

	auth, err := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		SafeBuild()
	if err != nil {
		return nil, fmt.Errorf("failed to create credentials: %w", err)
	}

	r, err := swrregion.SafeValueOf(region)
	if err != nil {
		return nil, fmt.Errorf("failed to get region: %w", err)
	}

	swrAuth, err := swr.SwrClientBuilder().
		WithRegion(r).
		WithCredential(auth).
		SafeBuild()
	if err != nil {
		return nil, fmt.Errorf("failed to create SWR client: %w", err)
	}

	client := swr.NewSwrClient(swrAuth)
	return &SWRAPI{
		client: client,
	}, nil
}

func (s *SWRAPI) CreateNamespace(namespace string) error {
	request := &model.CreateNamespaceRequest{}
	request.Body = &model.CreateNamespaceRequestBody{
		Namespace: namespace,
	}
	_, err := s.client.CreateNamespace(request)
	if err != nil {
		errorResponse := &models.ErrorResponse{}
		err = json.Unmarshal([]byte(fmt.Sprintf("%v", err)), errorResponse)
		if err != nil {
			return err
		}
		if errorResponse.ErrorMessage == "Namespace already exists" {
			return nil
		}
	}
	return nil
}

func (s *SWRAPI) UpdateRepo(namespace, repo string, body model.UpdateRepoRequestBody) error {
	request := &model.UpdateRepoRequest{}
	request.Namespace = namespace
	request.Repository = repo
	request.Body = &body
	_, err := s.client.UpdateRepo(request)
	if err != nil {
		return fmt.Errorf("failed to update repository: %w", err)
	}
	return nil
}
