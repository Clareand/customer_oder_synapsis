package httphelper

import (
	"encoding/json"
	"net/http"
)

//HTTPResponse struct
type HTTPResponse struct {
	StatusCode string      `json:"status_code"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
}

//HTTPError struct
type HTTPError struct {
	StatusCode   int         `json:"status_code"`
	Details      interface{} `json:"details,omitempty"`
	ResourceID   string      `json:"resource_id,omitempty"`
	ResourceType string      `json:"resource_type,omitempty"`
	Message      interface{} `json:"message"`
}

// ResponseJSON makes the response with payload as json format
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

//NewHTTPError func
func NewHTTPError(code int, details interface{}) HTTPError {
	return HTTPError{
		StatusCode: code,
		Details:    details,
	}
}
