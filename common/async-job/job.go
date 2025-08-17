package asyncjob

import (
	"context"
	"time"
)

// Job's function
// Job can do something (handler)
// Job can retry
// - Config retry tiems and duraction
// Should be stateful - start, stop, pause, load from config, import from db
// Should have job manager to manage jobs

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	return [6]string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type JobConfig struct {
	Name       string
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     JobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

// func NewJob(handler JobHandler, option ...OptionHdl) *job {
// 	j := job{
// 		config:     JobConfig{MaxTimeout: defaultMaxTimeout, Retries: defaultRetryTime},
// 		handler:    handler,
// 		retryIndex: -1,
// 		state: StateInit,
// 	}
// }
