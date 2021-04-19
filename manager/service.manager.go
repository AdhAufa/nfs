package manager

type ServiceManager interface {
	Auth() string
}

func Auth() string {
	return "hello"
}
