package v1

import "fmt"

type Error struct {
	Detail ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string      `json:"code"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"`
	Message string      `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %s, type: %s, param: %v, message: %s", e.Detail.Code, e.Detail.Type, e.Detail.Param, e.Detail.Message)
}
