package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appName = "uruk"
	cfgFile string
)

// RootCmd is the base command
var RootCmd = &cobra.Command{
	Use:   appName,
	Short: "Go app.",
	Long:  `App written in Go.`,
}

// Execute runs command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		fmt.Sprintf("config file (default is .env)"))

	_ = viper.BindPFlag(
		"config",
		RootCmd.PersistentFlags().Lookup("config"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Use config file from default path
		path, _ := os.Getwd()
		viper.SetConfigFile(fmt.Sprintf("%s/.env", path))
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
