package actionablejob_test

import (
	"github.com/octoblu/actionablejob"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DumbJob struct {}
func (dumbJob *DumbJob) GetKey() string {
	return "I am dumb"
}

var _ = Describe("ActionableRedisJob", func() {
	Describe("New", func() {
		Context("called with a key", func(){
			var sut actionablejob.ActionableJob

			BeforeEach(func(){
				sut = actionablejob.New("old-map")
			})

			It("should set the key", func(){
				Expect(sut.GetKey()).To(Equal("old-map"))
			})
		})
	})

	Describe("NewFromJob", func() {
		Context("called with a job", func(){
			var sut actionablejob.ActionableJob

			BeforeEach(func(){
				job := new(DumbJob)
				sut = actionablejob.NewFromJob(job)
			})

			It("should set the key", func(){
				Expect(sut.GetKey()).To(Equal("I am dumb"))
			})
		})
	})
})
