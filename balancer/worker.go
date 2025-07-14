package balancer

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/mmycin/intralb/config"
)

type worker struct {
	id      int
	groupID string
	handler http.Handler
	queue   chan *job
	load    int32
	cfg     *config.Options
	dbConn  any // Optional DB connection per worker group
}

type job struct {
	w    http.ResponseWriter
	r    *http.Request
	done chan struct{}
}

func newWorker(id int, base http.Handler, cfg *config.Options) *worker {
	w := &worker{
		id:      id,
		handler: base,
		queue:   make(chan *job, cfg.QueueSize),
		cfg:     cfg,
	}
	go w.processJobs()
	return w
}

// Additional constructor for workers with group context
func newWorkerWithGroup(id int, base http.Handler, cfg *config.Options, groupID string, dbConn any) *worker {
	w := newWorker(id, base, cfg)
	w.groupID = groupID
	w.dbConn = dbConn
	return w
}

func (w *worker) processJobs() {
	for j := range w.queue {
		atomic.AddInt32(&w.load, 1)
		w.serve(j)
		atomic.AddInt32(&w.load, -1)
		close(j.done)
	}
}

func (w *worker) serve(j *job) {
	ctx, cancel := context.WithTimeout(j.r.Context(), time.Duration(w.cfg.GracefulTimeoutSeconds)*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		w.handler.ServeHTTP(j.w, j.r.WithContext(ctx))
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		if !w.responseStarted(j.w) {
			http.Error(j.w, "Request Timeout", http.StatusGatewayTimeout)
		}
	}
}

func (w *worker) responseStarted(rw http.ResponseWriter) bool {
	if r, ok := rw.(interface{ HeaderWritten() bool }); ok {
		return r.HeaderWritten()
	}
	return false
}

func (w *worker) Enqueue(wr http.ResponseWriter, r *http.Request) {
	j := &job{
		w:    wr,
		r:    r,
		done: make(chan struct{}),
	}
	w.queue <- j
	<-j.done // Wait for processing to complete
}

func (w *worker) Load() int32 {
	return atomic.LoadInt32(&w.load)
}

func (w *worker) Stop() {
	close(w.queue)
}