package workq

import "github.com/MStoykov/workQ"

type Queue struct {
	queue workQ.WorkQ
	in    chan Item
	out   chan Item
}

func NewQueue() Queue {
	queue := workQ.NewWorkQ()
	in := make(chan Item)
	go func() {
		wrappedIn := queue.In()
		defer close(wrappedIn)
		for item := range in {
			wrappedIn <- item
		}
	}()
	out := make(chan Item)
	go func() {
		wrappedOut := queue.Out()
		defer close(out)
		for item := range wrappedOut {
			basicItem, ok := item.(Item)
			if !ok {
				panic("WAT!?!?!")
			}
			out <- basicItem
		}
	}()
	return Queue{
		queue: queue,
		in:    in,
		out:   out,
	}
}

func (q *Queue) Out() <-chan Item {
	return q.out
}
func (q *Queue) In() chan<- Item {
	return q.in
}
