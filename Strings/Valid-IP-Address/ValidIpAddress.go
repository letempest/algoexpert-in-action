// Given a string with digits, find out all valid ip addresses that could be formed
// by separating the string with dot delimiter
// O(1) Time | O(1) Space; Reason: both time and space of possible ip address has a upper bound O(2^32),
// regardless of what input is provided
package main

import (
	"fmt"
	"math"
	"strconv"
)

const ip = "1921680"

func main() {
	fmt.Printf("%#v\n", getValidIpList(ip))
	fmt.Printf("%#v\n", getValidIpList("2125074"))
	fmt.Printf("%#v\n", getValidIpList("11112"))
	fmt.Printf("%#v\n", getValidIpList("2553072"))
	fmt.Printf("%#v\n", getValidIpList("035345"))
}

func isValidPart(part string) bool {
	numPart, err := strconv.ParseInt(part, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if numPart > 255 {
		return false
	}
	return len(fmt.Sprint(numPart)) == len(part)
}

func getValidIpList(ipStr string) []string {
	var ipList []string
	var currentIP string
	firstIdx, secondIdx, thirdIdx := 1, 0, 0
	var first, second, third, fourth string

	if string(ipStr[0]) == "0" {
		return ipList
	}

	for firstIdx < int(math.Min(4, float64(len(ipStr))-2)) {
		first = ipStr[0:firstIdx]
		isValid := isValidPart(first)
		if !isValid {
			break
		}

		secondIdx = firstIdx + 1
		thirdIdx = secondIdx + 1
		for secondIdx < int(math.Min(7, float64(len(ipStr))-1)) {
			second = ipStr[firstIdx:secondIdx]
			isValid = isValidPart(second)
			if !isValid {
				break
			}

			thirdIdx = secondIdx + 1
			for thirdIdx < len(ipStr) {
				third = ipStr[secondIdx:thirdIdx]
				fourth = ipStr[thirdIdx:]

				if isValidPart(third) && isValidPart(fourth) {
					currentIP = fmt.Sprintf("%s.%s.%s.%s", first, second, third, fourth)
					// fmt.Println(currentIP)
					ipList = append(ipList, currentIP)
				}
				thirdIdx++
			}
			secondIdx++
		}
		firstIdx++
	}
	return ipList
}
