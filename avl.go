package avl

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Value int
	BF    int
	Left  *TreeNode
	Right *TreeNode
}

func FromArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	var tree *TreeNode
	for i := 0; i < len(arr); i++ {
		tree = tree.Add(arr[i])
	}
	return tree
}

func (n *TreeNode) AddFromArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr); i++ {
		n = n.Add(arr[i])
	}
	return n
}

func create(v int) *TreeNode {
	return &TreeNode{
		Value: v,
		BF:    0,
	}
}

func (n *TreeNode) depth() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.Left.depth(), n.Right.depth())
}

func (n *TreeNode) printn() {
	fmt.Printf("(%d, %d)  ", n.Value, n.BF)
}

func (n *TreeNode) PrettyPrint(padding string) {
	if n == nil {
		return
	}

	fmt.Printf("%s", padding)
	n.printn()
	fmt.Println()

	n.Left.PrettyPrint(padding + "  ")
	n.Right.PrettyPrint(padding + "  ")
}

func (n *TreeNode) BFS() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value)
	q := Queue{}
	q.enqueue(n.Left, n.Right)

	for !q.isEmpty() {
		x := q.dequeue()
		if x == nil {
			continue
		}
		fmt.Printf("%d ", x.Value)
		q.enqueue(x.Left, x.Right)
	}
}

func (n *TreeNode) PreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d \n", n.Value)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *TreeNode) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%d ", n.Value)
	n.Right.InOrder()
}

func (n *TreeNode) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Printf("%d ", n.Value)
}

func (n *TreeNode) rotateRight() *TreeNode {
	x := n.Left
	t2 := x.Right

	x.Right = n
	n.Left = t2

	n.BF = n.balanceFactor()
	x.BF = x.balanceFactor()

	return x
}

func (n *TreeNode) rotateLeft() *TreeNode {
	x := n.Right
	t2 := x.Left

	x.Left = n
	n.Right = t2

	n.BF = n.balanceFactor()
	x.BF = x.balanceFactor()

	return x
}

func (n *TreeNode) balance() *TreeNode {
	if n == nil {
		return n
	}

	n.BF = n.balanceFactor()

	//simples
	//rotação simples direita
	if n.BF >= 2 && n.Left.BF >= 0 {
		Debug("Nó %d com FB: %d. Rotação simples a direita\n", n.Value, n.BF)
		return n.rotateRight()
	}

	//rotação simples esquerda
	if n.BF <= -2 && n.Right.BF <= 0 {
		Debug("Nó %d com FB: %d. Rotação simples a esquerda\n", n.Value, n.BF)
		return n.rotateLeft()
	}

	//duplas
	//rotação dupla direita
	if n.BF >= 2 && n.Left.BF < 0 {
		Debug("Nó %d com FB: %d. Rotação dupla a direita\n", n.Value, n.BF)
		n.Left = n.Left.rotateLeft()
		return n.rotateRight()
	}

	//rotação dupla esquerda
	if n.BF <= -2 && n.Right.BF > 0 {
		Debug("Nó %d com FB: %d. Rotação dupla a esquerda\n", n.Value, n.BF)
		n.Right = n.Right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

func (n *TreeNode) addRec(v int, i *int) *TreeNode {
	*i++
	if n == nil {
		return create(v)
	}
	if v < n.Value {
		Debug("%d é menor do que %d\n", v, n.Value)
		n.Left = n.Left.addRec(v, i)
	} else if v > n.Value {
		Debug("%d é maior do que %d\n", v, n.Value)
		n.Right = n.Right.addRec(v, i)
	}

	return n.balance()
}

func (n *TreeNode) Add(v int) *TreeNode {
	i := 0
	t := n.addRec(v, &i)
	Debug("%d interações para inserir %d\n", i, v)
	return t
}

func (n *TreeNode) serachRec(v int, i *int) *TreeNode {
	*i++
	if n == nil || n.Value == v {
		return n
	}
	if v < n.Value {
		return n.Left.serachRec(v, i)
	}
	return n.Right.serachRec(v, i)
}

func (n *TreeNode) Serach(v int) *TreeNode {
	i := 0
	t := n.serachRec(v, &i)
	Debug("%d interações para buscar %d\n", i, v)
	return t
}

// FB(p) = h(sae(p)) - h(sad(p))
func (n *TreeNode) balanceFactor() int {
	return n.Left.depth() - n.Right.depth()
}

func (n *TreeNode) min() *TreeNode {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (n *TreeNode) removeRec(v int, i *int) *TreeNode {
	*i++
	if n == nil {
		return n
	}

	if v < n.Value {
		n.Left = n.Left.removeRec(v, i)
	} else if v > n.Value {
		n.Right = n.Right.removeRec(v, i)
	} else {
		// está no node para ser deletado
		if n.Left == nil || n.Right == nil {
			var temp *TreeNode
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
			n.Right = n.Right.removeRec(temp.Value, i)
		}
	}

	return n.balance()
}

func (n *TreeNode) Remove(v int) *TreeNode {
	i := 0
	t := n.removeRec(v, &i)
	Debug("%d interações para remover %d\n", i, v)
	return t
}

/*
type CountIter struct {
    i int
    operation func(int, *int) *TreeNode
}

func (n *TreeNode) countIter(op func(int, *int) *TreeNode) *CountIter {
    return &CountIter{
        operation: op,
    }
}

func (c *CountIter) exec(v int) *TreeNode {
	i := 0
    t := c.operation(v, &i)
	Debug("%d interações\n", i)
    return t
}
*/
