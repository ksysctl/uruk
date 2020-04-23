package actions

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
	"github.com/spf13/viper"

	"github.com/ksysctl/uruk/internal/repository"
)

// Define new storage engine
func (r Engine) storage() {
	dsn := "%s:%s@tcp(%s:%s)/%s?parseTime=True"
	source := fmt.Sprintf(dsn,
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

	err = db.DB().Ping()
	if err != nil {
		fmt.Println(err)
		panic("Failed to ping to database!")
	}

	repos := repository.New(gin.Mode(), db)

	r.router.Use(func(c *gin.Context) {
		c.Set("repos", repos)

		c.Next()
	})
}
