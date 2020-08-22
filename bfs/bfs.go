package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Value    int
	Children []*Node
}

type NodeWithDepth struct {
	node  *Node
	depth int
}

func maxDepth(root *Node) int {
	maxDepth := 0
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(NodeWithDepth{depth: 1, node: root})
	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		nodeWithDepth := elem.Value.(NodeWithDepth)
		depth := nodeWithDepth.depth
		if depth > maxDepth {
			maxDepth = depth
		}
		node := nodeWithDepth.node
		for _, n := range node.Children {
			queue.PushBack(NodeWithDepth{n, depth + 1})
		}
	}

	return maxDepth
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	queue := list.New()
	average := 0
	tmp := make(map[int]*Employee, len(employees))
	for _, e := range employees {
		tmp[e.Id] = e
	}
	queue.PushBack(tmp[id])

	for queue.Len() > 0 {
		element := queue.Front()
		poll := element.Value.(*Employee)
		queue.Remove(element)
		average += poll.Importance
		for _, e := range poll.Subordinates {
			queue.PushBack(tmp[e])
		}

	}

	return average
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func bfs(s *[][]int,level int,root *TreeNode){
	if root == nil{
		return
	}
	if level == len((*s)){
		(*s) = append((*s),[]int{})
	}
	(*s)[level] = append((*s)[level],root.Val)
	for _, n := range []*TreeNode{root.Left,root.Right}{
		bfs(s,level+1,n)
	}

}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	s := [][]int{}
	bfs(&s,0,root)
	for i := 0; i < len(s) / 2; i++ {
		s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
	}
	return s
}

func main() {

	root := &TreeNode{3,
		&TreeNode{9, nil, nil}, &TreeNode{20,
			&TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(levelOrderBottom(root))
}
