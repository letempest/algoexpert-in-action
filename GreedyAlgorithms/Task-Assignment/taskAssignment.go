package greedy

import "sort"

// Big O: O(nlog(n)) time | O(n) space
// O(nlog(n)) time because we're creating a sorted copy of input array,
// and the main logic of the algorithm is just one iteration of the sorted array, O(n) time
// in general, nlog(n) > n, still O(nlog(n)) time
// this function should return the indices pairs of the input tasks
// with which the combining duration of each pair is shortest, hence come optimal assignment
func TaskAssignment(k int, tasks []int) [][2]int {
	pairedTasks := [][2]int{}
	taskDurationsToIndices := getTaskDurationsToIndices(tasks)

	sortedTasks := append([]int(nil), tasks...)
	sort.Ints(sortedTasks)
	for idx := 0; idx < k; idx++ {
		task1Duration := sortedTasks[idx]
		indicesWithTask1Duration := taskDurationsToIndices[task1Duration]
		task1Index := indicesWithTask1Duration[len(indicesWithTask1Duration)-1]
		taskDurationsToIndices[task1Duration] = indicesWithTask1Duration[:len(indicesWithTask1Duration)-1]

		task2SortedIndex := len(tasks) - 1 - idx
		task2Duration := sortedTasks[task2SortedIndex]
		indicesWithTask2Duration := taskDurationsToIndices[task2Duration]
		task2Index := indicesWithTask2Duration[len(indicesWithTask2Duration)-1]
		taskDurationsToIndices[task2Duration] = indicesWithTask2Duration[:len(indicesWithTask2Duration)-1]

		pairedTasks = append(pairedTasks, [2]int{task1Index, task2Index})
	}

	return pairedTasks
}

func getTaskDurationsToIndices(tasks []int) map[int][]int {
	taskDurationsToIndices := map[int][]int{}

	for idx, taskDuration := range tasks {
		if _, exists := taskDurationsToIndices[taskDuration]; !exists {
			taskDurationsToIndices[taskDuration] = []int{idx}
		} else {
			taskDurationsToIndices[taskDuration] = append(taskDurationsToIndices[taskDuration], idx)
		}
	}
	return taskDurationsToIndices
}
