package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name   string
	Dept   string
	Salary int
}

func GroupByDept(employees []Employee) map[string][]Employee {
	result := make(map[string][]Employee)
	for _, e := range employees {
		result[e.Dept] = append(result[e.Dept], e)
	}
	return result
}

func main() {
	employees := []Employee{
		{"Alice", "Sales", 100}, {"Bob", "Sales", 110}, {"Carol", "Sales", 115},
		{"Dave", "Sales", 105}, {"Eve", "Sales", 120}, {"Frank", "Sales", 125},
		{"Grace", "Sales", 130}, {"Heidi", "Sales", 110}, {"Ivan", "Sales", 115},
		{"Judy", "Sales", 140}, {"Kevin", "Sales", 135}, {"Laura", "Sales", 120},
		{"Mike", "Sales", 125}, {"Nancy", "Sales", 150}, {"Oscar", "Sales", 145},
		{"Peggy", "Sales", 130}, {"Quinn", "Sales", 135}, {"Rachel", "Sales", 160},
		{"Sam", "IT", 140}, {"Tina", "IT", 145}, {"Uma", "IT", 150},
		{"Victor", "IT", 155}, {"Wendy", "IT", 160}, {"Xavier", "IT", 140},
		{"Yara", "IT", 145}, {"Zach", "IT", 150}, {"Amy", "IT", 165},
		{"Brian", "IT", 170}, {"Chloe", "IT", 155}, {"Dan", "IT", 160},
		{"Leo", "Marketing", 110}, {"Mia", "Marketing", 115}, {"Noah", "Marketing", 120},
		{"Olivia", "Marketing", 125}, {"Paul", "Marketing", 130}, {"Ruby", "Marketing", 115},
		{"Steve", "Marketing", 120}, {"Tara", "Marketing", 135},
		{"Emma", "HR", 90}, {"Finn", "HR", 95}, {"Gina", "HR", 100},
		{"Hugo", "HR", 105}, {"Iris", "HR", 90}, {"Jack", "HR", 110}, {"Kate", "HR", 115},
		{"Ethan", "Engineering", 160}, {"Fiona", "Engineering", 165}, {"George", "Engineering", 170},
		{"Hannah", "Engineering", 175}, {"Ian", "Engineering", 180},
	}

	groups := GroupByDept(employees)

	depts := make([]string, 0, len(groups))
	for d := range groups {
		depts = append(depts, d)
	}
	sort.Slice(depts, func(i, j int) bool {
		if len(groups[depts[i]]) != len(groups[depts[j]]) {
			return len(groups[depts[i]]) > len(groups[depts[j]])
		}
		return depts[i] < depts[j]
	})

	for _, d := range depts {
		fmt.Printf("%s: %d человек\n", d, len(groups[d]))
	}
}
