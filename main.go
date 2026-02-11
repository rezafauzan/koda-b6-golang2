package main

import "fmt"

func main() {
	scores := []int{50, 75, 66, 20, 32, 90,}

	index66 := 0
	for x := range len(scores) {
		if(scores[x] == 66){
			index66 = x
		}
	}

	var newScores []int = []int{}
	for y := range index66+1{
		newScores = append(newScores, scores[y])
	}
	newScores = append(newScores, 88)
	
	var restScores []int = scores[index66+1:]
	for z := range len(restScores){
		newScores = append(newScores, restScores[z])
	}
	for a := range len(newScores){
		fmt.Printf("Scores pada index - %d %s \n", a, fmt.Sprintf("%d", newScores[a]))
	}
}