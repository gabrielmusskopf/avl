package avl

import "fmt"

type Ordered[T any] interface {
	Compare(T) int
	Less(T) bool
}

type IndexedTree[K Ordered[K], V any] struct {
	Key   K
	Value V
	BF    int
	Left  *IndexedTree[K, V]
	Right *IndexedTree[K, V]
}

func newIndexedTree[K Ordered[K], V any](k K, v V) *IndexedTree[K, V] {
	return &IndexedTree[K, V]{
		Key:   k,
		Value: v,
		BF:    0,
	}
}

func (n *IndexedTree[K, V]) depth() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.Left.depth(), n.Right.depth())
}

func (n *IndexedTree[K, V]) printn() {
	fmt.Printf("(%v, %d)  ", n.Key, n.BF)
}

func (n *IndexedTree[K, V]) PrettyPrint(padding string) {
	if n == nil {
		return
	}

	fmt.Printf("%s", padding)
	n.printn()
	fmt.Println()

	n.Left.PrettyPrint(padding + "  ")
	n.Right.PrettyPrint(padding + "  ")
}

func (n *IndexedTree[K, V]) PreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%v \n", n.Key)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *IndexedTree[K, V]) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%v ", n.Key)
	n.Right.InOrder()
}

func (n *IndexedTree[K, V]) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Printf("%v ", n.Key)
}

func (n *IndexedTree[K, V]) rotateRight() *IndexedTree[K, V] {
	matches := n.Left
	t2 := matches.Right

	matches.Right = n
	n.Left = t2

	n.BF = n.balanceFactor()
	matches.BF = matches.balanceFactor()

	return matches
}

func (n *IndexedTree[K, V]) rotateLeft() *IndexedTree[K, V] {
	matches := n.Right
	t2 := matches.Left

	matches.Left = n
	n.Right = t2

	n.BF = n.balanceFactor()
	matches.BF = matches.balanceFactor()

	return matches
}

