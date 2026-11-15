// Условие: Дан слайс целых чисел. Раздели его на две группы: чётные и нечётные. Верни map[bool][]int, где ключ — true для чётных, false для нечётных. Порядок внутри групп не важен.

// Примеры:

// [1, 2, 3, 4, 5, 6] → {false: [1, 3, 5], true: [2, 4, 6]}
// []                 → {false: [], true: []}
// [7]                → {false: [7], true: []}
package main

import (
	"fmt"
	"slices"
)

func splitSlice(slice []int) map[bool][]int {
	resultMap := map[bool][]int{
		true:  {},
		false: {},
	}

	for _, n := range slice {
		even := n%2 == 0
		resultMap[even] = append(resultMap[even], n)
	}
	return resultMap
}

func main() {
	if res := splitSlice([]int{1, 2, 3, 4, 5, 6}); !slices.Equal(res[false], []int{1, 3, 5}) || !slices.Equal(res[true], []int{2, 4, 6}) {
		panic("Error: Test #1 is failed")
	}

	if res := splitSlice([]int{}); res[false] == nil || res[true] == nil || len(res[false]) != 0 || len(res[true]) != 0 {
		panic("Error: Test #2 is failed")
	}

	if res := splitSlice([]int{7}); !slices.Equal(res[false], []int{7}) || res[true] == nil || len(res[true]) != 0 {
		panic("Error: Test #3 is failed")
	}

	fmt.Println("All tests passed")
}
