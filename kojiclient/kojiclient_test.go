package kojiclient

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKojiClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KojiClient")
}

var _ = Describe("KojiClient Suite", func() {
	Describe("New instance", func() {
		It("Should create a KojiClient instance", func() {
			c, err := New("https://koji.hub")

			Expect(err).To(BeNil())
			Expect(c.Url).To(Equal("https://koji.hub"))
			Expect(c.api).ToNot(BeNil())

			c.Defer()
		})

		It("Should fail using an invalid url", func() {
			c, err := New(":koji")

			Expect(err).ToNot(BeNil())
			Expect(c).To(BeNil())
		})

		It("Should fail using an empty url", func() {
			c, err := New("")

			Expect(err).To(BeNil())
			Expect(c).To(BeNil())
		})
	})
})
