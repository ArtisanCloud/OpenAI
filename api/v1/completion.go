package v1

import (
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/openai/openaitype"
)

const createCompletionURI = "/v1/completions"

type Completion struct {
	h httphelper.Helper
}

type CreateCompletionRequest struct {
	Model            string                   `json:"model"`
	Prompt           openaitype.StringOrArray `json:"prompt,omitempty"`
	Suffix           string                   `json:"suffix,omitempty"`
	MaxTokens        *int                     `json:"max_tokens,omitempty"`
	Temperature      *float64                 `json:"temperature"`
	TopP             *float64                 `json:"top_p,omitempty"`
	N                *int                     `json:"n,omitempty"`
	Stream           *bool                    `json:"stream,omitempty"`
	Logprobs         *int                     `json:"logprobs,omitempty"`
	Echo             *bool                    `json:"echo,omitempty"`
	Stop             openaitype.StringOrArray `json:"stop,omitempty"`
	PresencePenalty  *float64                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64                 `json:"frequency_penalty,omitempty"`
	BestOf           *int                     `json:"best_of,omitempty"`
	LogitBias        map[string]int           `json:"logit_bias,omitempty"`
	User             string                   `json:"user,omitempty"`
}

type Choice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type CreateCompletionResult struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
	*Error
}

func (c *Completion) CreateCompletion(req *CreateCompletionRequest) (*CreateCompletionResult, error) {
	var result CreateCompletionResult
	err := c.h.Df().Uri(createCompletionURI).Method("POST").Json(req).Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}
