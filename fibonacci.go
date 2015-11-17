package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	next, current := 1, 0
	return func() int {
		s := next + current;
		defer func() {	
			current = next
		   	next = s
		}()
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
