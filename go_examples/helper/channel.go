package main

import (
    "fmt"
)

func sum(s []int, c chan int) {
    sum := 0
    for _, i := range s {
        sum += i
    }
    c <- sum
}

func main() {
    s :=  []int{10,20,30,40,50,60}
    fmt.Println(s[0:3])
    fmt.Println(s[3:6])
    c :=  make(chan int)
    go sum(s[0:3], c)
    go sum(s[3:6], c)
    x, y := <-c, <-c
    fmt.Println(x, y, x + y)
}