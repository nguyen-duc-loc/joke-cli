package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "joke-cli",
	Short: "Joke is a command line tool provides programming jokes",
	Long:  "A command line tool provides programming jokes built with love by nguyenducloc in Go\nhttps://github.com/nguyen-duc-loc/joke-cli",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
