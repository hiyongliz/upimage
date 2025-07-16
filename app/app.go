package app

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hiyongliz/upimage/internal/services/swr"
	"github.com/hiyongliz/upimage/pkg/swrapi"
	"github.com/hiyongliz/upimage/pkg/utils"
)

type Up struct {
	swrService *swr.SWRService
	region     string
}

func NewUp(region string) (*Up, error) {
	SWRAPI, err := swrapi.New(region)
	if err != nil {
		return nil, err
	}
	SWRService := swr.New(SWRAPI)
	return &Up{
		swrService: SWRService,
		region:     region,
	}, nil
}

func (u *Up) Execute(image string) error {
	repo, err := utils.GetRepoFromImage(image)
	if err != nil {
		return fmt.Errorf("failed to parse repository name from image %q: %w", image, err)
	}
	namespace, err := utils.GetNamespaceFromImage(image)
	if err != nil {
		return fmt.Errorf("failed to parse namespace from image %q: %w", image, err)
	}
	tag, err := utils.GetTagFromImage(image)
	if err != nil {
		return fmt.Errorf("failed to parse tag from image %q: %w", image, err)
	}

	fmt.Printf("Processing image: %s\n", image)
	fmt.Printf("  Namespace: %s\n", namespace)
	fmt.Printf("  Repository: %s\n", repo)
	fmt.Printf("  Tag: %s\n", tag)
	fmt.Printf("  Target region: %s\n", u.region)

	// Create SWR namespace if it doesn't exist
	fmt.Printf("Creating namespace %q if it doesn't exist...\n", namespace)
	if err := u.swrService.CreateNamespace(namespace); err != nil {
		return fmt.Errorf("failed to create namespace %q: %w", namespace, err)
	}

	// Download the image
	fmt.Printf("Pulling image %q...\n", image)
	cmd := exec.Command("docker", "pull", image)
	if err := runCmd(cmd); err != nil {
		return fmt.Errorf("failed to pull image %q: %w", image, err)
	}

	// Tag the image
	newImage := fmt.Sprintf("swr.%s.myhuaweicloud.com/%s/%s:%s", u.region, namespace, repo, tag)
	fmt.Printf("Tagging image as %q...\n", newImage)
	cmd = exec.Command("docker", "tag", image, newImage)
	if err := runCmd(cmd); err != nil {
		return fmt.Errorf("failed to tag image %q as %q: %w", image, newImage, err)
	}

	// Push the image to SWR
	fmt.Printf("Pushing image %q to SWR...\n", newImage)
	cmd = exec.Command("docker", "push", newImage)
	if err := runCmd(cmd); err != nil {
		return fmt.Errorf("failed to push image %q: %w", newImage, err)
	}

	// Register the image to SWR
	fmt.Printf("Setting repository %q/%q as public...\n", namespace, repo)
	if err := u.swrService.UpdateRepoPublic(namespace, repo); err != nil {
		return fmt.Errorf("failed to update repository %q/%q to public: %w", namespace, repo, err)
	}

	fmt.Printf("✅ Successfully synced image %q to %q\n", image, newImage)
	return nil
}

func runCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
