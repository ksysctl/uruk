package actions

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
	"github.com/spf13/viper"
)

// Define new storage engine
func (r Engine) storage() {
	str := "%s:%s@tcp(%s:%s)/%s?parseTime=True"
	source := fmt.Sprintf(str,
		viper.GetString("database_username"),
		viper.GetString("database_password"),
		viper.GetString("database_host"),
		viper.GetString("database_port"),
		viper.GetString("database_name"),
	)

	db, err := gorm.Open("mysql", source)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	r.router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
}
