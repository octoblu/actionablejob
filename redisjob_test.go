package redisjob_test

import (
	"github.com/octoblu/redisjob"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DumbJob struct {}
func (dumbJob *DumbJob) GetKey() string {
	return "I am dumb"
}

var _ = Describe("Redisjob", func() {
	Describe("New", func() {
		Context("called with a key", func(){
			var sut redisjob.Job

			BeforeEach(func(){
				sut = redisjob.New("old-map")
			})

			It("should set the key", func(){
				Expect(sut.GetKey()).To(Equal("old-map"))
			})
		})
	})

	Describe("NewFromJob", func() {
		Context("called with a job", func(){
			var sut redisjob.Job

			BeforeEach(func(){
				job := new(DumbJob)
				sut = redisjob.NewFromJob(job)
			})

			It("should set the key", func(){
				Expect(sut.GetKey()).To(Equal("I am dumb"))
			})
		})
	})
})
