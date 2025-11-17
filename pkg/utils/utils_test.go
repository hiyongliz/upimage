package utils

import (
	"os"
	"testing"
)

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

func TestSendMessageToTGBot(t *testing.T) {
	// This is a placeholder test. In a real scenario, you would mock the HTTP client.
	botToken := os.Getenv("TG_BOT_TOKEN")
	chatID := os.Getenv("TG_CHAT_ID")
	message := "Test message"

	err := SendMessageToTGBot(botToken, chatID, message)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}
