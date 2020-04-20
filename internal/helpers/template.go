package helpers

import (
	"fmt"
	"time"
)

// FormatAsDate renders time as string
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%-02d-%02d", year, month, day)
}
