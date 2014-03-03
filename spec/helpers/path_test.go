package helpers

import (
	. "github.com/grsmv/clio/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("github.com/grsmv/clio/helpers/path.go", func() {
	Describe("FixPath", func() {
		It("Should fix given path, basing on OS type", func() {
			path := "/a/b/c"
			if os.PathSeparator == '/' {
				Expect(FixPath(path)).To(Equal("/a/b/c"))
			} else {
				Expect(FixPath(path)).To(Equal("\\a\\b\\c"))
			}
		})
	})
})
