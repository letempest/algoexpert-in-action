package dp

import (
	"sort"
)

type sortByHeight [][3]int

func (a sortByHeight) Len() int           { return len(a) }
func (a sortByHeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByHeight) Less(i, j int) bool { return a[i][2] < a[j][2] }

// Big O: O(n^2) time | O(n) space
func DiskStacking(disks [][3]int) [][3]int {
	sortedDisks := append([][3]int(nil), disks...) // keep a copy of input disks sorted by height
	sort.Sort(sortByHeight(sortedDisks))

	heights := make([]int, len(disks))
	for i := range heights {
		heights[i] = sortedDisks[i][2]
	}

	sequences := make([]int, len(disks))
	for i := range sequences {
		sequences[i] = -1
	}

	var maxHeightIdx int

	for i := 1; i < len(sortedDisks); i++ {
		currentDisk := sortedDisks[i]
		for j := 0; j < i; j++ {
			otherDisk := sortedDisks[j]
			if isStackable(otherDisk, currentDisk) {
				newStackHeightWithOtherDisk := currentDisk[2] + heights[j]
				if newStackHeightWithOtherDisk > heights[i] {
					heights[i] = newStackHeightWithOtherDisk
					sequences[i] = j
				}
			}
		}

		if heights[i] > heights[maxHeightIdx] {
			maxHeightIdx = i
		}
	}

	return buildSequence(sortedDisks, sequences, maxHeightIdx)
}

// whether dimensions of otherDisk are all strictly smaller than they are in currentDisk
func isStackable(o, c [3]int) bool {
	return o[0] < c[0] && o[1] < c[1] && o[2] < c[2]
}

func buildSequence(disks [][3]int, sequences []int, currentIdx int) [][3]int {
	sequence := make([][3]int, 0)
	for currentIdx != -1 {
		sequence = append(sequence, disks[currentIdx])
		currentIdx = sequences[currentIdx]
	}
	return reversed(sequence)
}

func reversed(array [][3]int) [][3]int {
	for i, j := 0, len(array)-1; i < j; {
		array[i], array[j] = array[j], array[i]
		i += 1
		j -= 1
	}
	return array
}
