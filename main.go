// golang entry for each AlgoExpert challenge
// #[Greedy Algorithms]/Task Assignment

package main

import (
	"fmt"

	greedy "github.com/letempest/algoexpert-in-action/GreedyAlgorithms/Task-Assignment"
)

func main() {
	workers := 3
	tasks := []int{1, 3, 5, 3, 1, 4}
	pairedTasks := greedy.TaskAssignment(workers, tasks)
	fmt.Println(pairedTasks)
}
