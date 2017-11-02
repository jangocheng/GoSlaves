package slaves

import (
	"errors"
	"github.com/eapache/channels"
)

var (
	// ErrChanClosed shows error where channel is closed
	// where writing
	ErrChanClosed = errors.New("error: channel is closed")
)

// Jobs Handle multiple jobs
// enqueuing in buffered channel
type Jobs struct {
	ch *channels.InfiniteChannel
}

// Open creates jobs channel
func (jobs *Jobs) Open() {
	jobs.ch = channels.NewInfiniteChannel()
}

// Put send job to channel
func (jobs *Jobs) Put(job interface{}) {
	jobs.ch.In() <- job
}

// Len Gets the length of jobs to do
func (jobs *Jobs) Len() int {
	return jobs.ch.Len()
}

// Get gets a job from the buffered channel
// if error is returned Close() function have
// been called
func (jobs *Jobs) Get() (interface{}, error) {
	r, ok := <-jobs.ch.Out()
	if !ok {
		return nil, ErrChanClosed
	}
	return r, nil
}

// Close close job channel
func (jobs *Jobs) Close() {
	jobs.ch.Close()
}
