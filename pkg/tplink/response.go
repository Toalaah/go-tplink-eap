package tplink

import (
	"fmt"
)

type responseStatus struct {
	Success bool `json:"success,omitempty"`
	Timeout bool `json:"timeout,omitempty"`
	Error   int  `json:"error,omitempty"`
}

func (r *responseStatus) Ok() bool {
	return r.Success && !r.Timeout && r.Error == 0
}

func (r *responseStatus) String() string {
	return fmt.Sprintf(`{ "success": %t, "timeout": %t, "error": %d }`, r.Success, r.Timeout, r.Error)
}
