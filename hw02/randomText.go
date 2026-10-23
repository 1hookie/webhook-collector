package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	var (
		m      = make(map[byte]int)
		rare   = make([]byte, 0)
		common = make([]byte, 0)
	)

	str := randomString(1000)
	fmt.Println(str)
	// подсчитываем количество каждого символа в строке
	for _, v := range str {
		m[byte(v)]++
	}

	// добавляем в слайсы редкие и частые символы
	for k, v := range m {
		switch {
		case v < 15:
			rare = append(rare, k)
		case v >= 20:
			common = append(common, k)
		}
	}

	// сортируем слайсы
	slices.Sort(rare)
	slices.Sort(common)

	// выводим результаты
	fmt.Println("Unique chars:", len(m))
	fmt.Println("Rare chars:", string(rare))
	fmt.Println("Common chars:", string(common))

	//топ 3 биграммы
	top3Bigramms(str)

	// - Сгенерировать две случайные строки по 11 символу. Используя map как множество, найти:
	// - Пересечение — символы, которые есть в обеих строках.
	// - Разность A \ B — символы, которые есть в первой, но нет во второй.
	// - Симметрическую разность — символы, которые есть только в одной из строк (но не в обеих сразу).
	// - Для каждого символа X посчитать, сколько разных символов встречается сразу после него в строке. Символ, после которого идёт наибольшее разнообразие — "самый непредсказуемый".
	randStr1 := randomString(11)
	randStr2 := randomString(11)

	fmt.Println("Random String 1:", randStr1)
	fmt.Println("Random String 2:", randStr2)

	set1 := make(map[rune]struct{})
	set2 := make(map[rune]struct{})

	for _, char := range randStr1 {
		set1[char] = struct{}{}
	}
	for _, char := range randStr2 {
		set2[char] = struct{}{}
	}
	fmt.Println(set1, set2)

}

func top3Bigramms(str string) {
	// создаем мапу типа
	// биграмма : кол-во в строке
	bigramms := make(map[string]int)

	for i := 0; i < len(str)-1; i++ {
		bigramm := str[i : i+2]
		bigramms[bigramm]++
	}

	// слайс по ключам мапы
	keys := make([]string, 0, len(bigramms))

	for k := range bigramms {
		keys = append(keys, k)
	}

	// сортироуем по убыванию частоты встречаемости биграмм
	slices.SortFunc(keys, func(a, b string) int {
		if bigramms[a] != bigramms[b] {
			return bigramms[b] - bigramms[a]
		}
		return strings.Compare(a, b)
	})

	for i := 0; i < 3; i++ {
		fmt.Printf("Top %d: %s - %d\n", i+1, keys[i], bigramms[keys[i]])
	}
}

func randomString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
