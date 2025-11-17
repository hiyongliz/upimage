package utils

import (
	"fmt"
	"net/http"
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

func SendMessageToTGBot(botToken, chatID, message string) error {
	// Placeholder function for sending message to Telegram Bot
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	payload := fmt.Sprintf("chat_id=%s&text=%s", chatID, message)

	// Here you would typically use an HTTP client to send the POST request
	// For brevity, we'll just print the payload
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}
	return nil
}
