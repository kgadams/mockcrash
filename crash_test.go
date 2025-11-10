package mockcrash

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mock "github.com/stretchr/testify/mock"
)

func TestCrash(t *testing.T) {
	mockPtr := NewWithPtr_Mock(t)
	mockPtr.EXPECT().Do(mock.Anything).Return(nil)
	mockPtr.Do(nil)
}

var _ = Describe("Crash", func() {
	It("test", func() {
		mockPtr := NewWithPtr_Mock(GinkgoT())
		mockPtr.EXPECT().Do(mock.Anything).Return(nil)
		Expect(mockPtr.Do(nil)).To(Succeed())
	})
})
