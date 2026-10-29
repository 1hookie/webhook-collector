// Задача 3. «Пересечение двух массивов»

// Условие: Даны два слайса nums1 и nums2. Верни слайс их пересечения — элементы, которые встречаются в обоих массивах.
// Каждый элемент в результате должен быть уникальным (порядок не важен).

// Примеры:

// nums1 = [1, 2, 2, 1],  nums2 = [2, 2]       → [2]
// nums1 = [4, 9, 5],     nums2 = [9, 4, 9, 8] → [4, 9] (или [9, 4])

package main

import "fmt"

func intersectSlice(nums1 []int, nums2 []int) []int {
	// создаем мапу состоящую из элементов первого слайса
	seen := make(map[int]bool)
	for _, num := range nums1 {
		seen[num] = true
	}

	var result []int
	// проверяем есть ли элементы из второго слайса в мапе из элементов первого слайса
	// если есть -> добавляем в результирующий массив и удаляем из мапы, чтобы избежать дубликатов в результате
	for _, num := range nums2 {
		if seen[num] {
			result = append(result, num)
			delete(seen, num)
		}
	}
	return result
}

func main() {
	fmt.Println(intersectSlice([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersectSlice([]int{4, 9, 5}, []int{9, 4, 9, 8}))
}
