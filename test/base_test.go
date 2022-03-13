package test

import (
	"fmt"
	"testing"
)
func changeLocal(num []int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
func TestBase(t *testing.T)  {
	num := []int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}

