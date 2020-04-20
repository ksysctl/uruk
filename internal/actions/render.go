package actions

import (
	"html/template"

	"github.com/ksysctl/uruk/internal/helpers"
	"github.com/michelloworld/ez-gin-template"
)

// Define new template engine
func (r Engine) render() {
	render := eztemplate.New()
	render.Debug = true
	render.TemplatesDir = "web/templates/views/"

	// Attach custom functions
	render.TemplateFuncMap = template.FuncMap{
		"formatAsDate": helpers.FormatAsDate,
	}

	r.router.HTMLRender = render.Init()
}
