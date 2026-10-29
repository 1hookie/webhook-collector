// Условие: Дана строка s. Найди индекс первого символа, который встречается в строке ровно один раз. Если такого символа нет — верни -1.

// Примеры:

// "lqwerewqe"     → 0  (символ 'l' встречается 1 раз)
// "qowertqwert" → 1  (символ 'o' встречается 1 раз)
// "aabb"         → -1

package main

import "fmt"

func main() {
	if res := findUniqueChar("lqwerewqe"); res != 0 {
		panic("Error: Test #1 is faild")
	}

	if res := findUniqueChar("qowertqwert"); res != 1 {
		panic("Error: Test #2 is faild")
	}

	if res := findUniqueChar("aabb"); res != -1 {
		panic("Error: Test #3 is faild")
	}

	fmt.Println("All tests passed")
}

func findUniqueChar(s string) int {
	runesCount := make(map[rune]int)

	for _, r := range s {
		runesCount[r]++
	}

	for i, r := range s {
		if runesCount[r] == 1 {
			return i
		}
	}
	return -1
}
