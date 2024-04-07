package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Snipe",
	Short: "Snipe is a CLI tool for managing code snippets.",
	Long: `Snipe is a command-line interface (CLI) application designed to help developers 
manage and organize their code snippets effectively.

With snippet-manager, you can add, list, search, edit, and delete code snippets with ease.
It supports multi-line content and provides a flexible way to store and retrieve your valuable code snippets.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


