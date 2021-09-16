// Given two input, jobs and dependencies list, where jobs is an array of jobs, e.g. [1, 2, 3, 4],
// and dependencies list represents an array of dependency tuple, e.g. [[1,2], [1,3], [3,2], [4,2], [4,3]]
// then job 1 is the prerequsite of 2 and 3, in other word, 1 has to be completed before job 2, 3.
// return a sorted of array of jobs so that the order fulfills the restriction stated in dependencies list
// for example above, sorted result should be [1, 4, 3, 2] or [4, 1, 3, 2]

package topology

// Solution 1: depth-first traversal the jobGraph, *if current node has no prerequsites, then this node
// is safe to append into ordered array*, otherwise keep dfs traverse through its prerequsites array to
// populate the result array
// type jobGraph struct {
// 	nodes []*jobNode
// 	graph map[int]*jobNode // given a job (integer value), the graph map this job value to a pointer to its corresponding jobNode
// }

// func NewJobGraph(jobs []int, deps [][2]int) *jobGraph {
// 	jg := &jobGraph{
// 		graph: make(map[int]*jobNode),
// 	}
// 	for _, job := range jobs {
// 		jg.addNode(job)
// 	}

// 	for _, dep := range deps {
// 		prereq, job := dep[0], dep[1]
// 		jg.addPrereq(job, prereq)
// 	}

// 	return jg
// }

// func (j *jobGraph) addNode(job int) {
// 	j.graph[job] = newJobNode(job)
// 	j.nodes = append(j.nodes, j.graph[job])
// }

// func (j *jobGraph) addPrereq(job, prereq int) {
// 	jobNode := j.getNode(job)
// 	prereqNode := j.getNode(prereq)
// 	jobNode.prereqs = append(jobNode.prereqs, prereqNode)
// }

// func (j *jobGraph) getNode(job int) *jobNode {
// 	if _, exists := j.graph[job]; !exists {
// 		j.addNode(job)
// 	}
// 	return j.graph[job]
// }

// type jobNode struct {
// 	job      int
// 	prereqs  []*jobNode
// 	visited  bool
// 	visiting bool
// }

// func newJobNode(job int) *jobNode {
// 	jobNode := jobNode{
// 		job:      job,
// 		visited:  false,
// 		visiting: false,
// 	}
// 	return &jobNode
// }

// // Big O: O(j + d) time | O(j + d) space
// // where j is the number of jobs (vertices), and d is the number of dependencies (edges)
// // first initializing the graph with same space/time complexity, and essentially this
// // algorithm is doing a depth first search over the graph, while within each recursive
// // call, there's only constant time operation
// func TopologicalSort(jobs []int, deps [][2]int) []int {
// 	jobGraph := NewJobGraph(jobs, deps)

// 	return getOrderedJobs(*jobGraph)
// }

// func getOrderedJobs(graph jobGraph) []int {
// 	orderedJobs := make([]int, 0)
// 	nodes := graph.nodes
// 	for len(nodes) > 0 {
// 		// the way we remove node for depthFirstTraversal here will influence the eventual order,
// 		// since node '4' is poped before '1', result will be [4, 1, 3, 2] instead of [1, 4, 3, 2]
// 		node := nodes[len(nodes)-1]
// 		nodes = nodes[:len(nodes)-1]
// 		containsCycle := depthFirstTraverse(node, &orderedJobs)
// 		if containsCycle {
// 			return []int{}
// 		}
// 	}
// 	return orderedJobs
// }

// func depthFirstTraverse(node *jobNode, orderedJobs *[]int) bool {
// 	if node.visited {
// 		return false
// 	}
// 	// if a node is not visited but marked visiting by previous depthFirstTraverse,
// 	// then we know we have a cyclic graph
// 	if node.visiting {
// 		return true
// 	}
// 	node.visiting = true
// 	for _, prereqNode := range node.prereqs {
// 		containsCycle := depthFirstTraverse(prereqNode, orderedJobs)
// 		if containsCycle {
// 			return true
// 		}
// 	}
// 	node.visited = true
// 	node.visiting = false
// 	*orderedJobs = append(*orderedJobs, node.job)
// 	return false
// }

