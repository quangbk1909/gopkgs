package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	base = errors.New("base error")
)

func TestCause(t *testing.T) {
	err1 := WithStack(base)
	err2 := fmt.Errorf("text2: %w", err1)
	err3 := WithStack(err2)
	err4 := fmt.Errorf("text3: %w", err3)
	assert.Equal(t, base, Cause(err4))
}

func TestWithStack(t *testing.T) {
	err1 := WithStack(base)
	err2 := fmt.Errorf("text2: %w", err1)
	err3 := WithStack(err2)
	err4 := fmt.Errorf("text3: %w", err3)
	assert.Equal(t, "text3: text2: base error", err4.Error())
}

func TestWrap(t *testing.T) {
	err1 := WithStack(base)
	err2 := Wrap(err1, "text2")
	err3 := Wrap(err2, "text3")
	err4 := Wrap(err3, "text4")
	assert.Equal(t, "text4: text3: text2: base error", err4.Error())
	assert.Equal(t, base, Cause(err4))
	var stack1, stack2 interface {
		StackTrace() string
	}
	ok := errors.As(err1, &stack1)
	assert.Equal(t, true, ok)
	ok = errors.As(err4, &stack2)
	assert.Equal(t, true, ok)
	assert.Equal(t, stack2.StackTrace(), stack1.StackTrace())
}

func TestIsError(t *testing.T) {
	err1 := WithStack(base)
	err2 := Wrap(err1, "text1")
	err3 := fmt.Errorf("text2: %w", err2)
	assert.Equal(t, true, errors.Is(err3, base))
	assert.Equal(t, false, errors.Is(err3, errors.New("something")))
}

func Test_errorStack_StackTrace(t *testing.T) {
	type fields struct {
		error      error
		stackTrace string
		caller     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "happy case",
			fields: fields{
				error:      nil,
				stackTrace: "stack trace",
				caller:     "",
			},
			want: "stack trace",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorStack{
				error:      tt.fields.error,
				stackTrace: tt.fields.stackTrace,
				caller:     tt.fields.caller,
			}
			if got := e.StackTrace(); got != tt.want {
				t.Errorf("StackTrace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorStack_Caller(t *testing.T) {
	type fields struct {
		error      error
		stackTrace string
		caller     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "happy case",
			fields: fields{
				error:      nil,
				stackTrace: "",
				caller:     "caller",
			},
			want: "caller",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errorStack{
				error:      tt.fields.error,
				stackTrace: tt.fields.stackTrace,
				caller:     tt.fields.caller,
			}
			if got := e.Caller(); got != tt.want {
				t.Errorf("Caller() = %v, want %v", got, tt.want)
			}
		})
	}
}
