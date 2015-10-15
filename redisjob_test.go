package redisjob_test

import (
	"github.com/octoblu/redisjob"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redisjob", func() {
	Context("New with a key", func() {
		var sut redisjob.Job

		BeforeEach(func(){
			sut = redisjob.New("old-map")
		})

		It("should set the key", func(){
			Expect(sut.GetKey()).To(Equal("old-map"))
		})
	})
})
