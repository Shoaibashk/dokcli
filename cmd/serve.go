/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mbndr/figlet4go"
	"github.com/shoaibashk/dokcli/api"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server",
	Long:  `start the server`,
	Run: func(cmd *cobra.Command, args []string) {

		ascii := figlet4go.NewAsciiRender()

		bannerOptions := figlet4go.NewRenderOptions()
		bannerOptions.FontName = "larry3d"
		bannerOptions.FontColor = []figlet4go.Color{
			// Colors can be given by default ansi color codes...
			figlet4go.ColorGreen,
			figlet4go.ColorRed,
			figlet4go.ColorCyan,

			// ...or by an hex string...
			// figlet4go.NewTrueColorFromHexString("885DBA"),
			// ...or by an TrueColor object with rgb values
			// figlet4go.TrueColor{136, 93, 186},

		}

		// figlet4go.TrueColor{255, 198, 211}

		renderStr, _ := ascii.RenderOpts("Dok Cli", bannerOptions)
		fmt.Print(renderStr)
		// fmt.Println("Server started at : " + "http://localhost:1212")
		api.Server()
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
