package xmlrpc

import (
	"errors"
	"testing"

	"github.com/odra/cowji/pkg/koji-client/meta"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// fake xmlrpc client
type fakeCall func(reply any) error

type fakeClient struct {
	fakeCalls map[string]fakeCall
}

func (fc *fakeClient) Call(method string, args any, reply any) error {
	if handler, ok := fc.fakeCalls[method]; ok {
		return handler(reply)
	}

	return nil
}

func (fc *fakeClient) Close() error {
	return nil
}

func TestKojiClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XMLRPC Client Suite")
}

var _ = Describe("XMLRPC Client Suite", func() {
	Describe("Default Instance", func() {
		It("Should create a new deafult client", func() {
			client, err := FromURL("koji-hub:8443")

			Expect(client).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		It("Should fail to create a new default client", func() {
			client, err := FromURL(":koji-hub")

			Expect(client).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("Ping", func() {
		It("Should ping the remote server", func() {
			m := fakeClient{
				fakeCalls: map[string]fakeCall{
					"system.listMethods": func(reply any) error {
						obj := reply.(*meta.Methods)
						obj.Data = append(obj.Data, "aaa")

						return nil
					},
				},
			}
			client := New(&m)
			err := client.Ping()

			Expect(err).To(BeNil())
		})

		It("Should fail to ping the remote server", func() {
			m := fakeClient{
				fakeCalls: map[string]fakeCall{
					"system.listMethods": func(reply any) error {
						return errors.New("internal server error")
					},
				},
			}
			client := New(&m)
			err := client.Ping()

			Expect(err).ToNot(BeNil())
		})

		It("Should fail to retrieve rpc methods", func() {
			m := fakeClient{
				fakeCalls: map[string]fakeCall{
					"system.listMethods": func(reply any) error {
						obj := reply.(*meta.Methods)
						obj.Data = []string{}

						return nil
					},
				},
			}
			client := New(&m)
			err := client.Ping()

			Expect(err).ToNot(BeNil())
		})
	})
})
