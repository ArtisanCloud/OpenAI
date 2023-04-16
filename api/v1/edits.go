package v1

import (
	"github.com/artisancloud/httphelper"
)

const createEditURI = "/v1/edits"

type Edits struct {
	h httphelper.Helper
}

type CreateEditRequest struct {
	Model       string   `json:"model"`
	Input       string   `json:"input,omitempty"`
	Instruction string   `json:"instruction"`
	N           *int     `json:"n,omitempty"`
	Temperature *float64 `json:"temperature,omitempty"`
	TopP        *float64 `json:"top_p,omitempty"`
}

type EditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type CreateEditResult struct {
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Choices []EditChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
	*Error
}

func (e *Edits) CreateEdit(req *CreateEditRequest) (*CreateEditResult, error) {
	var result CreateEditResult
	err := e.h.Df().Uri(createEditURI).Method("POST").Json(req).Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}
