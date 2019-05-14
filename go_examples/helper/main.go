package main

import (
	"fmt"
	"util"
)

func getFullName(fname, lname string) string {
	return "Hello " + fname + " " + lname
}

func getRange() []string {
	names := []string{"nayan", "kumar", "hn"}
	names = names[0:2]
	return names
}

type User struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	res := getFullName("nayan", "kumar")
	fmt.Println(res)
	fmt.Println(getRange())
	a := 10
	b := &a
	fmt.Println(a)
	user := User{"Nayan", 27, "M"}

	customer := &user
	fmt.Println(*customer)
	fmt.Println(&a)
	fmt.Println(&b)

	util.testCode("nayan")
}
