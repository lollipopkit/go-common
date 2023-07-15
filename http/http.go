package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var (
	httpClient = http.DefaultClient
)

func HttpDo(method, url string, content any, headers map[string]string) ([]byte, int, error) {
	var body io.Reader
	switch content.(type) {
	case string:
		body = strings.NewReader(content.(string))
	case []byte:
		body = bytes.NewReader(content.([]byte))
	case nil:
		// do nothing
	default:
		jsonBytes, err := json.Marshal(content)
		if err != nil {
			return nil, 0, err
		}
		body = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}
