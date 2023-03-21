/*
Copyright © 2023 Shoaib Shaikh <shoaibashk.2000@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "(untracked)"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the Version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
