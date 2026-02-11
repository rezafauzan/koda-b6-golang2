package main

import "fmt"

type best struct {
	best string
}
type the struct {
	the best
}
type are struct {
	are the
}
type we struct {
	we are
}

func main() {
	var we are = are{
		are: the{
			the: best{
				best: "Koda",
			},
		},
	}

	fmt.Println(we.are.the.best)
}
