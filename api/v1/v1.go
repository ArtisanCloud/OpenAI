package v1

import "github.com/artisancloud/httphelper"

type V1 struct {
	Model       Model
	Completion  Completion
	Chat        Chat
	Edits       Edits
	Images      Images
	Embeddings  Embeddings
	Audio       Audio
	Files       Files
	FineTunes   FineTunes
	Moderations Moderations
}

func NewV1(helper httphelper.Helper) *V1 {
	return &V1{
		Model: Model{
			h: helper,
		},
		Completion: Completion{
			h: helper,
		},
		Chat: Chat{
			h: helper,
		},
		Edits: Edits{
			h: helper,
		},
		Images: Images{
			h: helper,
		},
		Embeddings: Embeddings{
			h: helper,
		},
		Audio: Audio{
			h: helper,
		},
		Files: Files{
			h: helper,
		},
		FineTunes: FineTunes{
			h: helper,
		},
		Moderations: Moderations{
			h: helper,
		},
	}
}
