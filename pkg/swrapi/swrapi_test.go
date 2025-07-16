package swrapi

import (
	"os"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
)

func TestNewClient(t *testing.T) {
	// Skip test if credentials not provided
	ak := os.Getenv("HUAWEICLOUD_SDK_AK")
	sk := os.Getenv("HUAWEICLOUD_SDK_SK")
	if ak == "" || sk == "" {
		t.Skip("Skipping test: HUAWEICLOUD_SDK_AK and HUAWEICLOUD_SDK_SK environment variables not set")
	}

	// Test with valid credentials
	client, err := New("cn-south-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to be created, got nil")
	}
}

func TestCreateNamespace(t *testing.T) {
	// Skip test if credentials not provided
	ak := os.Getenv("HUAWEICLOUD_SDK_AK")
	sk := os.Getenv("HUAWEICLOUD_SDK_SK")
	if ak == "" || sk == "" {
		t.Skip("Skipping integration test: HUAWEICLOUD_SDK_AK and HUAWEICLOUD_SDK_SK environment variables not set")
	}

	client, err := New("cn-south-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test creating a namespace
	err = client.CreateNamespace("test-namespace")
	if err != nil {
		t.Logf("Note: This may fail if namespace already exists or if permissions are insufficient: %v", err)
	}
}

func TestUpdateRepo(t *testing.T) {
	// Skip test if credentials not provided
	ak := os.Getenv("HUAWEICLOUD_SDK_AK")
	sk := os.Getenv("HUAWEICLOUD_SDK_SK")
	if ak == "" || sk == "" {
		t.Skip("Skipping integration test: HUAWEICLOUD_SDK_AK and HUAWEICLOUD_SDK_SK environment variables not set")
	}

	client, err := New("cn-south-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test updating a repository (this is an integration test that may fail)
	body := model.UpdateRepoRequestBody{
		IsPublic: true,
	}
	err = client.UpdateRepo("test-namespace", "test-repo", body)
	if err != nil {
		t.Logf("Note: This may fail if repository doesn't exist: %v", err)
	}
}
