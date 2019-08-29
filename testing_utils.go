package endpass

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

const (
	charsetUTF8                    = "charset=UTF-8"
	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + charsetUTF8
)

func createServer(
	responseCode int,
	responseContentType string,
	data interface{},
) *httptest.Server {
	var (
		responseBody []byte
		err          error
	)
	switch typedData := data.(type) {
	case []byte:
		responseBody = typedData
	default:
		responseBody, err = json.Marshal(data)
		if err != nil {
			panic(err)
		}
	}
	requestHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", MIMEApplicationJSONCharsetUTF8)
		w.WriteHeader(responseCode)
		_, _ = w.Write(responseBody)
	})
	return httptest.NewServer(requestHandler)
}
