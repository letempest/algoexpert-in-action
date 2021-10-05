package recursion

// Big O: O(n) time | O(d) space
// where n is number of all the elements in the array (including all the nested elements in sub arrays)
// d is the depth of the deepest subarray (e.g. for [1, 2, [3, [4, [5, 6], 7]], 8], d = 4)
// d recursive calls on the callstack
func ProductSum(array []interface{}) int {
	return productSumHelper(array, 1)
}

func productSumHelper(array []interface{}, multiplyer int) int {
	sum := 0
	for _, element := range array {
		switch v := element.(type) {
		case int:
			sum += v
		case []interface{}:
			// mixed case, element could be in form [1, 2, [3, [4, 5]], 6]
			sum += productSumHelper(v, multiplyer+1)
		case []int:
			// element contains only integer, the comented code below
			// use additional space, just to make golang's type checker happy
			// convert []int back to []interface{}
			// var nestedArr []interface{}
			// for _, num := range v {
			// 	nestedArr = append(nestedArr, num)
			// }
			// sum += productSumHelper(nestedArr, multiplyer+1)
			nestedSum := 0
			for _, num := range v {
				nestedSum += num
			}
			sum += nestedSum * (multiplyer + 1)
		default:
			panic("You should never reach here")
		}
	}

	return sum * multiplyer
}
