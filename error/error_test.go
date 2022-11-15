package error

import (
	"errors"
	"fmt"
	"reflect"
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
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &LitaError{
				code:    tt.fields.code,
				message: tt.fields.message,
				err:     tt.fields.err,
			}
			if got := e.Message(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message() = %v, want %v", got, tt.want)
			}
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
