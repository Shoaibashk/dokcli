/*
Copyright © 2023 Shoaib Shaikh <shoaibashk.2000@gmail.com>
*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/shoaibashk/dokcli/model"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create Config file for dokcli server",
	Long: `Create Config file for dokcli server 
	
	usage: 
		$ dokcli init
		$ dokcli serve
	`,
	Run: func(cmd *cobra.Command, args []string) {

		yamlData, err := yaml.Marshal(model.Defaults)

		if err != nil {
			fmt.Printf("Error while Marshaling. %v", err)
		}
		filename := "configfile.yaml"
		err = ioutil.WriteFile(filename, yamlData, 0644)
		if err != nil {
			panic("Unable to write data into the file")

		}
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
