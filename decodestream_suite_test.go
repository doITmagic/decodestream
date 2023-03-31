package decodestream_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDecodestream(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decodestream Suite")
}
