package testdelete

import "fmt"

func main() {
	x := 20
	y := 20
	fmt.Println(sum(x, y))
}

func sum(x, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}
