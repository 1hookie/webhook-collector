package main

import (
	"fmt"
	"sort"
)

type Grade struct {
	Student string
	Subject string
	Score   int
}

type StudentResult struct {
	Name    string
	Average float64
}

type stats struct {
	Sum   int
	Count int
}

func TopStudents(grades []Grade, n int) []StudentResult {
	byStudent := make(map[string]stats)
	for _, g := range grades {
		s := byStudent[g.Student]
		s.Sum += g.Score
		s.Count++
		byStudent[g.Student] = s
	}

	results := make([]StudentResult, 0, len(byStudent))
	for name, s := range byStudent {
		results = append(results, StudentResult{
			Name:    name,
			Average: float64(s.Sum) / float64(s.Count),
		})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Average != results[j].Average {
			return results[i].Average > results[j].Average
		}
		return results[i].Name < results[j].Name
	})

	if n < 0 {
		n = 0
	}
	if len(results) > n {
		results = results[:n]
	}
	return results
}

func main() {
	grades := []Grade{
		{"Alice", "Math", 95}, {"Alice", "Physics", 90}, {"Alice", "Chemistry", 88},
		{"Alice", "Biology", 92}, {"Alice", "History", 94}, {"Alice", "English", 91}, {"Alice", "CS", 93},
		{"Bob", "Math", 75}, {"Bob", "Physics", 72}, {"Bob", "Chemistry", 78},
		{"Bob", "Biology", 70}, {"Bob", "History", 76}, {"Bob", "English", 74},
		{"Carol", "Math", 96}, {"Carol", "Physics", 92}, {"Carol", "Chemistry", 95},
		{"Carol", "Biology", 93}, {"Carol", "History", 94}, {"Carol", "English", 95},
		{"Dave", "Math", 65}, {"Dave", "Physics", 60}, {"Dave", "Chemistry", 70},
		{"Dave", "Biology", 68}, {"Dave", "History", 62},
		{"Eve", "Math", 90}, {"Eve", "Physics", 85}, {"Eve", "Chemistry", 88},
		{"Eve", "Biology", 92}, {"Eve", "History", 86},
		{"Frank", "Math", 70}, {"Frank", "Physics", 68}, {"Frank", "Chemistry", 72},
		{"Frank", "Biology", 75}, {"Frank", "History", 65},
		{"Grace", "Math", 80}, {"Grace", "Physics", 85}, {"Grace", "Chemistry", 82},
		{"Grace", "Biology", 78}, {"Grace", "History", 84},
		{"Henry", "Math", 78}, {"Henry", "Physics", 80}, {"Henry", "Chemistry", 75}, {"Henry", "Biology", 82},
		{"Ivy", "Math", 85}, {"Ivy", "Physics", 88}, {"Ivy", "Chemistry", 82}, {"Ivy", "Biology", 86},
		{"Jake", "Math", 60}, {"Jake", "Physics", 58}, {"Jake", "Chemistry", 62},
	}

	for i, r := range TopStudents(grades, 3) {
		fmt.Printf("%d. %s: средний %.1f\n", i+1, r.Name, r.Average)
	}
}
