package redisjob

// Job represents a job in the queue
type Job interface {
  GetKey() string
}

// RedisJob implements Job and stores/retrieves jobs from Redis
type RedisJob struct {
  key string
}

// New returns a new job
func New(key string) *RedisJob {
  return &RedisJob{key: key}
}

// NewFromJob returns a new job based on a Job
func NewFromJob(job Job) *RedisJob {
  return &RedisJob{key: job.GetKey()}
}

// GetKey returns the key to store the information about the job
func (job *RedisJob) GetKey() string {
  return job.key
}
