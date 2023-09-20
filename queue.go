package avl

import "log"

type Queue struct {
	values []*TreeNode
}

func (q *Queue) enqueue(v... *TreeNode) {
	q.values = append(q.values, v...)
}

func (q *Queue) dequeue() *TreeNode {
	if q.isEmpty() {
		log.Fatal("nada para remover da fila!")
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}

func (q *Queue) isEmpty() bool {
	return len(q.values) == 0
}

