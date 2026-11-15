package main

import (
	"fmt"
	"sort"
)

type Order struct {
	ClientID int
	Amount   float64
}

type ClientSummary struct {
	ClientID   int
	TotalSpent float64
	OrderCount int
}

func SummarizeOrders(orders []Order) []ClientSummary {
	byClient := make(map[int]*ClientSummary)
	for _, o := range orders {
		s, ok := byClient[o.ClientID]
		if !ok {
			s = &ClientSummary{ClientID: o.ClientID}
			byClient[o.ClientID] = s
		}
		s.TotalSpent += o.Amount
		s.OrderCount++
	}

	result := make([]ClientSummary, 0, len(byClient))
	for _, s := range byClient {
		result = append(result, *s)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ClientID < result[j].ClientID
	})
	return result
}

func ordersWord(n int) string {
	switch {
	case n%100 >= 11 && n%100 <= 14:
		return "заказов"
	case n%10 == 1:
		return "заказ"
	case n%10 >= 2 && n%10 <= 4:
		return "заказа"
	default:
		return "заказов"
	}
}

func main() {
	orders := []Order{
		{1, 150}, {2, 200}, {3, 75}, {1, 300}, {4, 125},
		{5, 180}, {6, 90}, {7, 250}, {8, 110}, {9, 400},
		{10, 60}, {11, 175}, {12, 220}, {1, 125}, {2, 350},
		{3, 200}, {4, 180}, {5, 95}, {6, 140}, {7, 300},
		{8, 225}, {9, 150}, {10, 280}, {11, 110}, {12, 175},
		{1, 400}, {2, 125}, {3, 90}, {4, 250}, {5, 200},
		{6, 180}, {7, 150}, {8, 95}, {9, 220}, {10, 125},
		{11, 300}, {1, 175}, {2, 280}, {3, 150}, {4, 110},
		{5, 225}, {6, 250}, {7, 180}, {8, 300}, {1, 225},
		{2, 150}, {3, 175}, {4, 200}, {1, 100}, {2, 225},
	}

	for _, s := range SummarizeOrders(orders) {
		fmt.Printf("Клиент %d: %d %s, всего %.0f\n",
			s.ClientID, s.OrderCount, ordersWord(s.OrderCount), s.TotalSpent)
	}
}
