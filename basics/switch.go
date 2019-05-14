package main

import (
	system "fmt"
	"time"
)

func main() {
	var names = []string{"nayan", "kumar"}
	system.Println(time.Now().Weekday())
	for _, name := range names {
		switch name {
		case "nayan":
			system.Println("Nayan name is called...")
		case "kumar":
			system.Println("Kumar name is called...")
		default:
			system.Println("No name is found")
		}
	}

	switch time.Now().Weekday() {
	case time.Saturday:
		system.Println("Today is Saturday")
	case time.Sunday:
		system.Println("Today is Sunday")
	case time.Monday:
		system.Println("Today is Monday")
		fallthrough
	case time.Tuesday:
		system.Println("Today is Tuesday")
		fallthrough
	case time.Wednesday:
		system.Println("Today is Wednesday")
	}

	system.Println(time.Now())

	switch hour := time.Now().Hour(); {
	case hour < 12:
		system.Println("Good Morning...")
	case hour < 17:
		system.Println("Good Afternoon...")
	default:
		system.Println("Good Night...")
	}
}
