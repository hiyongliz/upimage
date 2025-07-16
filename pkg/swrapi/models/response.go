package models

type ErrorResponse struct {
	StatusCode                  int    `json:"status_code"`
	RequestId                   string `json:"request_id"`
	ErrorCode                   string `json:"error_code"`
	ErrorMessage                string `json:"error_message"`
	EncodedAuthorizationMessage string `json:"encoded_authorization_message,omitempty"`
}
