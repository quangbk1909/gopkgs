package workersgr

type Option interface {
	apply(settings *settings)
}

type optionFunc func(*settings)

func (f optionFunc) apply(settings *settings) {
	f(settings)
}

func OptionErrHandler(f func(error)) Option {
	return optionFunc(func(s *settings) {
		s.errHandler = f
	})
}

func OptionMaxGoroutines(n int) Option {
	return optionFunc(func(s *settings) {
		s.maxGoroutines = n
	})
}
