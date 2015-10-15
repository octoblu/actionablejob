package actionablejob

// Job represents a job in the queue
type Job interface {
  GetKey() string
}

// ActionableJob represents a job in the queue that can know if it should have action taken upon it
type ActionableJob interface {
  GetKey() string
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
