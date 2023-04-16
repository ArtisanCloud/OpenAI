package openai

import (
	"github.com/artisancloud/httphelper"
	"github.com/artisancloud/httphelper/client"
	"github.com/artisancloud/httphelper/dataflow"
	v1 "github.com/artisancloud/openai/api/v1"

	"github.com/pkg/errors"
	"net/http"
)

type Config interface {
	Default()
	Validate() error
}

type V1Config struct {
	OpenAPIKey   string
	Organization string
	HttpDebug    bool
	ProxyURL     string
}

func (c *V1Config) Default() {
	return
}

func (c *V1Config) Validate() error {
	if c.OpenAPIKey == "" {
		return errors.New("openapi key is empty")
	}
	return nil
}

type ClientConfig struct {
	V1 V1Config
}

func (c *ClientConfig) Default() {
	return
}

func (c *ClientConfig) Validate() error {
	if c.V1 != (V1Config{}) {
		if c.V1.OpenAPIKey == "" {
			return errors.New("openapi key is empty")
		}
	}
	return nil
}

type Client struct {
	config Config
	*v1.V1
}

func NewClient(config Config) (*Client, error) {
	var clientConfig *ClientConfig
	switch c := config.(type) {
	case *ClientConfig:
		clientConfig = c
	case *V1Config:
		clientConfig = &ClientConfig{
			V1: *c,
		}
	}

	v1helperConfig := &httphelper.Config{
		BaseUrl: "https://api.openai.com",
		Config: &client.Config{
			ProxyURL: clientConfig.V1.ProxyURL,
		},
	}
	v1helper, err := httphelper.NewRequestHelper(v1helperConfig)
	v1helper.WithMiddleware(openAIV1AuthorizationMiddleware(clientConfig.V1), httphelper.HttpDebugMiddleware(clientConfig.V1.HttpDebug))
	if err != nil {
		return nil, err
	}

	c := Client{
		config: config,
		V1:     v1.NewV1(v1helper),
	}

	return &c, nil
}

func openAIV1AuthorizationMiddleware(config V1Config) dataflow.RequestMiddleware {
	return func(handle dataflow.RequestHandle) dataflow.RequestHandle {
		return func(request *http.Request, response *http.Response) error {
			request.Header.Add("Authorization", "Bearer "+config.OpenAPIKey)
			if config.Organization != "" {
				request.Header.Add("OpenAI-Organization", config.Organization)
			}
			return handle(request, response)
		}
	}
}
