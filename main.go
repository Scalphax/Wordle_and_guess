package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"wordle/game"
	"wordle/guessv2"
)

//go:embed 2k_word_list.txt
var fileContent string

func main() {
	wordList := strings.Split(strings.TrimSpace(fileContent), "\n")
	f, err := os.Open("word_list.txt")
	if err != nil {
		fmt.Printf("word_list.txt not found, using embedded file.\n")
	} else {
		defer f.Close()
		wordList = nil
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			wordList = append(wordList, scanner.Text())
		}
	}

	fmt.Printf("0.Play 1.Robot 2.Full word list benchmark 3+.Random benchmark(int for turn)\nSelect mode: ")
	mode := 0
	fmt.Scanf("%d\n", &mode)

	run(mode, wordList)
}

func run(mode int, wordList []string) {
	switch mode {
	case 0:
		var maxTurn = 5
		print("Input max turn: ")
		_, err := fmt.Scanf("%d\n", &maxTurn)
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		for {
			game.Game(maxTurn, wordList)
		}
	case 1:
		s := guessv2.NewSolver(wordList)
		charaState := make([]byte, len(wordList[0]))
		for {
			fmt.Printf("%s\n", s.MakeChoice(charaState))
			var input string
			fmt.Scanln(&input)
			for i, num := range input {
				charaState[i] = byte(num - '0')
			}
			if input == "22222" {
				s.Reset()
			}
		}
	case 2:
		turnSum := 0
		gameRound := len(wordList)
		start := time.Now()
		s := guessv2.NewSolver(wordList)

		for _, answer := range wordList {
			turnSum += game.Auto(answer, wordList, s)
		}

		totalTime := time.Since(start)
		fmt.Printf("avgTurn: %f, avgTime: %fs", float32(turnSum)/float32(gameRound), totalTime.Seconds()/float64(gameRound))

	default:
		turnSum := 0
		gameRound := mode
		start := time.Now()
		s := guessv2.NewSolver(wordList)
		for i := 0; i < gameRound; i++ {
			randInt := rand.Intn(len(wordList))
			answer := wordList[randInt]
			turnSum += game.Auto(answer, wordList, s)
		}
		totalTime := time.Since(start)
		fmt.Printf("avgTurn: %f, avgTime: %fs", float32(turnSum)/float32(gameRound), totalTime.Seconds()/float64(gameRound))
	}
}
