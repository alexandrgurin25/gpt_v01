package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Client struct {
	openApiKey string
}

func New() *Client {
	openApiKey, exists := os.LookupEnv("OPENAI_API_KEY")

	if !exists {
		log.Panic("OPENAI_API_KEY NOT FOUNT IN .env")	
	}

	return &Client{openApiKey: openApiKey}
}

type body struct {
	Model    string       `json:"model"`
	Messages []messageGpt `json:"messages"`
}

type messageGpt struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type responseGpt struct {
	Id           string   `json:"id"`
	Object       string   `json:"object"`
	Created      int      `json:"created"`
	Model        string   `json:"model"`
	SystemFinger string   `json:"system_fingerprint"`
	Choices      []choice `json:"choices"`
	Usage        usage    `json:"usage"`
}

type usage struct {
	PromptTokent     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type choice struct {
	Index        int        `json:"index"`
	Message      messageGpt `json:"message"`
	Logprobs     *string    `json:"logprobs"`
	FinishReason string     `json:"finish_reason"`
}

func (client *Client) Request(text string) ([]string, error) {
	body := body{
		Model: "gpt-3.5-turbo",
		Messages: []messageGpt{
			{
				Role:    "system",
				Content: "Ты помощник, который работает в России.",
			},
			{
				Role:    "user",
				Content: text,
			},
		},
	}

	bytesRepresentation, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal could not return JSON encoding: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(bytesRepresentation),
	)
	if err != nil {
		return nil, fmt.Errorf("NewRequest could not send request: %w", err)
	}

	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.openApiKey))

	c := &http.Client{}
	response, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("сould not get a response: %w", err) 
	}

	defer response.Body.Close()

	var result responseGpt

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return nil, fmt.Errorf("NewDecoder cound not return decoder %w", err)
	}

	answers := make([]string, 0, len(result.Choices))

	for _, answer := range result.Choices {
		answers = append(answers, answer.Message.Content)
	}
	return answers, nil
}
