package v1

import (
	"fmt"
	"github.com/artisancloud/httphelper/dataflow"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"

	"github.com/artisancloud/httphelper"
)

const (
	filesURI              = "/v1/files"
	fileContentURIFormat  = "/v1/files/%s/content"
	fileSpecificURIFormat = "/v1/files/%s"
)

type Files struct {
	h httphelper.Helper
}

type ListFilesResponse struct {
	Data   []FileInfo `json:"data"`
	Object string     `json:"object"`
	*Error
}

type FileInfo struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
	*Error
}

type DeleteFileResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
	*Error
}

func (f *Files) ListFiles() (*ListFilesResponse, error) {
	var result ListFilesResponse
	err := f.h.Df().Uri(filesURI).Method("GET").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

func (f *Files) UploadFile(file *os.File, purpose string) (*FileInfo, error) {
	var result FileInfo
	err := f.h.Df().Uri(filesURI).Method("POST").Multipart(func(multipart dataflow.MultipartDataflow) error {
		if file == nil {
			return errors.New("file is nil")
		}
		multipart.FileMem("file", filepath.Base(file.Name()), file)
		multipart.FieldValue("purpose", purpose)
		return nil
	}).Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

func (f *Files) DeleteFile(fileID string) (*DeleteFileResponse, error) {
	var result DeleteFileResponse
	err := f.h.Df().Uri(fmt.Sprintf(fileSpecificURIFormat, fileID)).Method("DELETE").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

func (f *Files) RetrieveFile(fileID string) (*FileInfo, error) {
	var result FileInfo
	err := f.h.Df().Uri(fmt.Sprintf(fileSpecificURIFormat, fileID)).Method("GET").Result(&result)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &result, nil
}

func (f *Files) RetrieveFileContent(fileID string, destination io.Writer) error {
	resp, err := f.h.Df().Uri(fmt.Sprintf(fileContentURIFormat, fileID)).Method("GET").Request()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(destination, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
