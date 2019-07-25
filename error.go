package endpass

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrNoAccessToken = errors.New("Before doing requests you must pass authorization and get access token. See \"AuthCodeURL\" and \"Exchange\" methods")
)

// ErrorHTTPResponse return this error if StatusCode from API not equal 200.
type ErrorHTTPResponse struct {
	HTTPResponse *http.Response
}

func NewErrorHTTPResponse(resp *http.Response) *ErrorHTTPResponse {
	return &ErrorHTTPResponse{
		HTTPResponse: resp,
	}
}

func (e *ErrorHTTPResponse) Error() string {
	bodyBytes, err := ioutil.ReadAll(e.HTTPResponse.Body)
	if err != nil {
		return fmt.Sprintf(
			"HTTPResponseError(Status: %s; Body: %s)",
			e.HTTPResponse.Status,
			err.Error(),
		)
	}
	return fmt.Sprintf(
		"HTTPResponseError(Status: %s; Body: %s)",
		e.HTTPResponse.Status,
		string(bodyBytes),
	)
}
