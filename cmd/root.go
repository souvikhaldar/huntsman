package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "huntsman",
	Short: "Swissknife for security related stuff",
	Long:  "A blazing fast swissknife for doing versatile security work like concurrent port scanning, tcp proxy, etc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome home, huntsman!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
