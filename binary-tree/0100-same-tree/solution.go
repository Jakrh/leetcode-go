package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS: Runtime 0ms
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val &&
			isSameTree1(p.Left, q.Left) &&
			isSameTree1(p.Right, q.Right)
	}

	return p == q
}

// -------------------------------------------------

// BFS: Runtime 0ms
type nodePair struct {
	p *TreeNode
	q *TreeNode
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	queue := make([]nodePair, 0, 100)
	queue = append(queue, nodePair{
		p: p,
		q: q,
	})

	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		if pair.p == nil || pair.q == nil {
			return false
		}

		if pair.p.Val != pair.q.Val {
			return false
		}

		if pair.p.Left != nil || pair.q.Left != nil {
			queue = append(queue, nodePair{
				p: pair.p.Left,
				q: pair.q.Left,
			})
		}

		if pair.p.Right != nil || pair.q.Right != nil {
			queue = append(queue, nodePair{
				p: pair.p.Right,
				q: pair.q.Right,
			})
		}
	}

	return true
}
