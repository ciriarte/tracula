package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var pathToMain string

func TestTracula(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tracula Suite")
}

var _ = BeforeSuite(func() {
	var err error
	pathToMain, err = gexec.Build("github.com/ciriarte/tracula")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
