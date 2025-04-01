package solution

type Node struct {
	Val       int
	Neighbors []*Node
}

// DFS
func dfs(node *Node, clonedMap map[*Node]*Node) {
	if node == nil {
		return
	}
	clonedNode := &Node{
		Val: node.Val,
	}
	clonedMap[node] = clonedNode
	for _, neighbor := range node.Neighbors {
		if _, ok := clonedMap[neighbor]; !ok {
			dfs(neighbor, clonedMap)
		}
		clonedNode.Neighbors = append(clonedNode.Neighbors, clonedMap[neighbor])
	}
}

func cloneGraph1(node *Node) *Node {
	if node == nil {
		return nil
	}
	clonedMap := map[*Node]*Node{}
	dfs(node, clonedMap)
	return clonedMap[node]
}

// ---------------------------------------------

// BFS
func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}
	clonedMap := map[*Node]*Node{
		node: {
			Val: node.Val,
		},
	}
	queue := []*Node{node}
	var currNode *Node
	for len(queue) > 0 {
		currNode, queue = queue[0], queue[1:]
		clonedNode := clonedMap[currNode]
		for _, neighbor := range currNode.Neighbors {
			if _, ok := clonedMap[neighbor]; !ok {
				clonedMap[neighbor] = &Node{
					Val: neighbor.Val,
				}
				queue = append(queue, neighbor)
			}
			clonedNode.Neighbors = append(clonedNode.Neighbors, clonedMap[neighbor])
		}
	}
	return clonedMap[node]
}
