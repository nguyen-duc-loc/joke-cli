package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Joke struct {
	Category string          `json:"category"`
	Type     string          `json:"type"`
	Content  string          `json:"joke"`
	Flags    map[string]bool `json:"flags"`
	Language string          `json:"lang"`
}

type GenerateRequest struct {
	Language       string
	BlacklistFlags []string
	ResponseFormat string
	Amount         uint8
}

type GenerateResponse struct {
	Error  bool   `json:"error"`
	Amount uint8  `json:"amount"`
	Jokes  []Joke `json:"jokes"`
	Joke   string `json:"joke"` // In case there is only one joke in response
}

func NewRequest(language string, blacklistFlags []string, responseFormat string, amount uint8) GenerateRequest {
	return GenerateRequest{
		Language:       language,
		BlacklistFlags: blacklistFlags,
		ResponseFormat: responseFormat,
		Amount:         amount,
	}
}

func GenerateJokes(request GenerateRequest) ([]Joke, error) {
	language := fmt.Sprintf("lang=%s", request.Language)
	blacklistFlags := fmt.Sprintf("blacklistFlags=%s", strings.Join(request.BlacklistFlags, ","))
	format := fmt.Sprintf("format=%s", request.ResponseFormat)
	jokeType := "type=single"
	amount := fmt.Sprintf("amount=%d", request.Amount)
	url := fmt.Sprintf("https://v2.jokeapi.dev/joke/Programming?%s", strings.Join([]string{language, blacklistFlags, format, jokeType, amount}, "&"))
	result := []Joke{}

	response, err := http.Get(url)
	if err != nil {
		return result, errors.New("opps!!! something went wrong, please try again later")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	err = response.Body.Close()
	if err != nil {
		return result, err
	}

	var responseBody GenerateResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return result, err
	}

	if len(responseBody.Jokes) == 0 {
		result = []Joke{{Content: responseBody.Joke}}
	} else {
		result = responseBody.Jokes
	}

	return result, nil
}
