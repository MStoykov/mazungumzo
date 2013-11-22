package workq

// TODO: Add mutex here.
type Queue []*Item

func (q *Queue) Push(item *Item) {
	*q = append(*q, item)
	go item.Translate()
}

func (q *Queue) Pop() *Item {
	if !q.IsEmpty() {
		item := (*q)[0]
		<-item.Done
		*q = (*q)[1:len((*q))]
		return item
	}
	return nil
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
