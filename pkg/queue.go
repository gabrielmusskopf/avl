package avl

import "log"

type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Enqueue(v ...T) {
	q.values = append(q.values, v...)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		log.Fatal("nada para remover da fila!")
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue[T]) Walk(each func(T)) {
	for _, v := range q.values {
		each(v)
	}
}
