package xmlrpc

import (
	xmlrpcClient "alexejk.io/go-xmlrpc"
	"errors"
	"github.com/odra/kubeji/pkg/koji-client/meta"
)

type caller interface {
	Call(method string, args any, reply any) error
	Close() error
}

type Client struct {
	client caller
}

func New(c caller) *Client {
	return &Client{
		client: c,
	}
}

func FromURL(url string) (*Client, error) {
	client, err := xmlrpcClient.NewClient(url)
	if err != nil {
		return nil, err
	}

	return New(client), nil
}

func (c *Client) Ping() error {
	methods := meta.Methods{}

	err := c.client.Call("system.listMethods", nil, &methods)
	if err != nil {
		return err
	}

	if len(methods.Data) == 0 {
		return errors.New("system.listMethods returned an empty list")
	}

	return nil
}

func (c *Client) Defer() error {
	var err error

	defer func() {
		err = c.client.Close()
	}()

	return err
}
