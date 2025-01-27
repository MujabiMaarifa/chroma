package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chroma",
	Short: "chroma is a command line documentation generator",
	Long: `A Fast and Flexible documentation generator built with go and powered by Mistral ai Codestral
               `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate documentation from source code",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "[ERROR]: please provide a filepath")
			os.Exit(1)
		}

		if err := generate(); err != nil {
			fmt.Fprintln(os.Stderr, "[ERROR]: please provide a filepath")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
