package balancer

import (
    "github.com/mmycin/intralb/config"
    "net/http"
)

type routerPool struct {
    workers []*worker
    cfg     *config.Options
}

func newRouterPool(cfg *config.Options) *routerPool {
    return &routerPool{cfg: cfg}
}

func (rp *routerPool) spawn(id int, base http.Handler) *worker {
    w := newWorker(id, base, rp.cfg)
    rp.workers = append(rp.workers, w)
    return w
}

func (rp *routerPool) pickLeastBusy() *worker {
    var best *worker
    minLoad := int32(rp.cfg.MaxConcurrentPerRouter + 1)

    for _, w := range rp.workers {
        l := w.Load()
        if l < minLoad {
            minLoad = l
            best = w
        }
    }

    return best
}
