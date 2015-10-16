package claimablejob

import (
  "fmt"
  "strconv"
  "time"
  "github.com/garyburd/redigo/redis"
)

// Job represents a job in the queue
type Job interface {
  GetKey() string
}

// Conn is what all claimable jobs require
type Conn interface {
	Do(commandName string, args... interface{}) (interface{}, error)
}

// ClaimableJob represents a job in the queue that can know if it should have action taken upon it
type ClaimableJob interface {
  Claim() (bool,error)
  GetKey() string
  PushKeyIntoQueue(name string) error
}

// ClaimableRedisJob implements ClaimableJob and stores/retrieves jobs from Redis
type ClaimableRedisJob struct {
  key string
  conn Conn
}

// New returns a new job
func New(key string, conn Conn) *ClaimableRedisJob {
  return &ClaimableRedisJob{key: key, conn: conn}
}

// NewFromJob returns a new job based on a Job
func NewFromJob(job Job, conn Conn) *ClaimableRedisJob {
  return &ClaimableRedisJob{key: job.GetKey(), conn: conn}
}

// Claim returns true when the caller succesfully claims the job
func (job *ClaimableRedisJob) Claim() (bool,error) {
  now  := time.Now().Unix()
  then := now + 1

  result,err := job.conn.Do("GETSET", job.tickKey(), then)
  if err != nil {
    return false, err
  }

  nextTick := parseNextTick(result)

	if now < nextTick {
		return false, nil
	}

  return true, nil
}

// GetKey returns the key to store the information about the job
func (job *ClaimableRedisJob) GetKey() string {
  return job.key
}

// PushKeyIntoQueue pushes a key into a queue
func (job *ClaimableRedisJob) PushKeyIntoQueue(queueName string) error {
  var redisConn redis.Conn
  var err error

  redisConn, err = redis.Dial("tcp", ":6379")

  if err != nil {
    return err
  }

  _,err = redisConn.Do("LPUSH", queueName, job.GetKey())
  return nil
}

func (job *ClaimableRedisJob) tickKey() string {
  return fmt.Sprintf("[namespace]-%s", job.GetKey())
}

func parseNextTick(redisResult interface{}) int64 {
  strNextTick, ok  := redisResult.([]uint8)
  if !ok {
    return 0
  }

  nextTick,err := strconv.ParseInt(string(strNextTick), 10, 64)
  if err != nil {
    return 0
  }

  return nextTick
}
