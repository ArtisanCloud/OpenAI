package v1

import (
	"fmt"
	"github.com/artisancloud/httphelper"
)

// GET https://api.openai.com/v1/models
const listModelsURI = "/v1/models"

// GET https://api.openai.com/v1/models/{model}
const retrieveModelURI = "/v1/models/%s"

type Model struct {
	h httphelper.Helper
}

type Permission map[string]interface{}

type ListModelsResultData struct {
	ID         string       `json:"id"`
	Object     string       `json:"object"`
	OwnedBy    string       `json:"owned_by"`
	Permission []Permission `json:"permission"`
}

type ListModelsResult struct {
	Data   []ListModelsResultData `json:"data"`
	Object string                 `json:"object"`
	Root   string                 `json:"root"`
	// todo un_know type
	Parent interface{} `json:"parent"`
	*Error
}

// ListModels https://platform.openai.com/docs/api-reference/models/list
func (m *Model) ListModels() (*ListModelsResult, error) {
	var result ListModelsResult
	err := m.h.Df().Uri(listModelsURI).Method("GET").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

type RetrieveModelResult struct {
	ID         string       `json:"id"`
	Object     string       `json:"object"`
	OwnedBy    string       `json:"owned_by"`
	Permission []Permission `json:"permission"`
	Root       string       `json:"root"`
	// todo un_know type
	Parent interface{} `json:"parent"`
	*Error
}

// RetrieveModel https://platform.openai.com/docs/api-reference/models/retrieve
func (m *Model) RetrieveModel(model string) (*RetrieveModelResult, error) {
	var result RetrieveModelResult
	err := m.h.Df().Uri(fmt.Sprintf(retrieveModelURI, model)).Method("GET").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}
