// Задача 4. «Изоморфные строки»

// Условие: Даны две строки s и t. Они изоморфны, если каждый символ из s
//  можно однозначно заменить символом из t (и наоборот — без коллизий). Верни true или false.

// Примеры:

// s = "egg",   t = "add"   → true   (e→a, g→d)
// s = "foo",   t = "bar"   → false  (o не может маппиться и в b, и в r)
// s = "paper", t = "title" → true   (p→t, a→i, e→l, r→e)
// s = "badc",  t = "baba"  → false  (a и c оба маппятся в a — коллизия)

package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sr, tr := []rune(s), []rune(t)
	if len(sr) != len(tr) {
		return false
	}

	sToT := make(map[rune]rune)
	tToS := make(map[rune]rune)

	for i := range sr {
		a, b := sr[i], tr[i]

		if mapped, ok := sToT[a]; ok && mapped != b {
			return false
		}
		if mapped, ok := tToS[b]; ok && mapped != a {
			return false
		}

		sToT[a] = b
		tToS[b] = a
	}
	return true
}

func main() {
	if !isIsomorphic("egg", "add") {
		panic("Error: Test #1 is failed")
	}

	if isIsomorphic("foo", "bar") {
		panic("Error: Test #2 is failed")
	}

	if !isIsomorphic("paper", "title") {
		panic("Error: Test #3 is failed")
	}

	if isIsomorphic("badc", "baba") {
		panic("Error: Test #4 is failed")
	}

	fmt.Println("All tests passed")
}
