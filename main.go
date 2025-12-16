package main

import (
	"bufio"
	"fmt"
	"os"
	"wordle/game"
	"wordle/guess"
)

func main() {
	f, _ := os.Open("wordle-answers-alphabetical.txt")
	defer f.Close()

	var wordList []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	guess.Guess(wordList)
	return

	var maxTurn = 5
	print("Input max turn: ")
	_, err := fmt.Scanf("%d\n", &maxTurn)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for {
		game.Game(maxTurn, wordList)
	}
}
