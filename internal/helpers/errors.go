package helpers

import (
	"fmt"
)

// ErrorHelper represents an helper error
type ErrorHelper struct {
	msg string
}

// Error satisfies the error interface
func (e *ErrorHelper) Error() string {
	return fmt.Sprintf("%s", e.msg)
}
