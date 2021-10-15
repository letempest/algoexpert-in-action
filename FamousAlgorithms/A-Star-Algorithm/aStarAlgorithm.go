package famous

import (
	"fmt"
	"math"
)

type node struct {
	id                     string
	row                    int
	col                    int
	value                  int
	distanceFromStart      float64 // G score
	estimatedDistanceToEnd float64 // F score
	cameFrom               *node
}

func newNode(row, col, value int) *node {
	node := &node{
		id:                     fmt.Sprintf("%v-%v", row, col),
		row:                    row,
		col:                    col,
		value:                  value,
		distanceFromStart:      math.Inf(1), // initialize each node's G score to be possitive infinity
		estimatedDistanceToEnd: math.Inf(1), // F score to be possitive infinity
	}
	return node
}

// Big O: O(w * h * log(w * h)) time | O(w * h) space, where w, h is the width and height of graph array respectively
func AStarAlgorithm(startRow, startCol, endRow, endCol int, graph [][]int) [][2]int {
	nodes := initializeNodes(graph) // O(w*h) space

	startNode := nodes[startRow][startCol]
	endNode := nodes[endRow][endCol]

	startNode.distanceFromStart = 0
	startNode.estimatedDistanceToEnd = calculateManhattanDistance(startNode, endNode)

	nodesToVisit := buildHeap([]*node{startNode})

	for !nodesToVisit.isEmpty() { // at most O(w*h) time, because at most w*h nodes in the min heap
		currentMinDistanceNode := nodesToVisit.remove() // O(log(w*h)) time to retrive node and siftDown the min heap

		if currentMinDistanceNode == endNode {
			break
		}

		neighbors := getNeighboringNodes(currentMinDistanceNode, nodes)
		for _, neighbor := range neighbors {
			// value = 1 denotes that the neighbor node is an obstacle
			if neighbor.value == 1 {
				continue
			}

			tentativeDistanceToNeighbor := currentMinDistanceNode.distanceFromStart + 1

			if tentativeDistanceToNeighbor >= neighbor.distanceFromStart {
				continue
			}
			neighbor.cameFrom = currentMinDistanceNode
			neighbor.distanceFromStart = tentativeDistanceToNeighbor
			neighbor.estimatedDistanceToEnd = tentativeDistanceToNeighbor + calculateManhattanDistance(neighbor, endNode)

			if !nodesToVisit.containsNode(neighbor) {
				nodesToVisit.insert(neighbor) // O(log(w*h)) time to insert - siftUp()
			} else {
				nodesToVisit.update(neighbor) // O(log(w*h)) time to update - siftUp()
			}
		}
	}

	return reconstructPath(endNode)
}

func initializeNodes(graph [][]int) [][]*node {
	nodes := make([][]*node, 0)

	for row := range graph {
		nodes = append(nodes, []*node{})
		for col, value := range graph[row] {
			node := newNode(row, col, value)
			nodes[row] = append(nodes[row], node)
		}
	}

	return nodes
}

func calculateManhattanDistance(currentNode, endNode *node) float64 {
	currentRow := currentNode.row
	currentCol := currentNode.col
	endRow := endNode.row
	endCol := endNode.col

	return math.Abs(float64(currentRow-endRow)) + math.Abs(float64(currentCol-endCol))
}

func getNeighboringNodes(n *node, nodes [][]*node) []*node {
	neighbors := make([]*node, 0)

	numRows := len(nodes)
	numCols := len(nodes[0])
	row := n.row
	col := n.col

	if row < numRows-1 { // DOWN
		neighbors = append(neighbors, nodes[row+1][col])
	}
	if row > 0 { // UP
		neighbors = append(neighbors, nodes[row-1][col])
	}
	if col < numCols-1 { // RIGHT
		neighbors = append(neighbors, nodes[row][col+1])
	}
	if col > 0 { // LEFT
		neighbors = append(neighbors, nodes[row][col-1])
	}

	return neighbors
}

