package actions

import (
	"html/template"

	"github.com/ksysctl/uruk/internal/helpers"

	"github.com/gin-gonic/gin"
	"github.com/michelloworld/ez-gin-template"
	"github.com/spf13/viper"
)

// Define new template engine
func (r Engine) render() {
	render := eztemplate.New()
	render.TemplatesDir = "web/templates/views/"

	if viper.GetString("server_mode") == gin.DebugMode {
		render.Debug = true
	}

	// Attach custom functions
	render.TemplateFuncMap = template.FuncMap{
		"formatAsDate": helpers.FormatAsDate,
	}

	r.router.HTMLRender = render.Init()
}
