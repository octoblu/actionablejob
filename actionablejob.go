package actionablejob

import (
  "strconv"
  "time"
  "github.com/garyburd/redigo/redis"
)

// Job represents a job in the queue
type Job interface {
  GetKey() string
}

// ActionableJob represents a job in the queue that can know if it should have action taken upon it
type ActionableJob interface {
  GetKey() string
  Claim() (bool,error)
}

// ActionableRedisJob implements ActionableJob and stores/retrieves jobs from Redis
type ActionableRedisJob struct {
  key string
}

// New returns a new job
func New(key string) *ActionableRedisJob {
  return &ActionableRedisJob{key: key}
}

// NewFromJob returns a new job based on a Job
func NewFromJob(job Job) *ActionableRedisJob {
  return &ActionableRedisJob{key: job.GetKey()}
}

// GetKey returns the key to store the information about the job
func (job *ActionableRedisJob) GetKey() string {
  return job.key
}

// Claim returns true when the caller succesfully claims the job
func (job *ActionableRedisJob) Claim() (bool,error) {
  var err error
  var redisConn redis.Conn
  var result interface{}

	redisConn, err = redis.Dial("tcp", ":6379")

  if err != nil {
    return false, err
  }

  now  := time.Now().Unix()
  then := now + 1
  result,err = redisConn.Do("GETSET", "[namespace]-faulty", then)

  nextTick := parseNextTick(result)

	if now < nextTick {
		return false, nil
	}

  return true, nil
}

func parseNextTick(redisResult interface{}) int64 {
  strNextTick,ok := redisResult.([]uint8)
  if !ok {
    return 0
  }

  nextTick,err := strconv.ParseInt(string(strNextTick), 10, 64)
  if err != nil {
    return 0
  }

  return nextTick
}
