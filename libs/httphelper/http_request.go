package httphelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//NewJSONRequest func
func NewJSONRequest(method string, url string, payload interface{}, headers []map[string]string) (map[string]interface{}, error) {
	var respBody map[string]interface{}
	var err error

	bodyData, _ := json.Marshal(payload)
	body := bytes.NewReader(bodyData)

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return respBody, err
	}

	req.Header.Set("Content-Type", "application/json")

	if len(headers) > 0 {
		for _, header := range headers {
			for key, value := range header {
				req.Header.Set(key, value)
			}

		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		return respBody, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return respBody, err
	}

	if resp.StatusCode != 200 {
		msg := respBody["message"]
		if msg == "" {
			msg = resp.Status
		}
		return respBody, fmt.Errorf("error %s", msg)
	}

	return respBody, err
}
