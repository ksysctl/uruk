package response

// Supported response content types
const (
	typeJSON = iota
	typeHTML
)

const (
	// TypeJSON represents JSON Content-Type
	TypeJSON = "json"
	// TypeHTML represents HTML Content-Type
	TypeHTML = "html"
)

var types = map[string]int{
	TypeJSON: typeJSON,
	TypeHTML: typeHTML,
}

// SetType maps supported Content-Type
func SetType(value string) string {
	_, f := types[value]
	if !f {
		return TypeJSON
	}

	return value
}
