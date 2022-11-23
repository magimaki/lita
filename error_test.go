package lita

import (
	"errors"
	"fmt"
	"testing"
)

func TestLitaError_Message(t *testing.T) {
	type fields struct {
		code    string
		message interface{}
		err     error
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test1",
			fields: fields{
				message: map[string]string{"err": "1"},
			},
		},
		{
			name: "test1",
			fields: fields{
				message: []string{"1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				code:    tt.fields.code,
				message: tt.fields.message,
				err:     tt.fields.err,
			}
			fmt.Println(e.Message())
		})
	}
}

func TestNewError(t *testing.T) {
	type args struct {
		code    string
		message interface{}
		err     error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				code:    ErrCmdInput,
				message: "error1",
			},
		},
		{
			name: "test2",
			args: args{
				code:    ErrCmdInput,
				message: "error2",
				err:     errors.New("failed to pass the test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewError(tt.args.code, tt.args.message, tt.args.err)
			fmt.Println(err.Error())
		})
	}
}
