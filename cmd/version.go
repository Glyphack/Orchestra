package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Pring version number",
	Long:  "C'mon you know what a version is :)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Orchestra version 0.1.0")
	},
}