func reconstructPath(endNode *node) [][2]int {
	if endNode.cameFrom == nil {
		return [][2]int{}
	}

	path := make([][2]int, 0)
	currentNode := endNode

	for currentNode != nil {
		path = append(path, [2]int{currentNode.row, currentNode.col})
		currentNode = currentNode.cameFrom
	}

	return reverse(path)
}

func reverse(path [][2]int) [][2]int {
	for i, j := 0, len(path)-1; i < j; {
		path[i], path[j] = path[j], path[i]
		i += 1
		j -= 1
	}
	return path
}

// ===========================================================
// min heap construction below
type minHeap struct {
	heap                []*node
	nodePositionsInHeap map[string]int // in format: (node.id) : (index of node in heap array)
}

// Big O: O(n) time | O(1) space
func buildHeap(array []*node) *minHeap {
	firstParentIdx := int(math.Floor(float64(len(array)-2) / 2))
	nodePositionsInHeap := make(map[string]int)

	minHeap := &minHeap{
		heap:                array,
		nodePositionsInHeap: nodePositionsInHeap,
	}
	for currentIdx := firstParentIdx; currentIdx >= 0; currentIdx-- {
		minHeap.siftDown(currentIdx, len(array)-1)
	}
	for idx, node := range array {
		minHeap.nodePositionsInHeap[node.id] = idx
	}

	return minHeap
}

func (h *minHeap) isEmpty() bool {
	return len(h.heap) == 0
}

// O(log(n)) time | O(1) space
func (h *minHeap) siftDown(currentIdx, endIdx int) {
	childOneIdx := 2*currentIdx + 1
	var childTwoIdx int
	for childOneIdx <= endIdx {
		if 2*currentIdx+2 <= endIdx {
			childTwoIdx = 2*currentIdx + 2
		} else {
			childTwoIdx = -1
		}
		idxToSwap := currentIdx
		if childTwoIdx != -1 && h.heap[childTwoIdx].estimatedDistanceToEnd < h.heap[childOneIdx].estimatedDistanceToEnd {
			idxToSwap = childTwoIdx
		} else {
			idxToSwap = childOneIdx
		}
		if h.heap[currentIdx].estimatedDistanceToEnd > h.heap[idxToSwap].estimatedDistanceToEnd {
			h.swap(currentIdx, idxToSwap)
			currentIdx = idxToSwap
			childOneIdx = 2*currentIdx + 1
		} else {
			break
		}
	}
}

// O(log(n)) time | O(1) space
func (h *minHeap) siftUp(currentIdx int) {
	parentIdx := int(math.Floor((float64(currentIdx - 1)) / 2))
	for currentIdx > 0 && h.heap[currentIdx].estimatedDistanceToEnd < h.heap[parentIdx].estimatedDistanceToEnd {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = int(math.Floor((float64(currentIdx - 1)) / 2))
	}
}

// O(log(n)) time | O(1) space
func (h *minHeap) remove() *node {
	h.swap(0, len(h.heap)-1)
	node := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	delete(h.nodePositionsInHeap, node.id)
	h.siftDown(0, len(h.heap)-1)
	return node
}

func (h *minHeap) insert(node *node) {
	h.heap = append(h.heap, node)
	h.nodePositionsInHeap[node.id] = len(h.heap) - 1
	h.siftUp(len(h.heap) - 1)
}

func (h *minHeap) swap(i, j int) {
	h.nodePositionsInHeap[h.heap[i].id] = j
	h.nodePositionsInHeap[h.heap[j].id] = i
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *minHeap) containsNode(node *node) bool {
	_, exists := h.nodePositionsInHeap[node.id]
	return exists
}

// whenever the update func is called, we are updating the corresponding distance
// to a smaller value, hence it has to be sifted up
func (h *minHeap) update(node *node) {
	h.siftUp(h.nodePositionsInHeap[node.id])
}
