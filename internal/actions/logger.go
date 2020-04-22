package actions

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Define logger format
func (r Engine) logger() {
	r.router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %s | %d - %s %s %s \n",
			viper.GetString("app_name"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.TimeStamp.Format(time.RFC850),
			param.ClientIP,
			param.ErrorMessage,
		)
	}))
}
