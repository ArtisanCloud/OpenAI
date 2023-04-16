## OpenAI Golang SDK

The OpenAI Golang SDK is a Golang library for interacting with the OpenAI API.

Related documentation: [OpenAI API Documentation](https://platform.openai.com/docs/api-reference)

### Installation

You can easily install the OpenAI Golang SDK using Go Modules. In your project directory, run the following command:

arduinoCopy code

`go get -u github.com/artisancloud/openai` 

### Usage

Here is a simple example that demonstrates how to use the SDK to create an API client and generate a piece of text:

```
func main() {
	// Create an OpenAI API client
	config  := openai.V1Config{
		OpenAPIKey: "your-openai-key",     // Required, your OpenAI API key
		Organization: "your-organization", // Optional, your organization ID
		HttpDebug:  true,                  // Optional, enable debug mode (outputs HTTP request information using default log)
		ProxyURL: "proxy-url",             // Optional, proxy address, e.g. http://xxxxxx:port, if including basic authentication, format is http://username:password@xxxxxx:port
	}
	client, err := openai.NewClient(&config)
	if err != nil {
		panic(err)
	}

	// Create ChatCompletion
	req1 := v1.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []v1.Message{
			{
				Role:    "user",
				Content: "Hello!",
			},
		},
	}
	result1, err1 := client.V1.Chat.CreateChatCompletion(&req1)
	if err1 != nil {
		if e, ok := err1.(*v1.Error); ok {
			fmt.Printf("open ai error: %v", e.Detail.Message)
		}
		fmt.Printf("error: %v", err1)
		return
	}

	// Print the result
	fmt.Println(result1)

	// Create Completion
	req2 := v1.CreateCompletionRequest{
		Model:       "text-davinci-003",
		Prompt:      openaitype.StringOrArray{"Say this is a test", "Say Hello World"},
		MaxTokens:   openaitype.PointInt(150),
		Temperature: openaitype.PointFloat64(0),
		TopP:        openaitype.PointFloat64(1),
		N:           openaitype.PointInt(1),
		Stream:      openaitype.PointBool(false),
		Stop:        openaitype.StringOrArray{"\n"},
	}

	result2, err2 := client.V1.Completion.CreateCompletion(&req2)
	if err2 != nil {
		if e, ok := err2.(*v1.Error); ok {
			fmt.Printf("open ai error: %v", e.Detail.Message)
		}
		fmt.Printf("error: %v", err2)
		return
	}

	// Print the result
	fmt.Println(result2)
}
```

### Handling Errors

If you encounter an error while making an HTTP request using the SDK or if OpenAI returns an error (see `openai.api.v1.Error`), these errors will be returned as the `error` argument. If you need to operate on the error details returned by OpenAI, you can use `result.Error` or try type assertion `err.(*v1.Error)`.

### Additional Type Definitions

In addition to the OpenAI Golang SDK itself, some type definitions are provided to accommodate optional fields in OpenAI. These type definitions and utility functions can be found in the `openaitype` package. These type definitions and functions can help you interact more conveniently with the OpenAI API and ensure the API request format is correct.

Here are the definitions in the `openaitype` package:

- `StringOrArray` type: A string or array of strings type that implements `MarshalJSON` and `UnmarshalJSON` methods for serialization and deserialization between JSON and `StringOrArray`.
    
- Conversion functions: These functions are used to convert basic types to their corresponding pointer types:
    
    - `PointInt(i int) *int`
    - `PointInt32(i int32) *int32`
    - `PointInt64(i int64) *int64`
    - `PointBool(b bool) *bool`
    - `PointFloat32(f float32) *float32`
    - `PointFloat64(f float64) *float64`
    - `PointString(s string) *string`

### Contributions

Powered by [Artisan Cloud](https://github.com/ArtisanCloud)

- [Northseadl](https://github.com/northseadl)

Documentation generated by ChatGPT