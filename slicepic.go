package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy, dy)
	for i := range s {
		s[i] = make([]uint8, dx, dx)
		for j := range s[i] {
			s[i][j] = uint8((i+j)/2)
		}
	}
	return s
}

func main() {
	var p = Pic(8, 8)
	fmt.Println(p)
	pic.Show(Pic)
}
