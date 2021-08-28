// Given a directed, un-weighted graph, decide whether there's a cycle in it
// (in which starting from any node, following the direction, you can eventually revisit the starting node again)
// srcGraph below is an adjacency list, index denotes distinct node, value denotes collection of other nodes it can traverse to
// e.g. for index 0, which is node 0, you can traverse to node 1 and 3.
// cycles in this graph: 1) 0 -> 1 -> 2 -> 0;   2) 0 -> 1 -> 4 -> 2 -> 0
package main

import "fmt"

func main() {
	srcGraph := [][]int{{1, 3}, {2, 3, 4}, {0}, {}, {2, 5}, {}}
	fmt.Println(cycleInGraph(srcGraph))
}

// Solution 1
// Big O: O(v + e) time, depth-first search | O(v) space; v = vertex, e = edge
// func cycleInGraph(adjacencyList [][]int) bool {
// 	visited := make([]bool, len(adjacencyList))
// 	currentlyInStack := make([]bool, len(adjacencyList))

// 	// for loop for all nodes ensure that even for a graph with multiple disconnected sub-graphs, every node still has a chance to be visited
// 	for node := range adjacencyList {
// 		if visited[node] {
// 			continue
// 		}
// 		containsCycle := isNodeInCycle(node, adjacencyList, &visited, &currentlyInStack)
// 		if containsCycle {
// 			return true
// 		}
// 	}
// 	return false
// }

// func isNodeInCycle(node int, adjacencyList [][]int, visited, currentlyInStack *[]bool) bool {
// 	(*visited)[node] = true
// 	(*currentlyInStack)[node] = true

// 	for _, neighbor := range adjacencyList[node] {
// 		if !(*visited)[neighbor] {
// 			containsCycle := isNodeInCycle(neighbor, adjacencyList, visited, currentlyInStack)
// 			if containsCycle {
// 				return true
// 			}
// 		} else if (*currentlyInStack)[neighbor] {
// 			// if the program makes into this block, then the neighbor node is visited and in the stack already
// 			// which means current node is a descendant linking back to its ancestor, cycle detected!
// 			// for depth-first search to work, ancestor node is removed AFTER all descendant nodes are visited (i.e. removed from stack)
// 			return true
// 		}
// 	}

// 	(*currentlyInStack)[node] = false
// 	return false
// }

// ============================================================================
// Solution 2
const (
	UNVISITED = iota
	INSTACK
	FINISHED
)

// O(v + e) time | O(v) space
func cycleInGraph(adjacencyList [][]int) bool {
	statuses := make([]int, len(adjacencyList))

	for node := range adjacencyList {
		if statuses[node] != UNVISITED {
			continue
		}
		containsCycle := traverseAndMarkNodes(node, adjacencyList, &statuses)
		if containsCycle {
			return true
		}
	}

	return false
}

func traverseAndMarkNodes(node int, adjacencyList [][]int, statuses *[]int) bool {
	(*statuses)[node] = INSTACK

	for _, neighbor := range adjacencyList[node] {
		neighborStatus := (*statuses)[neighbor]
		if neighborStatus == INSTACK {
			return true
		} else if neighborStatus != UNVISITED {
			continue
		}

		// the neighbor node is not visited yet, invoke depth-first search on it
		containsCycle := traverseAndMarkNodes(neighbor, adjacencyList, statuses)
		if containsCycle {
			return true
		}
	}

	// depth-first search on current node is finished, mark it like so
	(*statuses)[node] = FINISHED
	return false
}
