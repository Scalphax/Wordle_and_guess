package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"wordle/guessv2"
)

func main() {
	f, _ := os.Open("word_list.txt")
	defer f.Close()

	var wordList []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	turnSum := 0
	gameRound := len(wordList)
	start := time.Now()
	guessv2.NewSolver(wordList)
	//for _, answer := range wordList {
	//	//randInt := rand.Intn(len(wordList))
	//	//answer := wordList[randInt]
	//	turnSum += game.Auto(answer, wordList)
	//}
	time := time.Since(start)
	fmt.Printf("avgTurn: %f, avgTime: %fs", float32(turnSum)/float32(gameRound), time.Seconds()/float64(gameRound))

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
