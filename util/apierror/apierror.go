package apierror

import (
	"fmt"
)

type ApiError struct {
	ErrorCode    string `json:"error_code"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorMessage)
}
