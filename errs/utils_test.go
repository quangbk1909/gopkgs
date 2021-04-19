package errs

import (
	"errors"
	"testing"
)

func TestAny(t *testing.T) {
	type args struct {
		errors []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "has error",
			args:    args{[]error{errors.New("sample")}},
			wantErr: true,
		},
		{
			name:    "empty error list",
			args:    args{[]error{}},
			wantErr: false,
		},
		{
			name:    "empty slice error",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Any(tt.args.errors...); (err != nil) != tt.wantErr {
				t.Errorf("Any() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
