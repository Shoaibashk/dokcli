/*
Copyright © 2023 Shoaib Shaikh <shoaibashk.2000@gmail.com>
*/
package cmd

import (
	"fmt"

	// "github.com/shoaibashk/dokcli/api"
	// "github.com/shoaibashk/dokcli/server"

	"github.com/shoaibashk/dokcli/api"
	"github.com/shoaibashk/dokcli/model"
	"github.com/shoaibashk/dokcli/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port       *string
	url        *string
	configFile *string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server",
	Long:  `start the server`,
	Run: func(cmd *cobra.Command, args []string) {

		isConfigFileSet := cmd.Flags().ShorthandLookup("f")
		isConfigFileSet = cmd.Flags().Lookup("file")

		vi := viper.New()

		for k, v := range model.Defaults {
			vi.SetDefault(k, v)
		}

		vi.SetConfigFile(*configFile)

		vi.ReadInConfig()

		// fmt.Println(vi.GetString("name"))
		// fmt.Println(vi.GetStringMap("tags"))
		// fmt.Printf("%v", vi.GetStringMapString("spec-url")["petstore"])

		banner, _ := server.DokcliBanner()
		fmt.Print(banner)

		// fmt.Println(isConfigFileSet.Changed)

		if isConfigFileSet.Changed {
			var url = vi.GetStringMapString("spec-url")["petstore"]
			var port = vi.GetString("port")
			fmt.Println("Server started at : " + "http://localhost:" + port)
			api.Server(port, url)
		} else {
			fmt.Println("Server started at : " + "http://localhost:" + *port)
			api.Server(*port, *url)
		}

	},
}

func init() {

	configFile = serveCmd.Flags().StringP(
		"file",
		"f",
		"",
		"configuartion file for dokcli to serve (use:  dokcli serve -f .filechange.yaml)")

	port = serveCmd.Flags().StringP(
		"port",
		"p",
		"1212",
		"Specific port for Dokcli Server to listen")

	url = serveCmd.Flags().StringP(
		"url",
		"u",
		"https://petstore.swagger.io/v2/swagger.json",
		"Specific swagger file for Dokcli")

	rootCmd.AddCommand(serveCmd)

}
