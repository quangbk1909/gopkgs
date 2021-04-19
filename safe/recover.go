package safe

import (
	"fmt"

	"gitlab.id.vin/platform/gopkgs/internal/log"

	"go.uber.org/zap"
)

const (
	msgPanicRecovered = "recovered from panic"
)

var (
	defaultPanicHandler = func(e interface{}) {
		log.Logger().Error(msgPanicRecovered, zap.String("error", fmt.Sprint(e)))
	}
	panicHandler = defaultPanicHandler
)

func ReplacePanicHandler(f func(interface{})) {
	panicHandler = f
}

func WithRecover(f func()) {
	defer func() {
		if e := recover(); e != nil {
			if panicHandler != nil {
				panicHandler(e)
			}
		}
	}()
	f()
}

func WithRecoverError(f func() error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s: %v", msgPanicRecovered, e)
		}
	}()
	return f()
}
