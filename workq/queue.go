package workq

type Queue struct {
	ch chan *Item
}

func NewQueue() *Queue {
	q := new(Queue)
	q.ch = make(chan *Item, 10)
	return q
}

func (q *Queue) Push(item *Item) {
	q.ch <- item
	go item.Translate()
}

func (q *Queue) Pop() <-chan *Item {
	ch := make(chan *Item)
	go func() {
		for item := range q.ch {
			<-item.Done
			ch <- item
		}
	}()
	return ch
}
