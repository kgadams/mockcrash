package mockcrash_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/types"
	. "github.com/onsi/gomega"
)

func TestMockcrash(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mockcrash Suite", types.ReporterConfig{
		NoColor: true,
	})
}
