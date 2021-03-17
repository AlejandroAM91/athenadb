package main

import "fmt"

const (
	minBranchDegree = 2
	maxBranchDegree = minBranchDegree * 2
)

type pair struct {
	key int
	val string
}

type node interface {
	insert(value pair)
	search(key int) pair
	split() (node, node)
	full() bool
	key() int
}

type inode struct {
	nodes []node
}

func newINode() *inode {
	return &inode{nodes: make([]node, 0, maxBranchDegree)}
}

func (n *inode) insert(value pair) {
	pos := 1
	for pos < len(n.nodes) && value.key > n.nodes[pos].key() {
		pos++
	}
	pos--
	n.nodes[pos].insert(value)

	if n.nodes[pos].full() {
		l, r := n.nodes[pos].split()
		n.nodes[pos] = l
		n.nodes = append(n.nodes, r)
		pos++
		copy(n.nodes[pos+1:], n.nodes[pos:])
		n.nodes[pos] = r
	}
}

func (n inode) search(key int) pair {
	pos := 1
	for pos < len(n.nodes) && n.nodes[pos].key() <= key {
		pos++
	}
	return n.nodes[pos-1].search(key)
}

func (n inode) split() (node, node) {
	left := newINode()
	for _, v := range n.nodes[:minBranchDegree] {
		left.nodes = append(left.nodes, v)
	}

	right := newINode()
	for _, v := range n.nodes[minBranchDegree:] {
		right.nodes = append(right.nodes, v)
	}
	return left, right
}

func (n inode) full() bool {
	return len(n.nodes) == maxBranchDegree
}

func (l inode) key() int {
	return l.nodes[0].key()
}

type leaf struct {
	values []pair
}

func newLeaf() *leaf {
	return &leaf{values: make([]pair, 0, maxBranchDegree)}
}

func (l *leaf) insert(value pair) {
	l.values = append(l.values, value)
	for i := len(l.values) - 1; i > 0; i-- {
		if value.key < l.values[i-1].key {
			l.values[i] = l.values[i-1]
			l.values[i-1] = value
		}
	}
}

func (l leaf) search(key int) pair {
	fmt.Println("Search 32 in node key:", l.key())
	pos := 0
	for pos < len(l.values) && key < l.values[pos].key {
		pos++
	}
	return l.values[pos]
}

func (l leaf) split() (node, node) {
	left := newLeaf()
	for _, v := range l.values[:minBranchDegree] {
		left.values = append(left.values, v)
	}

	right := newLeaf()
	for _, v := range l.values[minBranchDegree:] {
		right.values = append(right.values, v)
	}
	return left, right
}

func (l leaf) full() bool {
	return len(l.values) == maxBranchDegree
}

func (l leaf) key() int {
	return l.values[0].key
}

type tree struct {
	root node
}

func (t *tree) insert(value pair) {
	t.root.insert(value)

	if t.root.full() {
		l, r := t.root.split()
		tmp := newINode()
		tmp.nodes = append(tmp.nodes, l, r)
		t.root = tmp
	}
}

func (t tree) search(key int) pair {
	return t.root.search(key)
}

func main() {
	tree := tree{root: newLeaf()}
	tree.insert(pair{key: 32, val: "32"})
	tree.insert(pair{key: 16, val: "16"})
	tree.insert(pair{key: 128, val: "128"})
	tree.insert(pair{key: 8, val: "8"})
	tree.insert(pair{key: 4, val: "4"})
	tree.insert(pair{key: 2, val: "2"})
	tree.insert(pair{key: 256, val: "256"})
	tree.insert(pair{key: 512, val: "512"})
	tree.insert(pair{key: 64, val: "64"})
	fmt.Println("Root key:", tree.root.key())
	fmt.Println("Node 1 key:", tree.root.(*inode).nodes[0].key())
	fmt.Println("Node 1.1 key:", tree.root.(*inode).nodes[0].(*inode).nodes[0].key())
	fmt.Println("Node 1.1 values:", tree.root.(*inode).nodes[0].(*inode).nodes[0].(*leaf).values)
	fmt.Println("Node 1.2 key:", tree.root.(*inode).nodes[0].(*inode).nodes[1].key())
	fmt.Println("Node 1.1 values:", tree.root.(*inode).nodes[0].(*inode).nodes[1].(*leaf).values)
	fmt.Println("Node 2 key:", tree.root.(*inode).nodes[1].key())
	fmt.Println("Node 2.1 key:", tree.root.(*inode).nodes[1].(*inode).nodes[0].key())
	fmt.Println("Node 1.1 values:", tree.root.(*inode).nodes[1].(*inode).nodes[0].(*leaf).values)
	fmt.Println("Node 2.2 key:", tree.root.(*inode).nodes[1].(*inode).nodes[1].key())
	fmt.Println("Node 1.1 values:", tree.root.(*inode).nodes[1].(*inode).nodes[1].(*leaf).values)

	fmt.Println("Search 32:", tree.search(32))
}
