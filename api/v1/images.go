package v1

import (
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/httphelper/dataflow"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strconv"
)

const (
	createImageURI          = "/v1/images/generations"
	createImageEditURI      = "/v1/images/edits"
	createImageVariationURI = "/v1/images/variations"
)

type Images struct {
	h httphelper.Helper
}

type CreateImageRequest struct {
	// String (Required)
	/* A text description of the desired image(s). The maximum length is 1000 characters.*/
	/* 对所需图像的文本描述。最大长度为1000个字符。*/
	Prompt string `json:"prompt"`

	// Integer (Optional) Defaults to 1
	/* The number of images to generate. Must be between 1 and 10.*/
	/* 要生成的图像数。必须介于1和10之间。*/
	N *int `json:"n,omitempty"`

	// String (Optional) Defaults to 1024x1024
	/* The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.*/
	/* 生成图像的尺寸。必须是256x256、512x512或1024x1024之一。*/
	Size string `json:"size,omitempty"`

	// String (Optional) Defaults to url
	/* The format in which the generated images are returned. Must be one of url or b64_json.*/
	/* 以何种格式返回生成的图像。必须是url或b64_json之一。*/
	ResponseFormat string `json:"response_format,omitempty"`

	// String (Optional)
	/* A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.*/
	/* 代表您的最终用户的唯一标识符，可以帮助OpenAI监控和检测滥用。*/
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids)
	User string `json:"user,omitempty"`
}

type CreateImageEditRequest struct {
	// FileReader (Required)
	/* The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.*/
	/* 用作变体的基础的图像。必须是有效的PNG文件，小于4MB，并且是正方形的。*/
	Image *os.File

	// FileReader (Optional)
	/* An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where image should be
	edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as image.*/
	/* 另一个图像，其完全透明的区域（例如，其中alpha为零）指示应编辑图像的位置。必须是有效的PNG文件，小于4MB，并且具有与图像相同的尺寸。*/
	Mask *os.File

	// String (Required)
	/* A text description of the desired image(s). The maximum length is 1000 characters.*/
	/* 对所需图像的文本描述。最大长度为1000个字符。*/
	Prompt string

	// Integer (Optional) Defaults to 1
	/* The number of images to generate. Must be between 1 and 10.*/
	/* 要生成的图像数。必须介于1和10之间。*/
	N *int

	// String (Optional) Defaults to 1024x1024
	/* The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.*/
	/* 生成图像的尺寸。必须是256x256、512x512或1024x1024之一。*/
	Size string

	// String (Optional) Defaults to url
	/* The format in which the generated images are returned. Must be one of url or b64_json.*/
	/* 以何种格式返回生成的图像。必须是url或b64_json之一。*/
	ResponseFormat string

	// String (Optional)
	/* A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.*/
	/* 代表您的最终用户的唯一标识符，可以帮助OpenAI监控和检测滥用。*/
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids)
	User string
}

type CreateImageVariationRequest struct {
	// FileReader (Required)
	/* The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.*/
	/* 用作变体的基础的图像。必须是有效的PNG文件，小于4MB，并且是正方形的。*/
	Image *os.File

	// Integer (Optional) Defaults to 1
	/* The number of images to generate. Must be between 1 and 10.*/
	/* 要生成的图像数。必须介于1和10之间。*/
	N *int

	// String (Optional) Defaults to 1024x1024
	/* The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.*/
	/* 生成图像的尺寸。必须是256x256、512x512或1024x1024之一。*/
	Size string

	// String (Optional) Defaults to url
	/* The format in which the generated images are returned. Must be one of url or b64_json.*/
	/* 以何种格式返回生成的图像。必须是url或b64_json之一。*/
	ResponseFormat string

	// String (Optional)
	/* A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse.*/
	/* 代表您的最终用户的唯一标识符，可以帮助OpenAI监控和检测滥用。*/
	User string
}

type ImageData struct {
	URL string `json:"url"`
}

type ImageResponse struct {
	Created int64       `json:"created"`
	Data    []ImageData `json:"data"`
}

func (i *Images) CreateImage(req *CreateImageRequest) (*ImageResponse, error) {
	var result ImageResponse
	err := i.h.Df().Uri(createImageURI).Method("POST").Json(req).Result(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *Images) CreateImageEdit(req *CreateImageEditRequest) (*ImageResponse, error) {
	var result ImageResponse
	err := i.h.Df().Uri(createImageEditURI).Method("POST").Multipart(func(multipart dataflow.MultipartDataflow) error {
		if req.Image == nil {
			return errors.New("image is required")
		}
		if req.Prompt == "" {
			return errors.New("prompt is required")
		}

		multipart.FileMem("image", filepath.Base(req.Image.Name()), req.Image)
		multipart.FieldValue("prompt", req.Prompt)
		if req.Mask != nil {
			multipart.FileMem("mask", filepath.Base(req.Mask.Name()), req.Mask)
		}
		if req.N != nil {
			multipart.FieldValue("n", strconv.Itoa(*req.N))
		}
		if req.Size != "" {
			multipart.FieldValue("size", req.Size)
		}
		if req.ResponseFormat != "" {
			multipart.FieldValue("response_format", req.ResponseFormat)
		}
		if req.User != "" {
			multipart.FieldValue("user", req.User)
		}
		return nil
	}).Result(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (i *Images) CreateImageVariation(req *CreateImageVariationRequest) (*ImageResponse, error) {
	var result ImageResponse
	err := i.h.Df().Uri(createImageVariationURI).Method("POST").Multipart(func(multipart dataflow.MultipartDataflow) error {
		multipart.FileMem("image", filepath.Base(req.Image.Name()), req.Image)
		if req.N != nil {
			multipart.FieldValue("n", strconv.Itoa(*req.N))
		}
		if req.Size != "" {
			multipart.FieldValue("size", req.Size)
		}
		if req.ResponseFormat != "" {
			multipart.FieldValue("response_format", req.ResponseFormat)
		}
		if req.User != "" {
			multipart.FieldValue("user", req.User)
		}
		return nil
	}).Result(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
