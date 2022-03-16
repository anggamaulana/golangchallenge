package main

import (
	"fmt"
	"time"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func possibleTimes(digits []int) int {
	// Your code here
	var result []string
	str := fmt.Sprintf("%d%d%d%d", digits[0], digits[1], digits[2], digits[3])

	Perm([]rune(str), func(a []rune) {
		input := string(a[:2]) + ":" + string(a[2:])
		format := "15:04"
		_, err := time.Parse(format, input)
		if err != nil {
			return
		} else {
			result = append(result, input)
		}
	})

	return len(result)
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
}
