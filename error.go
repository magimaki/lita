package lita

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	ErrCmdInput = "error command input"
)

type Error struct {
	code    string
	message interface{}
	err     error
}

func NewError(code string, message interface{}, err error) error {
	return &Error{
		code:    code,
		message: message,
		err:     err,
	}
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Message() interface{} {
	return e.message
}

func (e *Error) Error() string {
	if e.err == nil {
		return fmt.Sprintf("[%s] %s: %s", strings.ToUpper(Lita), e.code, e.message)
	}
	return fmt.Sprintf("[%s] %s: %s: %s", strings.ToUpper(Lita), e.code, e.message, e.err.Error())
}

func (e *Error) MarshalJSON() ([]byte, error) {
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
