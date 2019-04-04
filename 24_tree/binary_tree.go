

//二叉查找树
package _4_tree


type Node struct {
	data int
	left *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewNode(data int) *Node {
	return &Node{data,nil , nil}
}

func New(root *Node) *BinaryTree {
	return &BinaryTree{root}
}

func (tree *BinaryTree) Find(data int) *Node {
	p := tree.root
	for  p != nil && p.data != data {
		if p.data > data {
			p = p.left
		} else {
			p = p.right
		}
	}
	return p
}

func (tree *BinaryTree) Add(data int)  {
	n := &Node{data:data}
	if tree == nil {
		tree = New(n)
		return
	}
	p := tree.root
	for p != nil {
		if data < p.data  {
			if p.left == nil {
				p.left = n
				return
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = n
				return
			}
			p = p.right
		}
	}

}

func (tree *BinaryTree) Delete(data int)  {
	var pp *Node
	p := tree.root

	for p != nil && p.data != data {
		pp = p
		if data < p.data  {
			p = p.left
		} else {
			p = p.right
		}
	}
	if p == nil { //没找到
		return
	}

	//删除节点有两个子节点
	if p.right != nil && p.left != nil {
		minPP := p
		minP := p.right
		for minP.left != nil {
			minPP = minP
			minP = minP.left
		}
		p.data = minP.data
		p = minP
		pp = minPP
	}
	//如果删除节点是叶子节点或者只有一个子节点
	var child *Node // p 节点的子节点
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil {//删除的是头节点
		tree.root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}

}


