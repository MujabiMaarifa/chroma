package cfg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/chachacollins/chroma/utils"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

const (
	default_star     = ` You are one of the greatest programmers to ever live, you will receive code and your job would be to generate markdown documentation elaborating what the code does and provide examples where necessary. Return ONLY the raw markdown content without any wrapper blocks or formatting. Start directly with the frontmatter (---). Do not include any backticks blocks around the entire response."`
	default_markdown = " You are one of the greatest programmers to ever live, you will receive code and your job would be to generate markdown documentation elaborating what the code does. You will return markdown and only markdown.Make sure to keep your documentation brief but super clear.Start directly with the frontmatter (---). Do not include any backticks blocks around the entire response. "
	default_inline   = " You are one of the greatest programmers to ever live, you will receive code and your job would be to add inline comments explaining what the code does. Keep the comments short but clear do not alter the file in any other way than to add comments and do not return anything other than the file provided with the comments.Do not return markdown or any other format just the code you have been given back with comments"
)

type Config struct {
	Apikey         string `json:"apikey"`
	MarkdownPrompt string `json:"md_prompt"`
	InlinePrompt   string `json:"il_prompt"`
	StarPrompt     string `json:"star_prompt"`
}

func new() *Config {
	return &Config{}
}

func Init() {
	value := renderTextInput("Enter your mistral ai api key ")
	config := new()
	config.Apikey = value
	config.MarkdownPrompt = default_markdown
	config.InlinePrompt = default_inline
	config.StarPrompt = default_star
	payLoad, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Could not Marshal json")
		os.Exit(1)
	}
	utils.WriteFile("config.json", string(payLoad))
}

func renderTextInput(prompt string) string {
	input := textinput.New()
	input.Placeholder = "Type here..."
	input.Focus()

	style := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	fmt.Println(style.Render(prompt))

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input.SetValue(scanner.Text())
	}

	return input.Value()
}
