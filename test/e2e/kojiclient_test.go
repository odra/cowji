package test_e2e

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kojiclient "github.com/odra/kubeji/pkg/koji-client"
)

func TestKojiClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KojiClient")
}

var _ = Describe("KojiClient Suite", func() {
	Describe("Ping", func() {
		It("Should ping a koji-hub instance", func() {
			kc, kcError := kojiclient.New("https://koji.fedoraproject.org/kojihub")

			pingError := kc.Ping()

			Expect(kcError).To(BeNil())
			Expect(pingError).To(BeNil())

			kc.Defer()
		})
	})
})
