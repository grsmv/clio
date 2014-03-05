package core

import (
	"github.com/grsmv/clio/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("github.com/grsmv/clio/core/template.go", func() {
	Describe("IncludeStylesheets", func() {
		It("Should return generated list of stylesheet incudes", func() {
			stylesheets := []string{
				"stylesheets/application.css",
				"stylesheets/templates.css",
				"stylesheets/core.css",
			}
			Expect(core.IncludeStylesheets(stylesheets...)).To(Equal(""))
		})
	})
})
