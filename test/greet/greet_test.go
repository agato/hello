package greet

import (
	. "github.com/r7kamura/gospel"
	"testing"

	"github.com/agato/hello/greet"
)

func TestDescribe(t *testing.T) {
	Describe(t, "hello", func() {
		Context("say hello", func() {
			It("should be 'hello'", func() {
				Expect(greet.Speak()).To(Equal, "hello")
			})
		})
	})
}
