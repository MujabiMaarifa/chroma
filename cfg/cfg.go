package cfg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/chachacollins/chroma/utils"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

const (
	// Default prompt for generating markdown documentation
	default_star = ` You are one of the greatest programmers to ever live, you will receive code and your job would be to generate markdown documentation elaborating what the code does and provide examples where necessary. Return ONLY the raw markdown content without any wrapper blocks or formatting. Start directly with the frontmatter (---). Do not include any backticks blocks around the entire response."`

	// Default prompt for generating markdown documentation
	default_markdown = " You are one of the greatest programmers to ever live, you will receive code and your job would be to generate markdown documentation elaborating what the code does. You will return markdown and only markdown.Make sure to keep your documentation brief but super clear.Start directly with the frontmatter (---). Do not include any backticks blocks around the entire response. "

	// Default prompt for adding inline comments
	default_inline = " You are one of the greatest programmers to ever live, you will receive code and your job would be to add inline comments explaining what the code does. Keep the comments short but clear do not alter the file in any other way than to add comments and do not return anything other than the file provided with the comments.Start directly with the frontmatter.Do not include any backticks (```) blocks around the entire response."
)

// Config struct holds the configuration settings
type Config struct {
	Apikey         string `json:"apikey"`      // API key for the service
	MarkdownPrompt string `json:"md_prompt"`   // Prompt for generating markdown documentation
	InlinePrompt   string `json:"il_prompt"`   // Prompt for adding inline comments
	StarPrompt     string `json:"star_prompt"` // Prompt for generating star documentation
}

// new creates a new Config instance
func new() *Config {
	return &Config{}
}

// Init initializes the configuration settings
func Init() {
	// Render text input for API key
	value := renderTextInput("Enter your mistral ai api key ")
	config := new()
	config.Apikey = value
	config.MarkdownPrompt = default_markdown
	config.InlinePrompt = default_inline
	config.StarPrompt = default_star

	// Marshal the config to JSON
	payLoad, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Could not Marshal json")
		os.Exit(1)
	}

	// Determine the runtime OS and set the config path accordingly
	runTime := runtime.GOOS
	switch runTime {
	case "windows":
		appdata := os.Getenv("APPDATA")
		if appdata == "" {
			fmt.Fprint(os.Stderr, "Error: could not find appdata path")
			os.Exit(1)
		}
		configPath := filepath.Join(appdata, "chroma", "config.json")
		utils.WriteFile(configPath, string(payLoad))
	case "linux":
		configDir := filepath.Join(os.Getenv("HOME"), "chroma")
		utils.MakeDir(configDir)
		utils.WriteFile(configDir+"/config.json", string(payLoad))
	}
}

// Load loads the configuration settings from the file
func Load() Config {
	configDir := filepath.Join(os.Getenv("HOME"), "chroma")
	var config Config

	file, err := os.Open(configDir + "/config.json")
	if err != nil {
		return Config{}
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	err = json.NewDecoder(reader).Decode(&config)
	if err != nil {
		fmt.Println("could not decode json")
		os.Exit(1)
	}
	return config
}

// renderTextInput renders a text input prompt
func renderTextInput(prompt string) string {
	input := textinput.New()
	input.Placeholder = "Type here..."
	input.Focus()
	input.PlaceholderStyle.Italic(true)
	input.PromptStyle.Border(lipgloss.NormalBorder())

	style := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	fmt.Println(style.Render(prompt))
	fmt.Print(style.Render(">> "))

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input.SetValue(scanner.Text())
	}

	return input.Value()
}
