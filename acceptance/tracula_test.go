package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

const THE_BAT = `
 (_   T R A C U L A   _)
  /\                 /\
 / \'._   (\_/)   _.'/ \
/_.''._'--('.')--'_.''._\
| \_ / ';=/ " \=;' \ _/ |
\/  '\__|'\___/'|__/'  \/
 '       \(/|\)/       '
          " ' "
`

var _ = Describe("Tracula", func() {
	It("version command", func() {
		command := exec.Command(pathToMain)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
		Expect(string(session.Out.Contents())).To(ContainSubstring(THE_BAT))
	})
})
