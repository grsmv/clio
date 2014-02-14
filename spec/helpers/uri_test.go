package helpers

import (
    . "github.com/grsmv/clio/helpers"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("github.com/grsmv/clio/helpers/uri.go", func (){

    Describe("ParseQuery", func () {

        It("Should parse given paramenters into map", func (){
            Expect(ParseQuery("a=b&c=d")).To(Equal(map[string]string{"a": "b", "c": "d"}))
        })

        It("Should pair unmatched key with empty string", func (){
            Expect(ParseQuery("a=b&c")).To(Equal(map[string]string{"a":"b", "c": ""}))
        })

        It("Should split main path from parameters", func (){
            abs, params := SplitPath("/a/b/c?a=b&c")
            Expect(abs).To(Equal("/a/b/c"))
            Expect(params).To(Equal("a=b&c"))
        })

        It("Shoud do nothing if main url followed by question mark", func () {
            _, params := SplitPath("/a/b/c?")
            Expect(params).To(Equal(""))
        })

        It("Should produce empty absolute path and params if empty string given", func () {
            abs, params := SplitPath("")
            Expect(abs).To(Equal(""))
            Expect(params).To(Equal(""))
        })
    })

    Describe("PreparePattern", func () {

        It("Should prepare correct pattern for splat parsing", func () {
            Expect(PreparePattern("/users/:id/edit").String()).To(Equal("^/users/(?P<id>[\\p{L}\\d-_]{1,})/edit$"))
        })

        It("Should escape dots in given path", func () {
            Expect(PreparePattern("/a.mp3").String()).To(Equal("^/a\\.mp3$"))
        })
    })
})
