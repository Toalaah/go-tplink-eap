package tplink

import (
	"encoding/json"
)

type responseStatus struct {
	Success bool `json:"success,omitempty"`
	Timeout bool `json:"timeout,omitempty"`
	Error   int  `json:"error,omitempty"`
}

func responseStatusFromHTTP(body []byte) (*responseStatus, error) {
	r := &responseStatus{}
	err := json.Unmarshal(body, r)
	return r, err
}
