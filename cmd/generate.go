package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

const API_URL = "https://api.mistral.ai/v1/chat/completions"

var (
	star     = ` You are one of the greatest programmers to ever live, you will receive code and your job would be to generate markdown documentation elaborating what the code does. Return ONLY the raw markdown content without any wrapper blocks or formatting. Start directly with the frontmatter (---). Do not include any backticks blocks around the entire response."`
	markdown = " You are one of the greatest programmers to ever live, you will receive code and your job would be to generate markdown documentation elaborating what the code does. You will return markdown and only markdown.Make sure to keep your documentation brief but super clear.Start directly with the frontmatter (---). Do not include any backticks blocks around the entire response. "
	inline   = " You are one of the greatest programmers to ever live, you will receive code and your job would be to add inline comments explaining what the code does. Keep the comments short but clear do not alter the file in any other way than to add comments and do not return anything other than the file provided with the comments.Do not return markdown or any other format just the code you have been given back with comments"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
}

func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: loading .env file: %s", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "API_KEY not found in environment variables")
	}
	return apiKey
}

func readFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file")
	}
	return file
}

func generateMd(apiKey string, writeFilePath string, file []byte) {
	generateMdPayload := ChatRequest{
		Model: "codestral-latest",
		Messages: []Message{
			{
				Role:    "user",
				Content: string(file),
			},
			{
				Role:    "system",
				Content: markdown,
			},
		},
		Temperature: 0.5,
		MaxTokens:   100000,
	}

	payloadBytes, err := json.Marshal(generateMdPayload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: marshaling generateMd: %s", err)
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: creating generateMd: %s\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: making generateMd: %s\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "unexpected status code: %d\n", resp.StatusCode)
	}

	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Fprintf(os.Stderr, "Error: decoding response: %s\n", err)
	}

	if len(response.Choices) > 0 {
		writeFile(writeFilePath, response.Choices[0].Message.Content)
	}

}

func inlineComm(apiKey string, writeFilePath string, file []byte) {
	inlineCommPayload := ChatRequest{
		Model: "codestral-latest",
		Messages: []Message{
			{
				Role:    "user",
				Content: string(file),
			},
			{
				Role:    "system",
				Content: inline,
			},
		},
		Temperature: 0.5,
		MaxTokens:   100000,
	}

	payloadBytes, err := json.Marshal(inlineCommPayload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: marshaling inlineComm: %s\n", err)
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: creating inlineComm: %s\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: making inlineComm: %s\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Fprintf(os.Stderr, "Error: decoding response: %s\n", err)
	}

	if len(response.Choices) > 0 {
		writeFile(writeFilePath, response.Choices[0].Message.Content)
	}

}

func starLight(apiKey string, writeFilePath string, file []byte) {
	starLightPayload := ChatRequest{
		Model: "codestral-latest",
		Messages: []Message{
			{
				Role:    "user",
				Content: string(file),
			},
			{
				Role:    "system",
				Content: star,
			},
		},
		Temperature: 0.5,
		MaxTokens:   10000,
	}

	payloadBytes, err := json.Marshal(starLightPayload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: marshaling starLight: %s\n", err)
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: creating starLight: %s\n", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: making starLight: %s\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Fprintf(os.Stderr, "Error: decoding response: %s\n", err)
	}
	filePath := "./docs/src/content/docs/reference/" + writeFilePath + ".md"

	if len(response.Choices) > 0 {
		writeFile(filePath, response.Choices[0].Message.Content)
	}

}

func getDocs() {

	_, errr := os.ReadDir("./docs")
	if errr != nil {
		cmd := exec.Command("git", "clone", "--depth=1", "https://github.com/chachacollins/chromatemplate.git", "docs")
		_, err := cmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: could not run command: %s\n", err)
		}

		installDocs()
	}
}

func installDocs() {
	cmd := exec.Command("npm", "i")
	cmd.Dir = "./docs"
	_, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not run command: %s", err)
	}
}

func serveDocs() {
	_, err := os.ReadDir("docs")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not find docs directory: %s\n", err)
		os.Exit(1)
	}
	buildDocs()
	runPreviewDocs()
}
func buildDocs() {
	cmd := exec.Command("npm", "run", "build")
	cmd.Dir = "docs"
	_, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not run command: %s\n", err)
	}
}

func runPreviewDocs() {
	cmd := exec.Command("npm", "run", "preview")
	cmd.Dir = "docs"
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not run command: %s\n", err)
	}
	println(string(output))
}

func makeDir(dirName string) {
	err := os.MkdirAll(dirName, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not create a dir: %s", err)
	}
}

func writeFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not create file: %s", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not write to file: %s", err)
	}
	err = writer.Flush()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not flush buffer: %s", err)
	}
}
