package v1

import (
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/openai/openaitype"
)

const createChatCompletionURI = "/v1/chat/completions"

type Chat struct {
	h httphelper.Helper
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CreateChatCompletionRequest struct {
	Model            string                   `json:"model"`
	Messages         []Message                `json:"messages"`
	Temperature      *float64                 `json:"temperature,omitempty"`
	TopP             *float64                 `json:"top_p,omitempty"`
	N                *int                     `json:"n,omitempty"`
	Stream           *bool                    `json:"stream,omitempty"`
	Stop             openaitype.StringOrArray `json:"stop,omitempty"`
	MaxTokens        *int                     `json:"max_tokens,omitempty"`
	PresencePenalty  *float64                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64                 `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int           `json:"logit_bias,omitempty"`
	User             string                   `json:"user,omitempty"`
}

type ChatChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type CreateChatCompletionResult struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Choices []ChatChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
	*Error
}

func (c *Chat) CreateChatCompletion(req *CreateChatCompletionRequest) (*CreateChatCompletionResult, error) {
	var result CreateChatCompletionResult
	err := c.h.Df().Uri(createChatCompletionURI).Method("POST").Json(req).Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}
