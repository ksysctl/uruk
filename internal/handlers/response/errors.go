package response

import (
	"fmt"
)

// ErrorResponse represents an response error
type ErrorResponse struct {
	msg string
}

// Error satisfies the error interface
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s", e.msg)
}
