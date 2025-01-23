package http

import (
	"io"
	"net/http"
)

func CreateRequest(method, url, headerKey, headerValue string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headerKey, headerValue)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func GetResponseBody(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
