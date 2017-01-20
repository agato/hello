package greet

import (
	. "github.com/r7kamura/gospel"
	"testing"

)

func TestDescribe(t *testing.T) {
	Describe(t, "hello", func() {
		Context("say hello", func() {
			It("should be 'hello'", func() {
				Expect(Speak()).To(Equal, "Hello World")
			})
		})
	})
}
