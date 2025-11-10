package mockcrash

import (
	"testing"

	mock "github.com/stretchr/testify/mock"
)

func TestCrash(t *testing.T) {
	mockPtr := NewWithPtr_Mock(t)
	mockPtr.EXPECT().Do(mock.Anything).Return(nil)
	mockPtr.Do(nil)
}