func (n *IndexedTree[K, V]) balance() *IndexedTree[K, V] {
	if n == nil {
		return n
	}

	n.BF = n.balanceFactor()

	//simples
	//rotação simples direita
	if n.BF >= 2 && n.Left.BF >= 0 {
		return n.rotateRight()
	}

	//rotação simples esquerda
	if n.BF <= -2 && n.Right.BF <= 0 {
		return n.rotateLeft()
	}

	//duplas
	//rotação dupla direita
	if n.BF >= 2 && n.Left.BF < 0 {
		n.Left = n.Left.rotateLeft()
		return n.rotateRight()
	}

	//rotação dupla esquerda
	if n.BF <= -2 && n.Right.BF > 0 {
		n.Right = n.Right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

func (n *IndexedTree[K, V]) addRec(k K, v V, i *int) *IndexedTree[K, V] {
	*i++
	if n == nil {
		return newIndexedTree(k, v)
	}
	if k.Less(n.Key) {
		n.Left = n.Left.addRec(k, v, i)
	} else if n.Key.Less(k) {
		n.Right = n.Right.addRec(k, v, i)
	}

	return n.balance()
}

func (n *IndexedTree[K, V]) Add(k K, v V) *IndexedTree[K, V] {
	if n.Search(k) != nil {
		return n
	}
	var i int
	return n.addRec(k, v, &i)
}

func (n *IndexedTree[K, V]) searchRec(k K, i *int) *IndexedTree[K, V] {
	*i++
	if n == nil || n.Key.Compare(k) == 0 {
		return n
	}
	if k.Less(n.Key) {
		return n.Left.searchRec(k, i)
	}
	return n.Right.searchRec(k, i)
}

func (n *IndexedTree[K, V]) searchAllByRec(k K, match *[]*IndexedTree[K, V], matchFunc func(K, K) bool, compareFunc func(K, K) int, iter *int) {
	if n == nil {
		return
	}
	*iter++
	if matchFunc(n.Key, k) {
		*match = append(*match, n)
	}
	if compareFunc(k, n.Key) == -1 || compareFunc(k, n.Key) == 0 {
		n.Left.searchAllByRec(k, match, matchFunc, compareFunc, iter)
	}
	if compareFunc(k, n.Key) == 1 || compareFunc(k, n.Key) == 0 {
		n.Right.searchAllByRec(k, match, matchFunc, compareFunc, iter)
	}
}

// Busca todos os nós que corresponderem de acordo com as funções
// Recebe:
//   k            chave que é usada para comparar com o nó em matchFunc
//   matchFunc    recebe a key do nó e retorna flag indicando se deve ser selecionado
//   compareFunc  recebe a key do nó e retorna 
//                -1 ou 0 para buscar na subárvore esquerda, ou
//                1 ou 0 para buscar na subárvore direita
func (n *IndexedTree[K, V]) SearchAllBy(k K, matchFunc func(K, K) bool, compareFunc func(K, K) int) []*IndexedTree[K, V] {
	matches := make([]*IndexedTree[K, V], 0)
	iter, elasped := measure(func(i *int) {
		n.searchAllByRec(k, &matches, matchFunc, compareFunc, i)
	})
	Debug("%d interações para buscar %v %dns\n", iter, k, elasped)
	return matches
}

func (n *IndexedTree[K, V]) matchAllByRec(match *[]*IndexedTree[K, V], matchFunc func(K) bool, compareFunc func(K) int, iter *int) {
	if n == nil {
		return
	}
	*iter++
	if matchFunc(n.Key) {
		*match = append(*match, n)
	}
	if compareFunc(n.Key) == -1 || compareFunc(n.Key) == 0 {
		n.Left.matchAllByRec(match, matchFunc, compareFunc, iter)
	}
	if compareFunc(n.Key) == 1 || compareFunc(n.Key) == 0 {
		n.Right.matchAllByRec(match, matchFunc, compareFunc, iter)
	}
}

// Busca todos os nós que corresponderem de acordo com as funções
// matchFunc    recebe a key do nó e retorna flag indicando se deve ser selecionado
// compareFunc  recebe a key do nó e retorna 
//              -1 ou 0 para buscar na subárvore esquerda, ou
//              1 ou 0 para buscar na subárvore direita
func (n *IndexedTree[K, V]) MatchAllBy(matchFunc func(K) bool, compareFunc func(K) int) []*IndexedTree[K, V] {
	matches := make([]*IndexedTree[K, V], 0)
	iter, elasped := measure(func(i *int) {
		n.matchAllByRec(&matches, matchFunc, compareFunc, i)
	})
	Debug("%d interações para buscar %dns\n", iter, elasped)
	return matches
}

// Percorre todos os nós da árvore aplicando a walkFunc. Uso não é indicado pois tem tempo linear
func (n *IndexedTree[K, V]) WalkAllBy(walkFunc func(IndexedTree[K, V])) {
	if n == nil {
		return
	}
	walkFunc(*n)
	n.Left.WalkAllBy(walkFunc)
	n.Right.WalkAllBy(walkFunc)
}

// Busca nó com chave idência a informada
func (n *IndexedTree[K, V]) Search(k K) *IndexedTree[K, V] {
	var t *IndexedTree[K, V]
	iter, elasped := measure(func(i *int) {
		t = n.searchRec(k, i)
	})
	Debug("%d interações para buscar %v %dns\n", iter, k, elasped)
	return t
}

// FB(p) = h(sae(p)) - h(sad(p))
func (n *IndexedTree[K, V]) balanceFactor() int {
	return n.Left.depth() - n.Right.depth()
}

func (n *IndexedTree[K, V]) min() *IndexedTree[K, V] {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (n *IndexedTree[K, V]) removeRec(k K, i *int) *IndexedTree[K, V] {
	*i++
	if n == nil {
		return n
	}

	if k.Less(n.Key) {
		n.Left = n.Left.removeRec(k, i)
	} else if n.Key.Less(k) {
		n.Right = n.Right.removeRec(k, i)
	} else {
		// está no node para ser deletado
		if n.Left == nil || n.Right == nil {
			var temp *IndexedTree[K, V]
			if n.Left != nil {
				temp = n.Left
			} else {
				temp = n.Right
			}

			if temp == nil {
				// sem filho
				n = nil
			} else {
				// um filho
				*n = *temp
			}
		} else {
			temp := n.Right.min()
			n.Value = temp.Value
			n.Right = n.Right.removeRec(temp.Key, i)
		}
	}

	return n.balance()
}

func (n *IndexedTree[K, V]) Remove(k K) *IndexedTree[K, V] {
	var i int
	return n.removeRec(k, &i)
}
