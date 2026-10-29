package main

import (
	"fmt"
)

//func isPalindrome(n int) bool {
//	nStr := strconv.Itoa(n)
//	for i := 0; i < len(nStr)/2; i++ {
//		if nStr[i] != nStr[len(nStr)-i-1] {
//			return false
//		}
//	}
//	return true
//}

func isPalindrome(n int) bool {
	if (n < 0) || (n%10 == 0 && n == 0) {
		return false
	}

	reverted := 0
	for n > reverted {
		reverted = reverted*10 + n%10
		n /= 10
	}
	return n == reverted
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(33))
	fmt.Println(isPalindrome(12344321))
	fmt.Println(isPalindrome(12345678))
}
