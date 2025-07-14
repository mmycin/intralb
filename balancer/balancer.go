package balancer

import (
	"net/http"
	"sync"
	"time"

	"github.com/mmycin/intralb/config"
)

type LoadBalancer struct {
	cfg      *config.Options
	base     http.Handler
	mu       sync.Mutex
	workers  []*worker
	groups   map[string][]*worker
	counter  int
	groupSeq int
}

func New(baseRouter http.Handler, cfg *config.Options) *LoadBalancer {
	if cfg.MaxWorkers == 0 {
		cfg.MaxWorkers = 10
	}
	if cfg.WorkerGroupSize == 0 {
		cfg.WorkerGroupSize = 3
	}

	lb := &LoadBalancer{
		cfg:    cfg,
		base:   baseRouter,
		groups: make(map[string][]*worker),
	}
	lb.addWorkerGroup()
	return lb
}

func (lb *LoadBalancer) BalanceLoad() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		worker := lb.getWorker()
		worker.Enqueue(w, r)
	})
}

func (lb *LoadBalancer) getWorker() *worker {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	var leastBusy *worker
	minLoad := int32(lb.cfg.MaxConcurrentPerRouter) + 1

	for _, group := range lb.groups {
		for _, w := range group {
			if load := w.Load(); load < minLoad {
				minLoad = load
				leastBusy = w
			}
		}
	}

	if minLoad >= int32(lb.cfg.MaxConcurrentPerRouter) && len(lb.workers) < lb.cfg.MaxWorkers {
		return lb.addWorkerGroup()
	}

	return leastBusy
}

func (lb *LoadBalancer) addWorkerGroup() *worker {
	groupID := lb.generateGroupID()
	var dbConn any
	if lb.cfg.DBInitFunc != nil {
		dbConn = lb.cfg.DBInitFunc(groupID)
	}

	var newWorkers []*worker
	for i := 0; i < lb.cfg.WorkerGroupSize && len(lb.workers) < lb.cfg.MaxWorkers; i++ {
		w := newWorkerWithGroup(lb.counter, lb.base, lb.cfg, groupID, dbConn)
		lb.workers = append(lb.workers, w)
		newWorkers = append(newWorkers, w)
		lb.counter++
	}

	lb.groups[groupID] = newWorkers
	if len(newWorkers) > 0 {
		return newWorkers[0]
	}
	return nil
}

func (lb *LoadBalancer) generateGroupID() string {
	lb.groupSeq++
	return time.Now().Format("20060102-150405") + "-" + string(rune(lb.groupSeq))
}

func (lb *LoadBalancer) Shutdown() {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for _, w := range lb.workers {
		w.Stop()
	}
	lb.workers = nil
	lb.groups = nil
}