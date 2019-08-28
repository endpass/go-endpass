package endpass

import (
	"errors"
	"fmt"
)

var (
	ErrNoAccessToken   = errors.New("Before doing requests you must pass authorization and get access token. See \"AuthCodeURL\" and \"Exchange\" methods")
	ErrFileNotUploaded = errors.New("file not uploaded yet")
)

// ErrorHTTPResponse return this error if StatusCode from API not equal 200.
type ErrorHTTPResponse struct {
	Status       string
	ResponseBody string
}

func NewErrorHTTPResponse(status, responseBody string) *ErrorHTTPResponse {
	return &ErrorHTTPResponse{
		Status:       status,
		ResponseBody: responseBody,
	}
}

func (e *ErrorHTTPResponse) Error() string {
	return fmt.Sprintf("HTTPResponseError(status: %s; responseBody: %s)", e.Status, e.ResponseBody)
}
