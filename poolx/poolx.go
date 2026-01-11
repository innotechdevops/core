package poolx

import (
	"context"
	"runtime"

	"github.com/alitto/pond"
)

type loopConfig struct {
	maxWorker   int
	maxCapacity int
	ctx         context.Context
}

type LoopOption func(*loopConfig)

func MaxWorker(value int) LoopOption {
	return func(cfg *loopConfig) {
		cfg.maxWorker = value
	}
}

func MaxCapacity(value int) LoopOption {
	return func(cfg *loopConfig) {
		cfg.maxCapacity = value
	}
}

func Context(value context.Context) LoopOption {
	return func(cfg *loopConfig) {
		cfg.ctx = value
	}
}

func LoopCompute[T any](records []T, onCompute func(record T), options ...LoopOption) {
	cfg := loopConfig{
		maxWorker:   runtime.NumCPU(),
		maxCapacity: len(records),
		ctx:         context.Background(),
	}
	for _, option := range options {
		if option != nil {
			option(&cfg)
		}
	}
	if cfg.ctx == nil {
		cfg.ctx = context.Background()
	}
	if cfg.maxWorker <= 0 {
		cfg.maxWorker = runtime.NumCPU()
	}
	if cfg.maxCapacity <= 0 {
		cfg.maxCapacity = len(records)
	}

	if len(records) == 0 {
		return
	}

	pool := pond.New(cfg.maxWorker, cfg.maxCapacity, pond.Context(cfg.ctx))
	size := len(records)

	// Process on worker pool.
	for i := 0; i < size; i++ {
		record := records[i]
		pool.Submit(func() {
			onCompute(record)
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	pool.StopAndWait()
}
