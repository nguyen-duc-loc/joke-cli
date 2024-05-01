package cmd

import (
	"errors"
	"fmt"

	"github.com/nguyen-duc-loc/joke-cli/utils"
	"github.com/spf13/cobra"
)

var jokeNum uint8

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate proramming jokes",
	Long:  "Generate one or more proramming jokes. For example:\njoke-cli generate -n 5",
	RunE: func(cmd *cobra.Command, args []string) error {
		if jokeNum < 1 || jokeNum > 10 {
			return errors.New("number of jokes must be range from 1 - 10")
		}

		request := utils.NewRequest("en", []string{}, "json", jokeNum)
		jokes, err := utils.GenerateJokes(request)
		if err != nil {
			return err
		}

		for _, joke := range jokes {
			fmt.Printf("[+] %s\n\n", joke.Content)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().Uint8VarP(&jokeNum, "number", "n", 1, "Number of joke (1 - 10)")
}
