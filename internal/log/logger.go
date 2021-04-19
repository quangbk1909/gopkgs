package log

import (
	"fmt"

	"go.uber.org/zap"
)

var (
	logger = zap.NewNop()
)

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Println("zapinit pkg: cannot initialize zap logger, used default noop logger instead")
	}
}

func ReplaceNoop() {
	logger = zap.NewNop()
}

func Logger() *zap.Logger {
	return logger
}
