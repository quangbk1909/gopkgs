package shutdown

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"syscall"
	"testing"
)

func init() {
	logger = zap.NewNop()
}

func TestSigtermHandler(t *testing.T) {
	type ExpectedCall struct {
		cleanFunc    int
		cleanErrFunc int
	}
	tests := []struct {
		name              string
		cleaningFunctions *cleaningFuncMock
		numSignals        int
		expectedCall      ExpectedCall
	}{
		{
			name:              "normal",
			cleaningFunctions: newCleaningFuncMockNoError(),
			numSignals:        1,
			expectedCall:      ExpectedCall{cleanFunc: 1, cleanErrFunc: 1},
		},
		{
			name:              "clean error func returns non-nil error",
			cleaningFunctions: newCleaningFuncMockError(),
			numSignals:        1,
			expectedCall:      ExpectedCall{cleanFunc: 1, cleanErrFunc: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := initSigtermHandler()
			handler.RegisterFunc(tt.cleaningFunctions.Clean)
			handler.RegisterErrorFunc(tt.cleaningFunctions.CleanErr)
			for i := 0; i < tt.numSignals; i++ {
				handler.sigChannel <- syscall.SIGTERM
			}
			handler.Wait()
			tt.cleaningFunctions.AssertNumberOfCalls(t, "Clean", tt.expectedCall.cleanFunc)
			tt.cleaningFunctions.AssertNumberOfCalls(t, "CleanErr", tt.expectedCall.cleanErrFunc)
		})
	}
}

func newCleaningFuncMockError() *cleaningFuncMock {
	m := &cleaningFuncMock{}
	m.On("Clean")
	m.On("CleanErr").Return(errors.New("some error"))
	return m
}

func newCleaningFuncMockNoError() *cleaningFuncMock {
	m := &cleaningFuncMock{}
	m.On("Clean")
	m.On("CleanErr").Return(nil)
	return m
}

type cleaningFuncMock struct {
	mock.Mock
}

func (c *cleaningFuncMock) Clean() {
	c.Called()
}

func (c *cleaningFuncMock) CleanErr() error {
	args := c.Called()
	return args.Error(0)
}
