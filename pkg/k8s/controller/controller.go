package controller

type Controller struct {
}

func New() (*Controller, error) {
	return &Controller{}, nil
}
