package workq

import "sync"

type Queue struct {
	items []*Item
	mutex sync.Mutex
}

func (q *Queue) Push(item *Item) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
	go item.Translate()
}

func (q *Queue) Pop() *Item {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if !q.IsEmpty() {
		item := (q.items)[0]
		<-item.Done
		q.items = q.items[1:len(q.items)]
		return item
	}
	return nil
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
