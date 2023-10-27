package worker

import (
	"apod/service"
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	interval = 4 * time.Hour
)

type Worker struct {
	mu      sync.Mutex
	ctx     context.Context
	cancel  context.CancelFunc
	running bool
}

func NewWorker(ctx context.Context) *Worker {
	ctx, cancel := context.WithCancel(ctx)
	return &Worker{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (w *Worker) Start() {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.running {
		return
	}
	w.running = true
	go w.loop()
}

func (w *Worker) loop() {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-w.ctx.Done():
			return
		case now := <-ticker.C:
			fmt.Println("Worker is running.")
			fmt.Println("Current time: ", now.UTC().String())

			var apodService service.ApodService
			apodService.SaveApod()

			fmt.Println("Worker is idle.")
		}
	}
}
