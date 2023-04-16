package v1

import (
	"fmt"
	"github.com/artisancloud/httphelper"
	"strconv"
)

const createFineTuneURI = "/v1/fine-tunes"
const listFineTunesURI = "/v1/fine-tunes"
const retrieveFineTuneURI = "/v1/fine-tunes/%s"
const cancelFineTuneURI = "/v1/fine-tunes/%s/cancel"
const fineTuneEventsURI = "/v1/fine-tunes/%s/events"
const deleteFineTuneModelURI = "/v1/models/%s"

type FineTunes struct {
	h httphelper.Helper
}

type CreateFineTuneRequest struct {
	/* String (Required) */
	/**/
	// The ID of an uploaded file that contains training data.
	// See upload file for how to upload a file.
	// Your dataset must be formatted as a JSONL file, where each training example is a JSON object with the keys "prompt" and "completion". Additionally, you must upload your file with the purpose fine-tune.
	// See the fine-tuning guide for more details.
	/**/
	// 包含训练数据的上传文件的ID。请参阅上传文件以了解如何上传文件。
	// 您的数据集必须格式化为JSONL文件，其中每个训练示例都是具有键“提示”和“完成”的JSON对象。此外，您必须使用微调目的上传您的文件。
	// 有关详细信息，请参阅微调指南。
	TrainingFile string `json:"training_file"`

	/* String (Optional) */
	/**/
	// The ID of an uploaded file that contains validation data.
	// If you provide this file, the data is used to generate validation metrics periodically during fine-tuning. These metrics can be viewed in the fine-tuning results file. Your train and validation data should be mutually exclusive.
	// Your dataset must be formatted as a JSONL file, where each validation example is a JSON object with the keys "prompt" and "completion". Additionally, you must upload your file with the purpose fine-tune.
	// See the fine-tuning guide for more details.
	/**/
	// 包含验证数据的上传文件的ID。
	// 如果您提供此文件，则该数据将在微调期间定期生成验证指标。这些指标可以在微调结果文件中查看。您的训练和验证数据应互斥。
	// 您的数据集必须格式化为JSONL文件，其中每个验证示例都是具有键“提示”和“完成”的JSON对象。此外，您必须使用微调目的上传您的文件。
	// 有关详细信息，请参阅微调指南。
	ValidationFile string `json:"validation_file,omitempty"`

	/* String (Optional) Defaults to curie */
	/**/
	// The name of the base model to fine-tune. You can select one of "ada", "babbage", "curie", "davinci", or a fine-tuned model created after 2022-04-21. To learn more about these models, see the Models documentation.
	/**/
	// 微调的基本模型的名称。您可以选择“ada”，“babbage”，“curie”，“davinci”或在2022-04-21之后创建的微调模型之一。要了解有关这些模型的更多信息，请参阅模型文档。
	Model string `json:"model,omitempty"`

	// String (Optional) Defaults to 4.
	/**/
	// The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset.
	/**/
	// 训练模型的周期数。一个周期指的是通过训练数据集的一个完整周期。
	NEpochs int `json:"n_epochs,omitempty"`

	// Integer (Optional) Defaults to null.
	/**/
	// The batch size to use for training. The batch size is the number of training examples used to train a single forward and backward pass.
	// By default, the batch size will be dynamically configured to be ~0.2% of the number of examples in the training set, capped at 256 - in general, we've found that larger batch sizes tend to work better for larger datasets.
	/**/
	// 用于训练的批量大小。批量大小是用于训练单个正向和反向传递的训练示例数。
	// 默认情况下，批量大小将动态配置为训练集中示例数的~0.2%，上限为256 - 一般来说，我们发现对于较大的数据集，较大的批量大小往往效果更好。
	BatchSize *int `json:"batch_size,omitempty"`

	// Float (Optional) Defaults to null.
	/**/
	// The learning rate multiplier to use for training. The fine-tuning learning rate is the original learning rate used for pretraining multiplied by this value.
	// By default, the learning rate multiplier is the 0.05, 0.1, or 0.2 depending on final batch_size (larger learning rates tend to perform better with larger batch sizes). We recommend experimenting with values in the range 0.02 to 0.2 to see what produces the best results.
	/**/
	// 用于训练的学习率乘数。微调学习率是用于预训练的原始学习率乘以此值。
	// 默认情况下，学习率乘数是0.05，0.1或0.2，具体取决于最终的batch_size（较大的学习率往往在较大的批量大小下表现更好）。
	LearningRateMultiplier *float64 `json:"learning_rate_multiplier,omitempty"`

	// Float (Optional) Defaults to 0.01.
	/**/
	// The weight to use for loss on the prompt tokens. This controls how much the model tries to learn to generate the prompt (as compared to the completion which always has a weight of 1.0), and can add a stabilizing effect to training when completions are short.
	// If prompts are extremely long (relative to completions), it may make sense to reduce this weight so as to avoid over-prioritizing learning the prompt.
	/**/
	// 用于提示令牌损失的权重。这控制模型尝试学习生成提示的程度（与始终具有权重1.0的完成相比），并且当完成短时，可以为训练添加稳定效果。
	// 如果提示非常长（相对于完成），则可以减少此权重，以避免过度优先学习提示。
	PromptLossWeight *float64 `json:"prompt_loss_weight,omitempty"`

	// Boolean (Optional) Defaults to false.
	/**/
	// If set, we calculate classification-specific metrics such as accuracy and F-1 score using the validation set at the end of every epoch. These metrics can be viewed in the results file.
	// In order to compute classification metrics, you must provide a validation_file. Additionally, you must specify classification_n_classes for multiclass classification or classification_positive_class for binary classification.
	/**/
	// 如果设置，我们会在每个周期结束时使用验证集计算分类相关的指标，如准确率和F-1分数。这些指标可以在结果文件中查看。
	// 若要计算分类指标，您必须提供一个validation_file。此外，您必须为多分类任务指定classification_n_classes，或者为二分类任务指定classification_positive_class。
	ComputeClassificationMetrics bool `json:"compute_classification_metrics,omitempty"`

	// Integer (Optional) Defaults to null.
	/**/
	// The number of classes in a classification task.This parameter is required for multiclass classification.
	/**/
	// 分类任务中的类别数量。对于多类分类，此参数是必需的。
	ClassificationNClasses int `json:"classification_n_classes,omitempty"`

	// String (Optional) Defaults to null.
	/**/
	// The positive class in binary classification.
	// This parameter is needed to generate precision, recall, and F1 metrics when doing binary classification.
	/**/
	// 二元分类中的正类。
	// 在进行二元分类时，需要此参数来生成精确度、召回率和F1指标。
	ClassificationPositiveClass string `json:"classification_positive_class,omitempty"`

	// Array (Optional) Defaults to null.
	/**/
	// If this is provided, we calculate F-beta scores at the specified beta values. The F-beta score is a generalization of F-1 score. This is only used for binary classification.
	// With a beta of 1 (i.e. the F-1 score), precision and recall are given the same weight. A larger beta score puts more weight on recall and less on precision. A smaller beta score puts more weight on precision and less on recall.
	/**/
	// 如果提供了此参数，我们将计算指定beta值的F-beta分数。F-beta分数是F-1分数的泛化。这仅用于二元分类。
	// 当beta为1（即F-1分数）时，精确度和召回率具有相同的权重。较大的beta分数在召回率上赋予更大的权重，在精确度上赋予较小的权重。较小的beta分数在精确度上赋予更大的权重，在召回率上赋予较小的权重。
	ClassificationBetas []float64 `json:"classification_betas,omitempty"`

	// String (Optional) Defaults to null.
	/**/
	// A string of up to 40 characters that will be added to your fine-tuned model name.
	// For example, a suffix of "custom-model-name" would produce a model name like ada:ft-your-org:custom-model-name-2022-02-15-04-21-04.
	/**/
	// 最多40个字符的字符串，将添加到您的微调模型名称中。
	// 例如，后缀为“custom-model-name”的模型名称类似于ada:ft-your-org:custom-model-name-2022-02-15-04-21-04
	Suffix string `json:"suffix,omitempty"`
}

