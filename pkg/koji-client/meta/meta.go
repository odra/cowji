package meta

type APIClientSpec interface {
	Ping() error
	Defer() error
}

type Methods struct {
	Data []string
}
