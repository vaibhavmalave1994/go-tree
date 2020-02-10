package tree

type BinaryTree struct {
	root *Node
}

type Node struct {
	val   int64
	left  *Node
	right *Node
}

func (t *BinaryTree) Insert(val int64) *BinaryTree {
	if t.root == nil {
		t.root = &Node{val, nil, nil}
	} else {
		t.root.Insert(val)
	}

	return t
}

func (n *Node) Insert(data int64) {
	if n == nil {
		return
	} else if data <= n.val {
		if n.left == nil {
			n.left = &Node{data, nil, nil}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data, nil, nil}
		} else {
			n.right.Insert(data)
		}
	}
}

func PreOrderTraversal(n *Node) []int64 {
	var result []int64
	if n == nil {
		return result
	}
	result = append(result, n.val)
	left := PreOrderTraversal(n.left)
	right := PreOrderTraversal(n.right)

	result = append(result, left...)
	result = append(result, right...)
	return result
}

func InOrderTraversal(n *Node) []int64 {
	var result []int64
	if n == nil {
		return result
	}
	left := InOrderTraversal(n.left)
	result = append(result, n.val)

	result = append(result, left...)
	right := InOrderTraversal(n.right)
	result = append(result, right...)
	return result
}

func PostOrderTraversal(n *Node) []int64 {
	var result []int64
	if n == nil {
		return result
	}
	left := InOrderTraversal(n.left)
	right := InOrderTraversal(n.right)

	result = append(result, left...)
	result = append(result, right...)
	result = append(result, n.val)
	return result
}
