package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ksysctl/uruk/internal/actions"
)

// Run runs web server
func Run(cmd *cobra.Command, args []string) {
	gin.SetMode(
		fmt.Sprintf("%s", viper.GetString("server_mode")),
	)

	handler := actions.Engine{}.Init()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt32("server_port")),
		Handler: handler,
	}

	server.ListenAndServe()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the web app.",
	Long:  `starts the web app written in Go.`,
	Run:   Run,
}

func init() {
	serveCmd.PersistentFlags().Int(
		"server_port",
		9000,
		"Port to run Application server on")

	serveCmd.PersistentFlags().String(
		"server_mode",
		"release",
		"Run mode")

	serveCmd.PersistentFlags().String(
		"app_name",
		appName,
		"Application name")

	_ = viper.BindPFlag(
		"server_port",
		serveCmd.PersistentFlags().Lookup("server_port"))

	_ = viper.BindPFlag(
		"server_mode",
		serveCmd.PersistentFlags().Lookup("server_mode"))

	_ = viper.BindPFlag(
		"app_name",
		serveCmd.PersistentFlags().Lookup("app_name"))

	viper.SetDefault(
		"license",
		"apache")

	RootCmd.AddCommand(serveCmd)
}
