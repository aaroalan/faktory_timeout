package main

import (
	"context"
	fworker "github.com/contribsys/faktory_worker_go"
)

func main() {
mgr := fworker.NewManager()
	mgr.Register("some-name", ProcessJob)
	mgr.Concurrency = 10
	mgr.ProcessStrictPriorityQueues("some-queue")
	mgr.Run()
}

func ProcessJob(_ context.Context, _ ...interface{}) error {
	return nil
}
