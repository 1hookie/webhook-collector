// Задача 5. «Два числа, дающие сумму» (Two Sum)

// Условие: Дан слайс целых чисел nums и число target. Найди индексы двух элементов, сумма которых равна target. Гарантируется, что решение существует ровно одно. Один и тот же элемент использовать дважды нельзя.

// Примеры:

// nums = [2, 7, 11, 15], target = 9  → [0, 1]   (2 + 7 = 9)
// nums = [3, 2, 4],      target = 6  → [1, 2]   (2 + 4 = 6)
// nums = [3, 3],         target = 6  → [0, 1]   (3 + 3 = 6)

package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return nil
}

func main() {
	if res := twoSum([]int{2, 7, 11, 15}, 9); !slices.Equal(res, []int{0, 1}) {
		panic("Error: Test #1 is failed")
	}

	if res := twoSum([]int{3, 2, 4}, 6); !slices.Equal(res, []int{1, 2}) {
		panic("Error: Test #2 is failed")
	}

	if res := twoSum([]int{3, 3}, 6); !slices.Equal(res, []int{0, 1}) {
		panic("Error: Test #3 is failed")
	}

	fmt.Println("All tests passed")
}
