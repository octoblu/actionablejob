package redisjob

// Job represents a job in the queue
type Job interface {
  GetKey() string
}

// RedisJob implements Job and stores/retrieves jobs from Redis
type RedisJob struct {
  key string
}

// NewJob returns a new job
func NewJob(key string) *RedisJob {
  return &RedisJob{key: key}
}

// GetKey returns the key to store the information about the job
func (job *RedisJob) GetKey() string {
  return job.key
}
