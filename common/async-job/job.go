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
	Execute()
	Retry()
	State()
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

type JobHanlder func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateRetryFailed
)
