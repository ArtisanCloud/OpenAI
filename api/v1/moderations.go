package v1

import (
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/openai/openaitype"
)

const createModerationURI = "/v1/moderations"

type Moderations struct {
	h httphelper.Helper
}

type CreateModerationRequest struct {
	Input openaitype.StringOrArray `json:"input"`
	Model string                   `json:"model,omitempty"`
}

type CreateModerationResult struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Categories     map[string]bool    `json:"categories"`
		CategoryScores map[string]float64 `json:"category_scores"`
		Flagged        bool               `json:"flagged"`
	} `json:"results"`
	*Error
}

func (m *Moderations) CreateModeration(req *CreateModerationRequest) (*CreateModerationResult, error) {
	var result CreateModerationResult
	err := m.h.Df().Uri(createModerationURI).Method("POST").Json(req).Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}
