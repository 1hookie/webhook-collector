// Компиляция бинарников для разных ос
// GOOS = os GOARCH= arch go build -o main main.go

package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("Errors are values.")
	fmt.Println(getDays(12, 4, 2023, 13, 4, 2023))
	if col := getDays(12, 4, 2023, 13, 4, 2023); col != 1 {
		panic("Error#1")
	}
	fmt.Println(getDays(12, 3, 2023, 12, 4, 2023))
	if col := getDays(12, 3, 2023, 12, 4, 2023); col != 31 {
		panic("Error#2")
	}
	fmt.Println(getDays(112, 4, 1961, 12, 3, 2023))
	if col := getDays(12, 4, 1961, 12, 3, 2023); col != 22614 {
		panic("Error#3")
	}

}

// Альтернативный вариант решения задачи без использования пакета time
func checkYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

func yearToDays(year int) int {
	if checkYear(year) {
		return 366
	}
	return 365
}

func monthToDays(month, year int) int {
	months := make(map[int]int, 12)
	months[1] = 31
	months[2] = 28
	months[3] = 31
	months[4] = 30
	months[5] = 31
	months[6] = 30
	months[7] = 31
	months[8] = 31
	months[9] = 30
	months[10] = 31
	months[11] = 30
	months[12] = 31
	if month == 2 && checkYear(year) {
		return months[month] + 1
	}
	return months[month]
}

func dateToDays(d, m, y int) int {
	days := 0
	for year := 1; year < y; year++ {
		days += yearToDays(year)
	}
	for month := 1; month < m; month++ {
		days += monthToDays(month, y)
	}
	days += d
	return days
}

func getDays(d1, m1, y1, d2, m2, y2 int) int {
	days1 := dateToDays(d1, m1, y1)
	days2 := dateToDays(d2, m2, y2)
	return days2 - days1
}

// Альтернативный вариант решения задачи с использованием пакета time
func getDaysWithImport(d1, m1, y1, d2, m2, y2 int) int {
	date1 := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(y2, time.Month(m2), d2, 0, 0, 0, 0, time.UTC)
	diff := date2.Sub(date1)
	days := int(diff.Hours() / 24)
	return days
}
