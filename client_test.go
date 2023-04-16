package openai

import (
	"bytes"
	v1 "github.com/artisancloud/openai/api/v1"
	"github.com/artisancloud/openai/openaitype"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

/*
 *  SDK for OpenAI API V1 TEST
 *  由于部分测试所需参数涉及到平台状态及数据, 如果您需要自己测试, 请将测试代码中的参数替换为自己的参数
 */

// 注意此处的key是测试用的，已被废弃，如果你需要测试, 请到 https://platform.openai.com/ 申请自己的key
// 如果你设置了自己的key， 请不要上传到公开仓库
const openAIKey = "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"

var testClient *Client

func TestMain(m *testing.M) {
	config := V1Config{
		OpenAPIKey:   openAIKey,
		Organization: "org-fklWRKyhwySGHR6Ki37W6KtY",
		HttpDebug:    true,
	}
	client, err := NewClient(&config)
	if err != nil {
		panic(err)
	}
	testClient = client

	if config.HttpDebug {
		file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		//defer file.Close()

		log.SetOutput(file)
		log.Println("start test...")
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestListModels(t *testing.T) {
	models, err := testClient.V1.Model.ListModels()
	if err != nil {
		t.Error(err)
	}
	t.Log(models)
}

func TestRetrieveModel(t *testing.T) {
	model, err := testClient.V1.Model.RetrieveModel("davinci")
	if err != nil {
		t.Error(err)
	}
	t.Log(model)
}

func TestCreateCompletion(t *testing.T) {
	req := v1.CreateCompletionRequest{
		Model:       "text-davinci-003",
		Prompt:      openaitype.StringOrArray{"Say this is a test"},
		MaxTokens:   openaitype.PointInt(7),
		Temperature: openaitype.PointFloat64(0),
		TopP:        openaitype.PointFloat64(1),
		N:           openaitype.PointInt(1),
		Stream:      openaitype.PointBool(false),
		Stop:        openaitype.StringOrArray{"\n"},
	}
	result, err := testClient.V1.Completion.CreateCompletion(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateChatCompletion(t *testing.T) {
	req := v1.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []v1.Message{
			{
				Role:    "user",
				Content: "Hello!",
			},
		},
	}
	result, err := testClient.V1.Chat.CreateChatCompletion(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateEdit(t *testing.T) {
	req := v1.CreateEditRequest{
		Model:       "text-davinci-edit-001",
		Input:       "What day of the wek is it?",
		Instruction: "Fix the spelling mistakes",
	}
	result, err := testClient.V1.Edits.CreateEdit(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateImage(t *testing.T) {
	req := v1.CreateImageRequest{
		Prompt: "A cute baby sea otter",
		N:      openaitype.PointInt(2),
		Size:   "256x256",
	}
	result, err := testClient.V1.Images.CreateImage(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func downloadFile(url string, ext string) (*os.File, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to download: non-200 status code")
	}

	file, err := os.CreateTemp("", "*."+ext)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		file.Close()
		os.Remove(file.Name())
		return nil, err
	}

	if err := file.Sync(); err != nil {
		file.Close()
		os.Remove(file.Name())
		return nil, err
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		file.Close()
		os.Remove(file.Name())
		return nil, err
	}

	return file, nil
}

func TestCreateImageEdit(t *testing.T) {
	imageUrl := "https://pic.rmb.bdstatic.com/bjh/news/54bfff96fd5019478442ba17ae7317cc4121.png"
	maskUrl := "https://pic.rmb.bdstatic.com/bjh/news/99115c994cb1928647eb101357e91f383500.png"

	imageFile, err := downloadFile(imageUrl, "png")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(imageFile.Name())

	maskFile, err := downloadFile(maskUrl, "png")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(maskFile.Name())

	req := v1.CreateImageEditRequest{
		Image:  imageFile,
		Mask:   maskFile,
		Prompt: "A cute baby sea otter wearing a beret",
		N:      openaitype.PointInt(2),
		Size:   "256x256",
	}
	result, err := testClient.V1.Images.CreateImageEdit(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateImageVariation(t *testing.T) {
	imageUrl := "https://pic.rmb.bdstatic.com/bjh/news/54bfff96fd5019478442ba17ae7317cc4121.png"

	imageFile, err := downloadFile(imageUrl, "png")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(imageFile.Name())

	req := v1.CreateImageVariationRequest{
		Image: imageFile,
		N:     openaitype.PointInt(2),
		Size:  "256x256",
	}
	result, err := testClient.V1.Images.CreateImageVariation(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateEmbeddings(t *testing.T) {
	req := v1.CreateEmbeddingsRequest{
		Model: "text-embedding-ada-002",
		Input: openaitype.StringOrArray{"The food was delicious and the waiter..."},
	}
	result, err := testClient.V1.Embeddings.CreateEmbeddings(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateTranscription(t *testing.T) {
	const audioUrl = "https://mp3.t57.cn:7087/kwlink_d.php?id=149121488"
	audioFile, err := downloadFile(audioUrl, "mp3")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(audioFile.Name())

	req := v1.CreateTranscriptionRequest{
		Model: "whisper-1",
		File:  audioFile,
	}
	result, err := testClient.V1.Audio.CreateTranscription(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateTranslation(t *testing.T) {
	const audioUrl = "https://mp3.t57.cn:7087/kwlink_d.php?id=149121488"
	audioFile, err := downloadFile(audioUrl, "mp3")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(audioFile.Name())

	req := v1.CreateTranslationRequest{
		Model: "whisper-1",
		File:  audioFile,
	}
	result, err := testClient.V1.Audio.CreateTranslation(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestListFiles(t *testing.T) {
	result, err := testClient.V1.Files.ListFiles()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestUploadFile(t *testing.T) {
	const jsonl = `{"prompt":"aa", "completion": "bb"}
		{"prompt":"cc", "completion": " dd"}`
	jsonFile, err := os.CreateTemp("", "*.jsonl")
	jsonFile.WriteString(jsonl)
	jsonFile.Seek(0, io.SeekStart)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(jsonFile.Name())

	result, err := testClient.V1.Files.UploadFile(jsonFile, "fine-tune")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestDeleteFile(t *testing.T) {
	// file id must exist
	result, err := testClient.V1.Files.DeleteFile("file-cWnxIx8vsqukp60D31Rjhwwy")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestRetrieveFile(t *testing.T) {
	result, err := testClient.V1.Files.RetrieveFile("file-V5Sq8CFxiC8spmCaszWa2lCz")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestRetrieveFileContent(t *testing.T) {
	var buf bytes.Buffer
	err := testClient.V1.Files.RetrieveFileContent("file-V5Sq8CFxiC8spmCaszWa2lCz", &buf)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(buf.Bytes()))
}

func TestCreateFineTune(t *testing.T) {
	req := v1.CreateFineTuneRequest{
		TrainingFile: "file-V5Sq8CFxiC8spmCaszWa2lCz",
	}
	result, err := testClient.V1.FineTunes.CreateFineTune(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestListFineTunes(t *testing.T) {
	result, err := testClient.V1.FineTunes.ListFineTunes()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestRetrieveFineTune(t *testing.T) {
	result, err := testClient.V1.FineTunes.RetrieveFineTune("ft-xwD8fHYArSFxP1kJ66pV3Om7")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCancelFineTune(t *testing.T) {
	result, err := testClient.V1.FineTunes.CancelFineTune("ft-xwD8fHYArSFxP1kJ66pV3Om7")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestListFineTuneEvents(t *testing.T) {
	result, err := testClient.V1.FineTunes.ListFineTuneEvents(&v1.ListFineTuneEventsRequest{
		FineTuneID: "ft-xwD8fHYArSFxP1kJ66pV3Om7",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestDeleteFineTuneModel(t *testing.T) {
	result, err := testClient.V1.FineTunes.DeleteFineTuneModel("curie:ft-personal-2023-04-16-02-36-14")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestCreateModeration(t *testing.T) {
	req := v1.CreateModerationRequest{
		Input: openaitype.StringOrArray{"I want to kill them"},
	}
	result, err := testClient.V1.Moderations.CreateModeration(&req)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
