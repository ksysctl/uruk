package response

// Supported response content types
const (
	typeJSON = iota
	typeHTML
	typeXML
	typeYAML
)

const (
	// TypeJSON represents JSON Content-Type
	TypeJSON = "json"
	// TypeHTML represents HTML Content-Type
	TypeHTML = "html"
	// TypeXML represents XML Content-Type
	TypeXML = "xml"
	// TypeYAML represents YAML Content-Type
	TypeYAML = "yaml"
)

var types = map[string]int{
	TypeJSON: typeJSON,
	TypeHTML: typeHTML,
	TypeXML:  typeXML,
	TypeYAML: typeYAML,
}

// SetType maps supported Content-Type
func SetType(value string) string {
	_, f := types[value]
	if !f {
		return TypeJSON
	}

	return value
}
