package utils

import (
	"fmt"
	"strings"
)

func GetNamespaceFromImage(image string) (string, error) {
	// Split the image by '/' to get the namespace
	parts := strings.Split(image, "/")
	if len(parts) < 2 {
		return "lazylibrary", nil
	}

	// The namespace is the first part
	namespace := parts[0]
	return namespace, nil
}

func GetRepoFromImage(image string) (string, error) {
	// Split the image by '/' to get the repository
	parts := strings.Split(image, "/")
	if len(parts) < 1 {
		return "", fmt.Errorf("invalid image format: %s", image)
	}
	if strings.Count(image, ":") > 1 {
		return "", fmt.Errorf("invalid image format: %s", image)
	}

	// The repository is the last part (after the last ':')
	repoWithTag := parts[len(parts)-1]
	repoParts := strings.Split(repoWithTag, ":")
	return repoParts[0], nil
}

func GetTagFromImage(image string) (string, error) {
	// Split the image by ':' to get the tag
	if strings.Count(image, ":") == 0 {
		return "latest", nil
	}
	parts := strings.Split(image, ":")
	if len(parts) > 2 {
		return "", fmt.Errorf("invalid image format: %s", image)
	}
	return parts[1], nil
}
