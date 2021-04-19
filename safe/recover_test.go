package safe

import (
	"errors"
	"testing"

	"gitlab.id.vin/platform/gopkgs/internal/log"

	"github.com/stretchr/testify/mock"
)

type panicHandlerMock struct {
	mock.Mock
}

func (m *panicHandlerMock) handle(e interface{}) {
	m.Called(e)
}

func init() {
	log.ReplaceNoop()
	panicHandlerMock := &panicHandlerMock{}
	panicHandlerMock.On("handle", mock.Anything).Return()
	ReplacePanicHandler(panicHandlerMock.handle)
}

func TestDefaultPanicHandler(t *testing.T) {
	defaultPanicHandler(nil)
	defaultPanicHandler(errors.New("error"))
	defaultPanicHandler("test")
}

func TestWithRecover(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "call panic explicitly",
			args: args{
				f: func() {
					panic("")
				},
			},
		},
		{
			name: "call panic implicitly",
			args: args{
				f: func() {
					var x *int = nil
					_ = *x / 1
				},
			},
		},
		{
			name: "no panic",
			args: args{
				f: func() {
					_ = 100
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			WithRecover(args.f)
		})
	}
}

func TestWithRecoverError(t *testing.T) {
	type args struct {
		f func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "call panic explicitly",
			args: args{
				f: func() error {
					panic("")
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "call panic implicitly",
			args: args{
				f: func() error {
					var x *int = nil
					_ = *x / 1
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "no panic no error",
			args: args{
				f: func() error {
					_ = 100
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "no panic but func return error",
			args: args{
				f: func() error {
					_ = 100
					return errors.New("any")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WithRecoverError(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("WithRecoverError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
