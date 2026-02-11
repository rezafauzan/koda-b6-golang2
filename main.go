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

type world struct{
	world string
}

type academy struct{
	academy string
}

type tech struct{
	tech academy
}

type man struct{
	man [1]tech
}

type obj struct{
	str [4][2][3]man
}

func main() {
	var we are = are{
		are: the{
			the: best{
				best: "Koda",
			},
		},
	}

	var hello world = world{
		world: "Hello World",
	}

	var obj obj= obj{
		str: [4][2][3] man{
			{
				{{},{},{}},{{},{},{}},
			},
			{
				{{},{},{}},{{},{},{}},
			},
			{
				{{},{},{}},{{},{},{}},
			},
			{
				{{},{},{}},
				{
					{},
					{},
					man{[1]tech{tech{academy{academy: "Tech Academy"}}}},
				},
			},
		},
	}

	fmt.Println(we.are.the.best)
	fmt.Println(hello.world)
	fmt.Println(obj.str[3][1][2].man[0].tech.academy)
}
