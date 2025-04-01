package solution

import (
	"fmt"
	"sort"
)

// Adjacency List of [[0,1],[1,2],[2,0]]
// 0: 1, 2
// 1: 0, 2
// 2: 1, 0

// Adjacency List of [[0,1],[0,2],[3,5],[5,4],[4,3]]
// 0: 1, 2
// 1: 0
// 2: 0
// 3: 5, 4
// 4: 5, 3
// 5: 3, 4

// BFS
func validPath1(n int, edges [][]int, source int, destination int) bool {
	adjList := make(map[int][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make(map[int]bool, n)
	visited[source] = true
	q := make([]int, 0, n)
	q = append(q, source)

	var currVertex int
	for len(q) > 0 {
		currVertex = q[0]
		q = q[1:]

		if currVertex == destination {
			return true
		}

		for _, linkedVertex := range adjList[currVertex] {
			if linkedVertex == destination {
				return true
			}
			if !visited[linkedVertex] {
				q = append(q, linkedVertex)
				visited[linkedVertex] = true
			}
		}
	}

	return false
}

//--------------------------------------------------

// DFS
func dfs(currVertex, destination int, adjList map[int][]int, visited map[int]bool) bool {
	if currVertex == destination {
		return true
	}
	if _, ok := visited[currVertex]; ok {
		return false
	}

	visited[currVertex] = true
	for _, linkedVertex := range adjList[currVertex] {
		if dfs(linkedVertex, destination, adjList, visited) {
			return true
		}
	}

	return false
}

// DFS
func validPath2(n int, edges [][]int, source int, destination int) bool {
	adjList := make(map[int][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make(map[int]bool, n)

	return dfs(source, destination, adjList, visited)
}

//--------------------------------------------------

// Adjacency Metrix of [[0,1],[1,2],[2,0]]
//   0 1 2
// 0 1 1 1
// 1 1 1 1
// 2 1 1 1

// Adjacency Metrix of [[0,1],[0,2],[3,5],[5,4],[4,3]]
//   0 1 2 3 4 5
// 0 1 1 1 0 0 0
// 1 1 1 0 0 0 0
// 2 1 0 1 0 0 0
// 3 0 0 0 1 1 1
// 4 0 0 0 1 1 1
// 5 0 0 0 1 1 1

// Use Adjacency Matrix and BFS
// Out of memory on leetcode for 5000 vertexes
func validPath3(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	// Build an Adjacency Matrix for the edges
	vertices := map[int]bool{}
	for _, edge := range edges {
		vertices[edge[0]] = true
		vertices[edge[1]] = true
	}

	adjMetrix := make(map[int]map[int]bool, n)
	for vertex := range vertices {
		adjMetrix[vertex] = make(map[int]bool, n)
	}
	for _, edge := range edges {
		adjMetrix[edge[0]][edge[1]] = true
		adjMetrix[edge[1]][edge[0]] = true
		adjMetrix[edge[0]][edge[0]] = true
		adjMetrix[edge[1]][edge[1]] = true
	}

	// printAdjMetrix(adjMetrix)

	// Check the path exists or does not
	if adjMetrix[source][destination] {
		return true
	}

	for connectedVertex := range adjMetrix[source] {
		visited := make(map[int]bool, n)
		visited[connectedVertex] = true
		q := []int{connectedVertex}
		for len(q) > 0 {
			currVertex := q[0]
			q = q[1:]

			if currVertex == destination {
				return true
			}

			visited[currVertex] = true
			for nextVertex := range adjMetrix[currVertex] {
				if !visited[nextVertex] {
					q = append(q, nextVertex)
				}
			}
		}
	}

	return false
}

func printAdjMetrix(adjMetrix map[int]map[int]bool) {
	vertices := make([]int, 0, len(adjMetrix))
	for vertex := range adjMetrix {
		vertices = append(vertices, vertex)
	}
	sort.Ints(vertices)

	digits := func(num int) int {
		if num == 0 {
			return 1
		}

		count := 0
		for num != 0 {
			num /= 10
			count++
		}

		return count
	}

	maxDigits := digits(vertices[len(vertices)-1])
	printFormat := fmt.Sprintf("%%0%dd ", maxDigits)
	fmt.Printf(" ")
	for i := 0; i < maxDigits; i++ {
		fmt.Printf(" ")
	}
	for _, vertex := range vertices {
		fmt.Printf(printFormat, vertex)
	}
	fmt.Printf("\n")

	converter := map[bool]int{false: 0, true: 1}
	for _, vertexCol := range vertices {
		fmt.Printf(printFormat, vertexCol)
		for _, vertexRow := range vertices {
			fmt.Printf(printFormat, converter[adjMetrix[vertexCol][vertexRow]])
		}
		fmt.Printf("\n")
	}
}