// ===============================================================================
// Solution 2: for each jobNode, keep track of its deppendencies and number of prerequsites,
// iterate through all the graph nodes, if a node's number of prerequsites equal to 0,
// then this node is safe to be poped from array nodesWithoutPrereqs and  append into resulting
// ordered array, when we append such a node, also decrease all of its dependent node's number
// of prerequsites by 1, since from the algorithm's point of view, these nodes no longer depend
// on current node, and if a node's number of preqs equal to 0 after decrement, this node should
// be pushed into the array of nodesWithoutPrereqs, loop through the above logic until array
// nodesWithoutPrereqs is empty
type jobGraph struct {
	nodes []*jobNode
	graph map[int]*jobNode // given a job (integer value), the graph map this job value to a pointer to its corresponding jobNode
}

func NewJobGraph(jobs []int, deps [][2]int) *jobGraph {
	jg := &jobGraph{
		graph: make(map[int]*jobNode),
	}
	for _, job := range jobs {
		jg.addNode(job)
	}

	for _, dep := range deps {
		job, dep := dep[0], dep[1]
		jg.addDep(job, dep)
	}

	return jg
}

func (j *jobGraph) addNode(job int) {
	j.graph[job] = newJobNode(job)
	j.nodes = append(j.nodes, j.graph[job])
}

func (j *jobGraph) addDep(job, dep int) {
	jobNode := j.getNode(job)
	depNode := j.getNode(dep)
	jobNode.deps = append(jobNode.deps, depNode)
	depNode.numOfPrereqs += 1
}

func (j *jobGraph) getNode(job int) *jobNode {
	if _, exists := j.graph[job]; !exists {
		j.addNode(job)
	}
	return j.graph[job]
}

type jobNode struct {
	job          int
	deps         []*jobNode
	numOfPrereqs int
}

func newJobNode(job int) *jobNode {
	jobNode := jobNode{
		job: job,
	}
	return &jobNode
}

// Big O: O(j + d) time | O(j + d) space
func TopologicalSort(jobs []int, deps [][2]int) []int {
	jobGraph := NewJobGraph(jobs, deps)

	return getOrderedJobs(*jobGraph)
}

func getOrderedJobs(graph jobGraph) []int {
	orderedJobs := make([]int, 0)
	nodesWithoutPrereqs := make([]*jobNode, 0)
	for _, node := range graph.nodes {
		if node.numOfPrereqs == 0 {
			nodesWithoutPrereqs = append(nodesWithoutPrereqs, node)
		}
	}

	for len(nodesWithoutPrereqs) > 0 {
		node := nodesWithoutPrereqs[len(nodesWithoutPrereqs)-1]
		nodesWithoutPrereqs = nodesWithoutPrereqs[:len(nodesWithoutPrereqs)-1]
		orderedJobs = append(orderedJobs, node.job)
		removeDeps(node, &nodesWithoutPrereqs)
	}

	// after the for loop above, if any node in the graph still has prerequsites
	// then we know the graph is NOT acyclic, i.e. a cyclic graph,
	// cannot be sorted topologically
	graphHasEdges := false
	for _, node := range graph.nodes {
		if node.numOfPrereqs > 0 {
			graphHasEdges = true
		}
		if graphHasEdges {
			return []int{}
		}
	}
	return orderedJobs
}

func removeDeps(node *jobNode, nodesWithoutPrereqs *[]*jobNode) {
	for len(node.deps) > 0 {
		dep := node.deps[len(node.deps)-1]
		node.deps = node.deps[:len(node.deps)-1]
		dep.numOfPrereqs -= 1
		if dep.numOfPrereqs == 0 {
			*nodesWithoutPrereqs = append(*nodesWithoutPrereqs, dep)
		}
	}
}
