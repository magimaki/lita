package error

import (
	"encoding/json"
	"fmt"
)

var (
	ErrCmdInput = "error command input"
)

type LitaError struct {
	code    string
	message interface{}
	err     error
}

func NewError(code string, message interface{}, err error) error {
	return &LitaError{
		code:    code,
		message: message,
		err:     err,
	}
}

func (e *LitaError) Code() string {
	return e.code
}

func (e *LitaError) Message() interface{} {
	return e.message
}

func (e *LitaError) Error() string {
	if e.err == nil {
		return fmt.Sprintf("[Lita] %s: %s", e.code, e.message)
	}
	return fmt.Sprintf("[Lita] %s: %s: %s", e.code, e.message, e.err.Error())
}

func (e *LitaError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code    string      `json:"code"`
		Message interface{} `json:"message"`
		Err     error       `json:"err"`
	}{
		Code:    e.code,
		Message: e.message,
		Err:     e.err,
	})
}
