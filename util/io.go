package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	httpClient     = http.DefaultClient
	ErrBodySupport = errors.New("body only support string, []byte, map[any]any.")
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func HttpDo(method, url string, content any, headers map[string]string) ([]byte, int, error) {
	var body io.Reader
	switch content.(type) {
	case string:
		body = strings.NewReader(content.(string))
	case []byte:
		body = bytes.NewReader(content.([]byte))
	case map[any]any:
		jsonBytes, err := json.Marshal(content)
		if err != nil {
			return nil, 0, err
		}
		body = bytes.NewReader(jsonBytes)
	default:
		return nil, 0, errors.Join(ErrBodySupport, fmt.Errorf("but body type: %T", content))
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
