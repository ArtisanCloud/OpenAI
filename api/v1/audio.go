package v1

import (
	"fmt"
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/httphelper/dataflow"
	"os"
	"path/filepath"
)

const (
	createTranscriptionURI = "/v1/audio/transcriptions"
	createTranslationURI   = "/v1/audio/translations"
)

type Audio struct {
	h httphelper.Helper
}

type CreateTranscriptionRequest struct {
	// FileReader (Required)
	/* The audio file to translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.*/
	/* 要翻译的音频文件，格式为mp3、mp4、mpeg、mpga、m4a、wav或webm。*/
	File *os.File

	// String (Required)
	/* ID of the model to use. Only whisper-1 is currently available.*/
	/* 要使用的模型的ID。目前仅支持whisper-1。*/
	Model string

	// String (Optional)
	/* An optional text to guide the model's style or continue a previous audio segment. The prompt should be in English.*/
	/* 可选的文本，用于指导模型的风格或继续以前的音频段。提示应为英语。*/
	Prompt string

	// String (Optional) Defaults to json
	/* The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.*/
	/* 字幕输出的格式，可选项为json、text、srt、verbose_json或vtt。*/
	ResponseFormat string

	// Float (Optional) Defaults to 0
	/* The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.*/
	/* 采样温度，介于0和1之间。更高的值（如0.8）将使输出更随机，而更低的值（如0.2）将使其更专注和确定。如果设置为0，则模型将使用对数概率自动增加温度，直到达到某些阈值。*/
	Temperature *float64

	// String (Optional)
	/* The language of the input audio. Supplying the input language in ISO-639-1 format will improve accuracy and latency.*/
	/* 输入音频的语言。以ISO-639-1格式提供输入语言将提高准确性和延迟。*/
	Language string
}

type CreateTranslationRequest struct {
	// FileReader (Required)
	/* The audio file to translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.*/
	/* 要翻译的音频文件，格式为mp3、mp4、mpeg、mpga、m4a、wav或webm。*/
	File *os.File

	// String (Required)
	/* ID of the model to use. Only whisper-1 is currently available.*/
	/* 要使用的模型的ID。目前仅支持whisper-1。*/
	Model string

	// String (Optional)
	/* An optional text to guide the model's style or continue a previous audio segment. The prompt should be in English.*/
	/* 可选的文本，用于指导模型的风格或继续以前的音频段。提示应为英语。*/
	Prompt string

	// String (Optional) Defaults to json
	/* The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.*/
	/* 字幕输出的格式，可选项为json、text、srt、verbose_json或vtt。*/
	ResponseFormat string

	// String (Optional)
	/* The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use log probability to automatically increase the temperature until certain thresholds are hit.*/
	/* 采样温度，介于0和1之间。更高的值（如0.8）将使输出更随机，而更低的值（如0.2）将使其更专注和确定。如果设置为0，则模型将使用对数概率自动增加温度，直到达到某些阈值。*/
	Temperature float64
}

type TranscriptionResponse struct {
	Text string `json:"text"`
}

func (a *Audio) CreateTranscription(req *CreateTranscriptionRequest) (*TranscriptionResponse, error) {
	var result TranscriptionResponse
	err := a.h.Df().Uri(createTranscriptionURI).Method("POST").Multipart(func(multipart dataflow.MultipartDataflow) error {
		if req.File != nil {
			multipart.FileMem("file", filepath.Base(req.File.Name()), req.File)
		}
		multipart.FieldValue("model", req.Model)
		if req.Prompt != "" {
			multipart.FieldValue("prompt", req.Prompt)
		}
		if req.ResponseFormat != "" {
			multipart.FieldValue("response_format", req.ResponseFormat)
		}
		if req.Temperature != nil {
			multipart.FieldValue("temperature", fmt.Sprintf("%.1f", *req.Temperature))
		}
		if req.Language != "" {
			multipart.FieldValue("language", req.Language)
		}
		return nil
	}).Result(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *Audio) CreateTranslation(req *CreateTranslationRequest) (*TranscriptionResponse, error) {
	var result TranscriptionResponse
	err := a.h.Df().Uri(createTranslationURI).Method("POST").Multipart(func(multipart dataflow.MultipartDataflow) error {
		if req.File != nil {
			multipart.FileMem("file", filepath.Base(req.File.Name()), req.File)
		}
		multipart.FieldValue("model", req.Model)
		if req.Prompt != "" {
			multipart.FieldValue("prompt", req.Prompt)
		}
		if req.ResponseFormat != "" {
			multipart.FieldValue("response_format", req.ResponseFormat)
		}
		if req.Temperature != 0 {
			multipart.FieldValue("temperature", fmt.Sprintf("%.1f", req.Temperature))
		}
		return nil
	}).Result(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
