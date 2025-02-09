package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func MakeDir(dirName string) {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not create a dir: %s", err)
	}
}

func WriteFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		PrintError("Error: could not create file")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		PrintError("Error: could not write to file")
	}
	err = writer.Flush()

	if err != nil {
		PrintError("could not flush buffer")
	}
}

func PrintSuccess(msg string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")).
		Bold(true)
	fmt.Println(style.Render(msg))
}

func PrintError(msg string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	fmt.Println(style.Render(msg))
}

func PrintLog(msg string) {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true)

	fmt.Println(style.Render(msg))
}

func ExtractCodeBlocks(md string) []string {
	re := regexp.MustCompile("```(?s)(.*?)```")
	matches := re.FindAllStringSubmatch(md, -1)
	var results []string
	for _, match := range matches {
		if len(match) > 1 {
			content := strings.TrimSpace(match[1])
			if idx := strings.Index(content, "\n"); idx != -1 {
				content = strings.TrimSpace(content[idx:])
			}
			results = append(results, content)
		}
	}
	return results

}