type FineTuneEvent struct {
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type TrainingFile struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type CreateFineTuneResult struct {
	ID              string                 `json:"id"`
	Object          string                 `json:"object"`
	Model           string                 `json:"model"`
	CreatedAt       int64                  `json:"created_at"`
	Events          []FineTuneEvent        `json:"events"`
	FineTunedModel  string                 `json:"fine_tuned_model"`
	Hyperparams     map[string]interface{} `json:"hyperparams"`
	OrganizationID  string                 `json:"organization_id"`
	ResultFiles     []interface{}          `json:"result_files"`
	Status          string                 `json:"status"`
	ValidationFiles []interface{}          `json:"validation_files"`
	TrainingFiles   []TrainingFile         `json:"training_files"`
	UpdatedAt       int64                  `json:"updated_at"`
	*Error
}

type ListFineTunesResult struct {
	Object string                 `json:"object"`
	Data   []CreateFineTuneResult `json:"data"`
	*Error
}

type ListFineTuneEventsResult struct {
	Object string          `json:"object"`
	Data   []FineTuneEvent `json:"data"`
	*Error
}

type DeleteFineTuneModelResult struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
	*Error
}

type ListFineTuneEventsRequest struct {
	FineTuneID string
	Stream     *bool
}

// CreateFineTune
// Manage fine-tuning jobs to tailor a model to your specific training data.
//
// Related guide: [Fine-tune models](https://platform.openai.com/docs/guides/fine-tuning)
/* 管理微调作业以使模型适合您的特定训练数据。*/
func (ft *FineTunes) CreateFineTune(req *CreateFineTuneRequest) (*CreateFineTuneResult, error) {
	var result CreateFineTuneResult
	err := ft.h.Df().Uri(createFineTuneURI).Method("POST").Json(req).Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

// ListFineTunes
// List your organization's fine-tuning jobs
/* 列出您的组织的微调作业*/
func (ft *FineTunes) ListFineTunes() (*ListFineTunesResult, error) {
	var result ListFineTunesResult
	err := ft.h.Df().Uri(listFineTunesURI).Method("GET").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

func (ft *FineTunes) RetrieveFineTune(fineTuneID string) (*CreateFineTuneResult, error) {
	var result CreateFineTuneResult
	err := ft.h.Df().Uri(fmt.Sprintf(retrieveFineTuneURI, fineTuneID)).Method("GET").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

func (ft *FineTunes) CancelFineTune(fineTuneID string) (*CreateFineTuneResult, error) {
	var result CreateFineTuneResult
	err := ft.h.Df().Uri(fmt.Sprintf(cancelFineTuneURI, fineTuneID)).Method("POST").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

func (ft *FineTunes) ListFineTuneEvents(req *ListFineTuneEventsRequest) (*ListFineTuneEventsResult, error) {
	var result ListFineTuneEventsResult
	df := ft.h.Df().Uri(fmt.Sprintf(fineTuneEventsURI, req.FineTuneID))
	if req.Stream != nil {
		df = df.Query("stream", strconv.FormatBool(*req.Stream))
	}
	err := df.Method("GET").Result(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ft *FineTunes) DeleteFineTuneModel(model string) (*DeleteFineTuneModelResult, error) {
	var result DeleteFineTuneModelResult
	err := ft.h.Df().Uri(fmt.Sprintf(deleteFineTuneModelURI, model)).Method("DELETE").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}
