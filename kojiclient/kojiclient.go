package kojiclient

import (
	"net/url"
	"strings"

	"github.com/odra/cowji/kojiclient/apis/xmlrpc"
	"github.com/odra/cowji/kojiclient/meta"
)

// A Koji API client struct, used to consume methods from Koji Hub.
// The *Url* argument should be the endpoint of the XML RPC API,
// such as *https://koji.fedoraproject.org/kojihub/*.
type KojiClient struct {
	Url string
	api meta.APIClientSpec
}

// Create a new KojiClient instance.
// The *Url* argument should be the endpoint of the XML RPC API,
// such as *https://koji.fedoraproject.org/kojihub/*.
// It will try to parse the provided url and return an
// error in case it is not able to.
func New(_url string) (*KojiClient, error) {
	trimmed := strings.TrimSpace(_url)
	u, err := url.Parse(trimmed)
	if u == nil || err != nil || strings.Trim(trimmed, " ") == "" {
		return nil, err
	}

	xmlClient, err := xmlrpc.FromURL(_url)
	if err != nil {
		return nil, err
	}

	return &KojiClient{
		Url: _url,
		api: xmlClient,
	}, nil
}

func (kc *KojiClient) Ping() error {
	return kc.api.Ping()
}

func (kc *KojiClient) Defer() error {
	return kc.api.Defer()
}
