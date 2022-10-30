/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/arjprd/crypt-service/driver"
	"github.com/arjprd/crypt-service/service"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start crypt-service",
	Long:  `command starts service on port and config specified`,
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Flag to get port to run the service on
	// serveCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "specify the port on which the service has to listen over")
	// viper.BindPFlag("")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve(cmd *cobra.Command, args []string) {
	config := driver.NewConfig(configPath)
	service, err := service.NewService(config)
	if err != nil {
		log.Fatal(err)
	}
	service.Start()
}
