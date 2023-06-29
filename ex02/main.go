package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var play []string
	var points int

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	formattedData := strings.Split(string(data), "\n")
	for i := 0; i < len(formattedData)-1; i++ {
		play = strings.Split(formattedData[i], " ")
		points += myPointsForShape(play) + enemyPointsForShape(play) + draw(play) + win(play) + lose(play)

	}
	if points > 0 {
		fmt.Println("You won the game with", points, "points ahead!")
	} else if points < 0 {
		fmt.Println("You lost the game with", points, "points behind!")
	} else {
		fmt.Println("The game finished with a draw!")
	}
}

//A  X -> Rock
//B  Y -> Paper
//C  Z -> Scissors

func win(play []string) int {
	if (play[0] == "A" && play[1] == "Y") ||
		(play[0] == "B" && play[1] == "Z") ||
		(play[0] == "C" && play[1] == "X") {

		return 6
	}
	return 0
}
func lose(play []string) int {
	if (play[0] == "A" && play[1] == "Z") ||
		(play[0] == "B" && play[1] == "X") ||
		(play[0] == "C" && play[1] == "Y") {
		return -6
	}
	return 0
}

func draw(play []string) int {
	if (play[0] == "A" && play[1] == "X") ||
		(play[0] == "B" && play[1] == "Y") ||
		(play[0] == "C" && play[1] == "Z") {
		return 0
	}
	return 0
}

func myPointsForShape(play []string) int {
	if play[1] == "X" {
		return 1
	}
	if play[1] == "Y" {
		return 2
	}
	return 3
}

func enemyPointsForShape(play []string) int {
	if play[0] == "A" {
		return 1 * -1
	}
	if play[0] == "B" {
		return 2 * -1
	}
	return 3 * -1
}
