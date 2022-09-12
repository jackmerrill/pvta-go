package pvtago

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type NetworkClient struct {
	baseURL string
}

func (c *NetworkClient) DoRequest(method string, path string, body map[string]interface{}) ([]byte, error) {
	httpClient := &http.Client{}
	var jsonBody []byte = nil

	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		jsonBody = j
	}

	req, err := http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func (c *NetworkClient) Parse(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}

func (c *NetworkClient) Get(path string, out interface{}) error {
	data, err := c.DoRequest("GET", path, nil)
	if err != nil {
		return err
	}
	return c.Parse(data, out)
}
