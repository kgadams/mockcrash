package mockcrash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mock "github.com/stretchr/testify/mock"
)

var _ = Describe("Crash", func() {
	It("test", func() {
		mockPtr := NewWithPtr_Mock(GinkgoT())
		val := Value{
			Parameter: "anything",
		}
		mockPtr.EXPECT().Do(mock.Anything).Return(val, nil)
		got, err := mockPtr.Do(&Value{})
		Expect(err).ToNot(HaveOccurred())
		Expect(got).To(Equal(val))
	})
})
