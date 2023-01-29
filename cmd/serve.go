/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/shoaibashk/dokcli/api"
	"github.com/shoaibashk/dokcli/server"
	"github.com/spf13/cobra"
)

var (
	port *string
	url  *string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server",
	Long:  `start the server`,
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := server.DokcliBanner()
		fmt.Print(banner)
		// fmt.Println("Server started at : " + "http://localhost:1212")
		api.Server(*port, *url)
		fmt.Println("Listening on :  " + *port)
	},
}

func init() {

	port = serveCmd.Flags().StringP("port", "p", "1212", "Specific port for Dokcli Server to listen")
	url = serveCmd.Flags().StringP("url", "u", "https://petstore.swagger.io/v2/swagger.json", "Specific swagger file for Dokcli")

	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type serveSpec struct {
	port *string
	url  *string
}
