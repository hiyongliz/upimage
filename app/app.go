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
	options    *UpOptions
}

type UpOptions struct {
	Region          string
	Namespace       string
	Public          bool
	CreateNamespace bool
	Registry        string // Added registry field
}

func NewUp(options UpOptions) (*Up, error) {
	if options.Registry != "swr" {
		return &Up{
			options: &options,
		}, nil
	} else {
		SWRAPI, err := swrapi.New(options.Region)
		if err != nil {
			return nil, err
		}
		SWRService := swr.New(SWRAPI)
		return &Up{
			swrService: SWRService,
			options:    &options,
		}, nil
	}
}

func (u *Up) Execute(image string) error {
	repo, err := utils.GetRepoFromImage(image)
	if err != nil {
		return fmt.Errorf("failed to parse repository name from image %q: %w", image, err)
	}
	// namespace, err := utils.GetNamespaceFromImage(image)
	// if err != nil {
	// 	return fmt.Errorf("failed to parse namespace from image %q: %w", image, err)
	// }
	tag, err := utils.GetTagFromImage(image)
	if err != nil {
		return fmt.Errorf("failed to parse tag from image %q: %w", image, err)
	}

	fmt.Printf("Processing image: %s\n", image)
	fmt.Printf("  Namespace: %s\n", u.options.Namespace)
	fmt.Printf("  Repository: %s\n", repo)
	fmt.Printf("  Tag: %s\n", tag)
	fmt.Printf("  Target region: %s\n", u.options.Region)

	if u.options.CreateNamespace {
		// Create SWR namespace if it doesn't exist
		fmt.Printf("Creating namespace %q if it doesn't exist...\n", u.options.Namespace)
		if err := u.swrService.CreateNamespace(u.options.Namespace); err != nil {
			return fmt.Errorf("failed to create namespace %q: %w", u.options.Namespace, err)
		}
	}

	// Download the image
	fmt.Printf("Pulling image %q...\n", image)
	cmd := exec.Command("docker", "pull", image)
	if err := runCmd(cmd); err != nil {
		return fmt.Errorf("failed to pull image %q: %w", image, err)
	}

	newImage := ""
	// Tag the image
	switch u.options.Registry {
	case "swr":
		newImage = fmt.Sprintf("swr.%s.myhuaweicloud.com/%s/%s:%s", u.options.Region, u.options.Namespace, repo, tag)
	case "acr":
		newImage = fmt.Sprintf("registry.%s.aliyuncs.com/%s/%s:%s", u.options.Region, u.options.Namespace, repo, tag)
	case "tcr":
		newImage = fmt.Sprintf("ccr.ccs.tencentyun.com/%s/%s:%s", u.options.Namespace, repo, tag)
	default:
		return fmt.Errorf("unsupported registry: %s", u.options.Registry)
	}
	fmt.Printf("Tagging image as %q...\n", newImage)
	cmd = exec.Command("docker", "tag", image, newImage)
	if err := runCmd(cmd); err != nil {
		return fmt.Errorf("failed to tag image %q as %q: %w", image, newImage, err)
	}

	// Push the image to SWR
	fmt.Printf("Pushing image %q to %s...\n", newImage, u.options.Registry)
	cmd = exec.Command("docker", "push", newImage)
	if err := runCmd(cmd); err != nil {
		return fmt.Errorf("failed to push image %q: %w", newImage, err)
	}

	if u.options.Public {
		fmt.Printf("Setting repository %q/%q as public...\n", u.options.Namespace, repo)
		switch u.options.Registry {
		case "swr":
			if err := u.swrService.UpdateRepoPublic(u.options.Namespace, repo); err != nil {
				return fmt.Errorf("failed to update repository %q/%q to public: %w", u.options.Namespace, repo, err)
			}
		default:
			fmt.Printf("The registry %s does not support public repositories.\n", u.options.Registry)
		}
	}

	fmt.Printf("✅ Successfully synced image %q to %q\n", image, newImage)
	return nil
}

func runCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
