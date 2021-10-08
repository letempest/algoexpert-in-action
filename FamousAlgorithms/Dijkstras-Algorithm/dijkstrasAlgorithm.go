package dijkstra

import (
	"math"
)

// Solution 1
// type void interface{}

// // Big O: O(v^2 + e) time | O(v) space
// // where v is the number of vertices and e is the number of edges
// func DijkstrasAlgorithm(start int, edges [][][2]int) []float64 {
// 	var empty void
// 	numberOfVertices := len(edges)
// 	visited := make(map[int]void)
// 	minDistances := make([]float64, numberOfVertices)
// 	for i := range minDistances {
// 		if i != start {
// 			minDistances[i] = math.Inf(1)
// 		}
// 	}

// 	for len(visited) < numberOfVertices {
// 		vertex, currentMinDistance := getUnvisitedVertexWithMinDistance(minDistances, visited)

// 		if currentMinDistance == math.Inf(1) {
// 			minDistances[vertex] = -1
// 			// continue instead of 'break', in case there are multiple separated vertices
// 			// which distance of them all need to be set to -1 to indicate 'NOT-connected'
// 			continue
// 		}

// 		visited[vertex] = empty

// 		neighbors := edges[vertex]
// 		for _, neighbor := range neighbors {
// 			neighborIdx, neighborWeight := neighbor[0], neighbor[1]

// 			if _, exists := visited[neighborIdx]; exists {
// 				continue
// 			}

// 			newPathDistance := currentMinDistance + float64(neighborWeight)
// 			if newPathDistance < minDistances[neighborIdx] {
// 				minDistances[neighborIdx] = newPathDistance
// 			}
// 		}
// 	}

// 	return minDistances
// }

// func getUnvisitedVertexWithMinDistance(minDistances []float64, visited map[int]void) (int, float64) {
// 	currentMinDistance := math.Inf(1)
// 	vertex := -1

// 	for i, distance := range minDistances {
// 		if _, exists := visited[i]; exists {
// 			continue
// 		}
// 		if distance <= currentMinDistance {
// 			currentMinDistance = distance
// 			vertex = i
// 		}
// 	}
// 	return vertex, currentMinDistance
// }

// ==============================================================================================
// Solution 2: min heap to retrive current min distance instead of iterating over distances array
// every time to find out minimum value
// Big O: O((v + e) * log(v)) time | O(v) space
// where v is the number of vertices and e is the number of edges
func DijkstrasAlgorithm(start int, edges [][][2]int) []float64 {
	numberOfVertices := len(edges)
	minDistances := make([]float64, numberOfVertices)
	for i := range minDistances {
		if i != start {
			minDistances[i] = math.Inf(1)
		}
	}
	minDistancesHeap := buildHeap(minDistances)
	minDistancesHeap.update(start, 0)

	for !minDistancesHeap.isEmpty() { // O(v) operation
		vertex, currentMinDistance := minDistancesHeap.remove() // O(log(v)) time

		if currentMinDistance == math.Inf(1) {
			break
		}

		neighbors := edges[vertex]
		for _, neighbor := range neighbors { // total e iterations, visiting each edge exactly once
			neighborIdx, neighborWeight := neighbor[0], neighbor[1]

			newPathDistance := currentMinDistance + float64(neighborWeight)
			if newPathDistance < minDistances[neighborIdx] {
				minDistances[neighborIdx] = newPathDistance
				minDistancesHeap.update(neighborIdx, newPathDistance) // update() call sifUp(), which yields O(log(v)) time
			}
		}
	}

	for i, distance := range minDistances {
		if math.IsInf(distance, 1) {
			minDistances[i] = -1
		}
	}

	return minDistances
}

type vertex struct {
	index            int
	distanceToVertex float64
}

type minHeap struct {
	heap      []vertex
	vertexMap map[int]int // in format: (vertexIdx in distance array) : (corresponding vertexIdx in minHeap.heap array)
}

// Big O: O(n) time | O(1) space
func buildHeap(array []float64) *minHeap {
	firstParentIdx := int(math.Floor(float64(len(array)-2) / 2))
	// encounter a weird case here for problem [Heaps/Sorted K-Sorted Array]
	// if we pass 'Heap: array' directly below, when sorting the same array
	// in place, the heap value is accidently change too, so make a copy here
	copyArr := make([]vertex, len(array))
	vertexMap := make(map[int]int)

	for i := range array {
		copyArr[i].index = i
		copyArr[i].distanceToVertex = math.Inf(1)
		vertexMap[i] = i
	}
	// copyArr[0].distanceToVertex = 0

	minHeap := &minHeap{
		heap:      copyArr,
		vertexMap: vertexMap,
	}
	for currentIdx := firstParentIdx; currentIdx >= 0; currentIdx-- {
		minHeap.siftDown(currentIdx, len(array)-1)
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
		if childTwoIdx != -1 && h.heap[childTwoIdx].distanceToVertex < h.heap[childOneIdx].distanceToVertex {
			idxToSwap = childTwoIdx
		} else {
			idxToSwap = childOneIdx
		}
		if h.heap[currentIdx].distanceToVertex > h.heap[idxToSwap].distanceToVertex {
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
	for currentIdx > 0 && h.heap[currentIdx].distanceToVertex < h.heap[parentIdx].distanceToVertex {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = int(math.Floor((float64(currentIdx - 1)) / 2))
	}
}

// O(log(n)) time | O(1) space
func (h *minHeap) remove() (int, float64) {
	h.swap(0, len(h.heap)-1)
	vertex := h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	delete(h.vertexMap, vertex.index)
	h.siftDown(0, len(h.heap)-1)
	return vertex.index, vertex.distanceToVertex
}

func (h *minHeap) swap(i, j int) {
	h.vertexMap[h.heap[i].index] = j
	h.vertexMap[h.heap[j].index] = i
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

// whenever the update func is called, we are updating the corresponding distance
// to a smaller value, hence it has to be sifted up
func (h *minHeap) update(vertexIdx int, value float64) {
	h.heap[h.vertexMap[vertexIdx]] = vertex{
		index:            vertexIdx,
		distanceToVertex: value,
	}
	h.siftUp(h.vertexMap[vertexIdx])
}
