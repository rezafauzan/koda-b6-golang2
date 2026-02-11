package slice

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

type world struct {
	world string
}

type academy struct {
	academy string
}

type tech struct {
	tech academy
}

type man struct {
	man [1]tech
}

type obj struct {
	str [4][2][3]man
}

type is struct {
	is string
}

type fruit struct {
	fruit is
}

type favourite struct {
	favourite [4]fruit
}

type my struct {
	my [1]favourite
}

type num struct{
	first [2]int
	second [3]int
}

func Slice() {
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

	var obj obj = obj{
		str: [4][2][3]man{
			{
				{{}, {}, {}}, {{}, {}, {}},
			},
			{
				{{}, {}, {}}, {{}, {}, {}},
			},
			{
				{{}, {}, {}}, {{}, {}, {}},
			},
			{
				{{}, {}, {}},
				{
					{},
					{},
					man{[1]tech{{academy{academy: "Tech Academy"}}}},
				},
			},
		},
	}

	var my [1]favourite = [1]favourite{
		favourite{
			favourite: [4]fruit{
				{},
				{},
				{},
				{
					fruit: is{
						is: "Apple",
					}, 
				},
			},
		},
	}

	var num num = num{
		first: [2]int{
			0,16,
		},
		second: [3]int{
			0,0,16,
		},
	}

	fmt.Println(we.are.the.best)
	fmt.Println(hello.world)
	fmt.Println(obj.str[3][1][2].man[0].tech.academy)
	fmt.Println(my[0].favourite[3].fruit.is)
	fmt.Printf("%s \n", fmt.Sprintf("%d",num.first[1] + num.second[2]))
}
