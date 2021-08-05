package main

import "context"

// Link: https://blog.golang.org/context-and-structs

// Best practices to pass context as argument
//
// Here, the (*Worker).Fetch and (*Worker).Process methods both accept a context directly.
// With this pass-as-argument design, users can set per-call deadlines, cancellation, and metadata.
// And, it's clear how the context.Context passed to each method will be used:
// there's no expectation that a context.Context passed to one method will be used by any other method.
// This is because the context is scoped to as small an operation as it needs to be,
// which greatly increases the utility and clarity of context in this package.

type Worker struct{}

type Work struct{}

func New() *Worker {
	return &Worker{}
}
func (this *Worker) Fetch(ctx context.Context) (*Work, error) {
	_ = ctx // A per-call ctx is used for cancellation, deadlines, and metadata.
	return &Work{}, nil
}

func (this *Worker) Process(ctx context.Context, work *Work) error {
	_ = ctx // A per-call ctx is used for cancellation, deadlines, and metadata.
	return nil
}

// Worst practices to store context in struct
//
// The (*Worker).Fetch and (*Worker).Process method both use a context stored in Worker.
// This prevents the callers of Fetch and Process (which may themselves have different contexts)
// from specifying a deadline, requesting cancellation, and attaching metadata on a per-call basis.
// For example: the user is unable to provide a deadline just for (*Worker).Fetch,
// or cancel just the (*Worker).Process call.
// The caller's lifetime is intermingled with a shared context,
// and the context is scoped to the lifetime where the Worker is created.

// type Worker struct {
// 	ctx context.Context
// }

// type Work struct{}

// func New(ctx context.Context) *Worker {
// 	return &Worker{ctx: ctx}
// }

// func (this *Worker) Fetch() (*Work, error) {
// 	_ = this.ctx // A shared this.ctx is used for cancellation, deadlines, and metadata.
// 	return &Work{}, nil
// }

// func (this *Worker) Process(work *Work) error {
// 	_ = this.ctx // A shared this.ctx is used for cancellation, deadlines, and metadata.
// 	return nil
// }

func main() {}
