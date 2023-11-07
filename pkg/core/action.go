package limiter

type Action interface {
	Execute() error
}

type ActionFunc func() error

func (af ActionFunc) Execute() error {
	return af()
}
