package main

import "fmt"
import "koda-b6-golang2/slice"

func main() {
	slice.Slice()
	scores := []int{50, 75, 20, 32, 90, 66,}
	var newScores []int = []int{}

	index66 := 0

	for x := range len(scores) {
		if(scores[x] == 66){
			index66 = x
		}
	}
	
	
	for y := range index66+1{
		if(y == index66){
			newScores = append(newScores, 88)
		}
		newScores = append(newScores, scores[y])
	}

	
	var restScores []int = scores[index66+1:]

	newScores = append(newScores, restScores...)

	for a := range len(newScores){
		fmt.Printf("Scores pada index - %d %s \n", a, fmt.Sprintf("%d", newScores[a]))
	}
}