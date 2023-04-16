package v1

import (
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/openai/openaitype"
)

const (
	createEmbeddingsURI = "/v1/embeddings"
)

type Embeddings struct {
	h httphelper.Helper
}

type CreateEmbeddingsRequest struct {
	// String (Required)
	/* ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.*/
	/* 要使用的模型的 ID。您可以使用 List models API 来查看所有可用模型，或查看我们的模型概述以了解它们的描述。*/
	Model string `json:"model"`

	// StringOrArray (Required)
	/* Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings for multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed 8192 tokens in length.*/
	/* 要获取嵌入的输入文本，编码为字符串或令牌数组。要在单个请求中获取多个输入的嵌入，请传递字符串数组或令牌数组数组。每个输入的长度不得超过 8192 个令牌。*/
	Input openaitype.StringOrArray `json:"input"`

	// String (Optional)
	/* ID of the prompt to use. If not provided, the model's default prompt will be used. You can use the List prompts API to see all of your available prompts, or see our Prompt overview for descriptions of them.*/
	/* 要使用的提示的 ID。如果未提供，则将使用模型的默认提示。您可以使用 List prompts API 来查看所有可用提示，或查看我们的提示概述以了解它们的描述。*/
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids)
	User string `json:"user,omitempty"`
}

type EmbeddingData struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

type EmbeddingsResponse struct {
	Object string          `json:"object"`
	Data   []EmbeddingData `json:"data"`
	Model  string          `json:"model"`
	Usage  struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

func (e *Embeddings) CreateEmbeddings(req *CreateEmbeddingsRequest) (*EmbeddingsResponse, error) {
	var result EmbeddingsResponse
	err := e.h.Df().Uri(createEmbeddingsURI).Method("POST").Json(req).Result(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
