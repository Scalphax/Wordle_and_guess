package main

import (
	"bufio"
	"fmt"
	"os"
	"wordle/game"
)

func main() {
	f, _ := os.Open("wordle-answers-alphabetical.txt")
	defer f.Close()

	var wordList []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	turnSum := 0
	gameRound := 4000
	for i := 0; i < gameRound; i++ {
		turnSum += game.Auto(wordList)
	}
	fmt.Printf("avgTurn: %f", float32(turnSum)/float32(gameRound))

	//var maxTurn = 5
	//print("Input max turn: ")
	//_, err := fmt.Scanf("%d\n", &maxTurn)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//
	//for {
	//	game.Game(maxTurn, wordList)
	//}
}
