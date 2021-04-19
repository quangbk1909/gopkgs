package timex

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	t, err := time.Parse("2006-01-02", "2019-10-19")
	if err != nil {
		panic(err)
	}
	now = func() time.Time {
		return t
	}
}

func TestGetCurrentMillisecond(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "", want: 1571443200000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CurrentUnixMillisecond())
		})
	}
}

func TestGetTimeMillisecond(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			want: 1571443200000,
			args: args{
				t: now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			assert.Equal(t, tt.want, ToUnixMillisecond(args.t))
		})
	}
}

func TestFromUnixMillisecond(t *testing.T) {
	type args struct {
		milli int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "without nano second",
			want: time.Unix(1608611075, 0),
			args: args{
				milli: 1608611075000,
			},
		},
		{
			name: "with nano second",
			want: time.Unix(1608611075, 234000000),
			args: args{
				milli: 1608611075234,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			assert.Equal(t, tt.want, FromUnixMillisecond(args.milli))
		})
	}
}
