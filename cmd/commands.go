package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/chachacollins/chroma/cfg"
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

var markdownCmd = &cobra.Command{
	Use:   "md",
	Short: "Generate markdown documentation and save it to a file",
	Long: `Takes a file as input and generates markdown documentation and saves it to the specified output file.
Warning: If a file already exists its contents will be replaced.
          `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Error: please provide a <code path> : <save file>")
			os.Exit(1)
		}
		fileName := args[0]
		writeFilePath := args[1]
		split := strings.Split(writeFilePath, ".")
		if split[1] != "md" {
			fmt.Fprintln(os.Stderr, "Error: please provide a filename with a markdown extension")
			os.Exit(1)
		}
		file := readFile(fileName)
		generateMd(writeFilePath, file)
	},
}

var inlineCmd = &cobra.Command{
	Use:   "il",
	Short: "Generate inline documentation and save it to a file",
	Long: `Takes a file as input and generates inline documentation(comments) and saves it to the  same  file.
Warning: The file contents will be replaced by the ai function.
          `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Error: please provide a <code path> ")
			os.Exit(1)
		}
		fileName := args[0]
		file := readFile(fileName)
		inlineComm(fileName, file)
	},
}

var starCmd = &cobra.Command{
	Use:   "star",
	Short: "Generates astro documentation",
	Long: `Creates a docs directory with the generated documentation of the astro framework.
Warning: The file contents will be replaced by the ai function.
          `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Error: please provide a <code path> or serve command ")
			os.Exit(1)
		}
		fileName := args[0]
		file := readFile(fileName)
		getDocs()
		starLight(fileName, file)
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize proper env variable",
	Long:  `Init`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg.Init()
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves generated astro docs",
	Long: `Creates a docs directory with the generated documentation of the astro framework.
Warning: The file contents will be replaced by the ai function.
          `,
	Run: func(cmd *cobra.Command, args []string) {
		serveDocs()
	},
}

func init() {
	rootCmd.AddCommand(markdownCmd)
	rootCmd.AddCommand(inlineCmd)
	rootCmd.AddCommand(starCmd)
	rootCmd.AddCommand(initCmd)
	starCmd.AddCommand(serveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
