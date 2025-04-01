package graph

import "testing"

type Node struct {
	Val       int
	Neighbors []*Node
}

func adjList2Graph(adjList [][]int) *Node {
	nodes := make([]*Node, len(adjList))
	for i := range adjList {
		nodes[i] = &Node{
			Val: i + 1,
		}
	}
	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}
	if len(nodes) > 0 {
		return nodes[0]
	} else {
		return nil
	}
}

func graphsEqual(t *testing.T, original, cloned *Node) bool {
	logEnabled := false
	oriNodes := map[int]*Node{}
	clonedNodes := map[int]*Node{}

	visited := map[*Node]bool{original: true}
	queue := []*Node{original}
	var currNode *Node
	for len(queue) > 0 {
		currNode, queue = queue[0], queue[1:]
		for _, neighbor := range currNode.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		oriNodes[currNode.Val] = currNode
	}

	visited = map[*Node]bool{original: true}
	queue = append(queue, cloned)
	for len(queue) > 0 {
		currNode, queue = queue[0], queue[1:]
		for _, neighbor := range currNode.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		clonedNodes[currNode.Val] = currNode
	}

	if logEnabled {
		t.Logf("%d, %d\n", len(oriNodes), len(clonedNodes))
	}

	if len(oriNodes) != len(clonedNodes) {
		return false
	}
	for k := range oriNodes {
		if logEnabled {
			t.Logf("(%d %v %p), (%d %v %p)\n",
				oriNodes[k].Val,
				oriNodes[k].Neighbors,
				oriNodes[k],
				clonedNodes[k].Val,
				clonedNodes[k].Neighbors,
				clonedNodes[k])
		}

		if oriNodes[k] == clonedNodes[k] {
			// It means they are the same node
			return false
		}

		if len(oriNodes[k].Neighbors) != len(clonedNodes[k].Neighbors) {
			return false
		}

		// Check if neighbors are the same
		neighborMap := map[int]bool{}
		for _, node := range oriNodes[k].Neighbors {
			neighborMap[node.Val] = true
		}
		if len(neighborMap) != len(clonedNodes[k].Neighbors) {
			return false
		}
		for _, node := range clonedNodes[k].Neighbors {
			if ok := neighborMap[node.Val]; !ok {
				return false
			}
		}

	}

	return true
}
