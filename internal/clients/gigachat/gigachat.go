package gigachat

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	GigachatAuthorizationDate string
	Client                    *http.Client
}

func New() *Client {
	GigachatAuthorizationDate, exists := os.LookupEnv("GIGACHAT_AUTH_DATE")

	if !exists {
		log.Panic("Gigachat_Auth_Date NOT FOUNT IN .env", GigachatAuthorizationDate)
	}

	caCert, err := os.ReadFile("..\\russian_trusted_sub_ca.cer")
	if err != nil {
		log.Fatal("Check func makeHTTPSClient ->", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientAccess := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	return &Client{GigachatAuthorizationDate: GigachatAuthorizationDate, Client: clientAccess}
}

type accessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expires_at"`
}

var token *accessToken

type body struct {
	Model             string      `json:"model"`
	Messages          []messageGC `json:"messages"`
	Temperature       float32     `json:"temperature"`
	TopT              float32     `json:"top_p"`
	N                 int         `json:"n"`
	Stream            bool        `json:"stream"`
	MaxTokens         int         `json:"max_tokens"`
	RepetitionPenalty int         `json:"repetition_penalty"`
}

type messageGC struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type reqMessage struct {
	MessageGC    messageGC `json:"message"`
	Index        int       `json:"index"`
	FinishReason string    `json:"finish_reason"`
}

type usage struct {
	PromtTokens      int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
	SystemTokens     int `json:"system_tokens"`
}

type choices struct {
	Choices []reqMessage `json:"choices"`
	Created int          `json:"created"`
	Model   string       `json:"model"`
	Usage   usage        `json:"usage"`
	Object  string       `json:"object"`
}

func (client *Client) RequestAuth() (*accessToken, error) {

	rqUID := uuid.New()

	data := url.Values{}
	data.Set("scope", "GIGACHAT_API_PERS")
	req, err := http.NewRequest(
		"POST",
		"https://ngw.devices.sberbank.ru:9443/api/v2/oauth",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("RqUID", rqUID.String()) //это рандом
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", client.GigachatAuthorizationDate))

	c := client.Client
	response, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("сould not get a response: %w", err)
	}

	defer response.Body.Close()

	var result accessToken

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return nil, fmt.Errorf("NewDecoder cound not return decoder %w", err)
	}

	return &result, nil

}

func (client *Client) Request(text string) ([]string, error) {

	if token == nil || token.ExpiresAt <= int(time.Now().UnixMilli()) {
		token, _ = client.RequestAuth()
	}
	body := body{
		Model: "GigaChat:latest",
		Messages: []messageGC{
			{
				Role:    "user",
				Content: text,
			},
		},
		Temperature:       1.0,
		TopT:              0.1,
		N:                 1,
		Stream:            false,
		MaxTokens:         512,
		RepetitionPenalty: 1,
	}
	bytesRepresentation, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://gigachat.devices.sberbank.ru/api/v1/chat/completions",
		bytes.NewBuffer(bytesRepresentation),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	c := client.Client
	response, err := c.Do(req)

	if err != nil {
		log.Println("%w", err)
		return nil, err
	}

	defer response.Body.Close()

	var resAns choices

	err = json.NewDecoder(response.Body).Decode(&resAns)

	if err != nil {
		return nil, fmt.Errorf("NewDecoder cound not return decoder %w", err)
	}

	answers := make([]string, 0, len(resAns.Choices))

	for _, answer := range resAns.Choices {

		answers = append(answers, answer.MessageGC.Content)
	}

	return answers, nil
}
