package openai

import (
	"app/internal/common"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

type body struct {
	Model          string       `json:"model"`
	ResponseFormat jsonResponse `json:"response_format"`
	Messages       []messageGpt `json:"messages"`
}

type jsonResponse struct {
	TypeResp string `json:"type"`
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
		Model:          "gpt-3.5-turbo",
		ResponseFormat: jsonResponse{TypeResp: "json_object"},
		Messages: []messageGpt{
			{
				Role:    "system",
				Content: "Ты помощник, который работает в России. А также возвращай сообщения в формате json",
			},
			{
				Role:    "user",
				Content: text,
			},
		},
	}

	bytesRepresentation, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(bytesRepresentation),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", common.OPENAIAPIKEY))

	c := &http.Client{}
	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result responseGpt

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	fmt.Println(result.Choices)
	answers := make([]string, 0, len(result.Choices))

	for _, answer := range result.Choices {
		answers = append(answers, answer.Message.Content)
	}
	return answers, nil
}
