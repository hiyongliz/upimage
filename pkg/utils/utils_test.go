package utils

import "testing"

func TestGetNamespaceFromImage(t *testing.T) {
	tests := []struct {
		image     string
		namespace string
		err       bool
	}{
		{"myregistry/myimage:latest", "myregistry", false},
		{"myregistry/myimage", "myregistry", false},
		{"myimage:latest", "lazylibrary", false}, // Should use default namespace
		{"myimage", "lazylibrary", false},        // Should use default namespace
	}

	for _, tt := range tests {
		t.Run(tt.image, func(t *testing.T) {
			ns, err := GetNamespaceFromImage(tt.image)
			if (err != nil) != tt.err {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if ns != tt.namespace {
				t.Errorf("expected namespace: %s, got: %s", tt.namespace, ns)
			}
		})
	}
}

func TestGetRepoFromImage(t *testing.T) {
	tests := []struct {
		image string
		repo  string
		err   bool
	}{
		{"myregistry/myimage:latest", "myimage", false},
		{"myregistry/myimage", "myimage", false},
		{"myimage:latest", "myimage", false},
		{"myimage", "myimage", false},
		{"invalid/image:tag:extra", "", true}, // Invalid format
	}

	for _, tt := range tests {
		t.Run(tt.image, func(t *testing.T) {
			repo, err := GetRepoFromImage(tt.image)
			if (err != nil) != tt.err {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if repo != tt.repo {
				t.Errorf("expected repo: %s, got: %s", tt.repo, repo)
			}
		})
	}
}
