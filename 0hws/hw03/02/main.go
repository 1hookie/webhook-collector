package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name     string
	Category string
	Price    float64
}

func MostExpensiveInCategory(products []Product) map[string]Product {
	best := make(map[string]Product)
	for _, p := range products {
		current, ok := best[p.Category]
		if !ok || p.Price > current.Price {
			best[p.Category] = p
		}
	}
	return best
}

func main() {
	products := []Product{
		{"iPhone 15", "Phones", 799}, {"iPhone 15 Pro", "Phones", 999}, {"iPhone 15 Pro Max", "Phones", 1199},
		{"iPhone 14", "Phones", 699}, {"Galaxy S24", "Phones", 799}, {"Galaxy S24 Ultra", "Phones", 1299},
		{"Galaxy S23", "Phones", 649}, {"Pixel 8", "Phones", 699}, {"Pixel 8 Pro", "Phones", 999},
		{"Pixel 7a", "Phones", 499}, {"OnePlus 12", "Phones", 799}, {"Xiaomi 14", "Phones", 899},
		{"Nothing Phone 2", "Phones", 599}, {"Sony Xperia 1 V", "Phones", 1099}, {"Motorola Edge 40", "Phones", 599},
		{"MacBook Pro 14", "Laptops", 1999}, {"MacBook Pro 16", "Laptops", 2499}, {"MacBook Air M3", "Laptops", 1099},
		{"Dell XPS 13", "Laptops", 1299}, {"Dell XPS 15", "Laptops", 1799}, {"ThinkPad X1 Carbon", "Laptops", 1699},
		{"ThinkPad T14", "Laptops", 1199}, {"HP Spectre x360", "Laptops", 1399}, {"ASUS ZenBook 14", "Laptops", 1099},
		{"Razer Blade 15", "Laptops", 2199}, {"Lenovo Yoga 9i", "Laptops", 1499}, {"Surface Laptop 5", "Laptops", 1299},
		{"iPad Pro 12.9", "Tablets", 1099}, {"iPad Air", "Tablets", 599}, {"iPad mini", "Tablets", 499},
		{"Galaxy Tab S9", "Tablets", 799}, {"Galaxy Tab S9 Ultra", "Tablets", 1199}, {"Surface Pro 9", "Tablets", 999},
		{"Xiaomi Pad 6", "Tablets", 349}, {"Lenovo Tab P12 Pro", "Tablets", 649}, {"Fire HD 10", "Tablets", 149},
		{"AirPods Pro", "Headphones", 249}, {"AirPods Max", "Headphones", 549}, {"Sony WH-1000XM5", "Headphones", 399},
		{"Bose QC Ultra", "Headphones", 429}, {"Sennheiser Momentum 4", "Headphones", 349}, {"JBL Tour One M2", "Headphones", 299},
		{"Beats Studio Pro", "Headphones", 349}, {"B&W Px8", "Headphones", 699},
		{"Apple Watch Ultra 2", "Smartwatches", 799}, {"Apple Watch Series 9", "Smartwatches", 399},
		{"Galaxy Watch 6", "Smartwatches", 329}, {"Garmin Fenix 7", "Smartwatches", 699},
		{"Fitbit Sense 2", "Smartwatches", 249}, {"Pixel Watch 2", "Smartwatches", 349},
	}

	best := MostExpensiveInCategory(products)

	categories := make([]string, 0, len(best))
	for c := range best {
		categories = append(categories, c)
	}
	sort.Strings(categories)

	for _, c := range categories {
		p := best[c]
		fmt.Printf("%-13s %s — $%.0f\n", c+":", p.Name, p.Price)
	}
}
