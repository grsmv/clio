package core

import (
	"github.com/grsmv/clio/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("github.com/grsmv/clio/core/template.go", func() {
	Describe("IncludeStylesheets", func() {
		It("Should return generated list of stylesheet incudes", func() {
			files := []string{
				"stylesheets/application",
				"stylesheets/templates",
				"stylesheets/core",
			}
			fixture := " <link rel=\"stylesheet\" href=\"/assets/stylesheets/application.css\" type=\"text/css\" media=\"screen\" charset=\"utf-8\">\n <link rel=\"stylesheet\" href=\"/assets/stylesheets/templates.css\" type=\"text/css\" media=\"screen\" charset=\"utf-8\">\n <link rel=\"stylesheet\" href=\"/assets/stylesheets/core.css\" type=\"text/css\" media=\"screen\" charset=\"utf-8\">\n"
			Expect(core.IncludeStylesheets(files...)).To(Equal(fixture))
		})

		It("generate link to a single stylesheet", func() {
			Expect(core.IncludeStylesheets("css/styles")).To(Equal(
				" <link rel=\"stylesheet\" href=\"/assets/css/styles.css\" type=\"text/css\" media=\"screen\" charset=\"utf-8\">\n",
			))
		})
	})

	Describe("IncludeJavascripts", func() {
		It("Should return generated list of javascript incudes", func() {
			files := []string{
				"js/application",
				"js/templates",
			}
			fixture := " <script type=\"text/javascript\" src=\"/assets/js/application.js\"></script>\n <script type=\"text/javascript\" src=\"/assets/js/templates.js\"></script>\n"
			Expect(core.IncludeJavascripts(files...)).To(Equal(fixture))
		})

		It("generate link to a single javascript file", func() {
			Expect(core.IncludeJavascripts("js/app")).To(Equal(
				" <script type=\"text/javascript\" src=\"/assets/js/app.js\"></script>\n",
			))
		})
	})
})
