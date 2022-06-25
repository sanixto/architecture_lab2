package engine

import "sync"

type cmdQueue struct {
	mu   sync.Mutex
	a    []Command
	wait bool

	notEmpty chan struct{}
}

func (cq *cmdQueue) push(cmd Command) {
	cq.mu.Lock()
	cq.a = append(cq.a, cmd)
	cq.mu.Unlock()
	if cq.wait {
		cq.wait = false
		cq.notEmpty <- struct{}{}
	}
}

func (cq *cmdQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.a) == 0 {
		cq.wait = true
		cq.mu.Unlock()
		<-cq.notEmpty
		cq.mu.Lock()
	}

	res := cq.a[0]
	cq.a[0] = nil
	cq.a = cq.a[1:]
	return res
}

func (cq *cmdQueue) empty() bool {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	return len(cq.a) == 0
}
