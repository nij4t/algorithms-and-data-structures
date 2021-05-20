package queue

type Queue []interface{}

func (q *Queue) Empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func (q *Queue) Enqueue(i interface{}) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() interface{} {
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped
}
