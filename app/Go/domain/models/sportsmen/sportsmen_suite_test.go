package sportsmen_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSportsmen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sportsmen Suite")
}
