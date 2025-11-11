package mockcrash

import (
	"context"

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
	It("templ", func() {
		mockPtr := NewWithPtr_Mock(GinkgoT())
		val := Value{
			Parameter: "anything",
		}
		mockPtr.EXPECT().Do(mock.Anything).Return(val, nil)
		ctx := context.WithValue(context.Background(), MockKey, mockPtr)
		got, err := Execute(ctx, `{{ doit }}`, map[string]any{})
		Expect(err).ToNot(HaveOccurred())
		Expect(got).To(Equal("Value(anything)"))
	})
})
