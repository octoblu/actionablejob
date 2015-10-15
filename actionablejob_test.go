package actionablejob_test

import (
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/octoblu/actionablejob"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DumbJob struct {}
func (dumbJob *DumbJob) GetKey() string {
	return "I am dumb"
}

var _ = Describe("ActionableRedisJob", func() {
	var redisConn redis.Conn
	var sut actionablejob.ActionableJob

	BeforeEach(func() {
		var err error

		redisConn, err = redis.Dial("tcp", ":6379")
		Expect(err).To(BeNil())
	})

	AfterEach(func(){
		redisConn.Close()
	})

	Describe("New", func() {
		Context("called with a key", func(){
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
			BeforeEach(func(){
				job := new(DumbJob)
				sut = actionablejob.NewFromJob(job)
			})

			It("should set the key", func(){
				Expect(sut.GetKey()).To(Equal("I am dumb"))
			})
		})
	})

	Describe("Claim", func(){
		BeforeEach(func(){
			sut = actionablejob.New("faulty")
		})

		Context("When the job has already run this second", func(){
			BeforeEach(func(){
				then := int64(time.Now().Unix() + 1)
				_,err := redisConn.Do("SET", "[namespace]-faulty", then)
				Expect(err).To(BeNil())
			})

			It("should return false", func(){
				Expect(sut.Claim()).To(BeFalse())
			})
		})

		Context("When the job ran in the previous second", func(){
			var gotClaim bool

			BeforeEach(func(){
				now := int64(time.Now().Unix())
				_,err := redisConn.Do("SET", "[namespace]-faulty", now)
				Expect(err).To(BeNil())

				gotClaim,err = sut.Claim()
			})

			It("should return false", func(){
				Expect(gotClaim).To(BeTrue())
			})

			It("should advance the time", func(){
				result,err := redisConn.Do("GET", "[namespace]-faulty")
				Expect(err).To(BeNil())

				nextTickStr := string(result.([]uint8))
			  nextTick,err := strconv.ParseInt(nextTickStr, 10, 64)
				Expect(err).To(BeNil())

				then := int64(time.Now().Unix()) + 1
				Expect(nextTick).To(Equal(then))
			})
		})
	})
})
