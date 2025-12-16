package game

import (
	_ "embed"
	"fmt"
	"math/rand"
)

func Game(maxTurn int, wordList []string) bool {
	fmt.Printf("\nNew game\n")
	var charConfirmed [26]int
	randInt := rand.Intn(len(wordList))
	answer := wordList[randInt]
	turn := 0
	for turn < maxTurn {
		var input []rune
		fmt.Printf("[%d]Input your answer: ", turn+1)
		fmt.Scanln(&input)
		if valid := check(string(input), wordList); valid {
			turn++
			if correct := checkAnswer(string(input), answer, &charConfirmed); correct {
				fmt.Printf("Correct!\n")
				return true
			}
		} else {
			continue
		}
	}
	fmt.Printf("Game over! The answer is: %s\n", answer)
	return false
}

func check(input string, wordList []string) bool {
	wordLen := len(wordList[0])
	if len(input) != wordLen {
		fmt.Printf("The input must be exactly %d characters.\n", wordLen)
		return false
	}

	l := 0
	r := len(wordList)
	for l <= r {
		mid := l + (r-l)/2
		result := compare(input, wordList[mid])
		if result == -1 {
			r = mid - 1
		} else if result == 1 {
			l = mid + 1
		} else {
			return true
		}
	}
	fmt.Printf("Not a valid word.\n")
	return false
}

func compare(input string, wordList string) int {
	for i := 0; i < len(wordList); i++ {
		if input[i] < wordList[i] {
			return -1
		} else if input[i] > wordList[i] {
			return 1
		}
	}
	return 0
}

func checkAnswer(input string, answer string, charConfirmed *[26]int) bool {
	var charCount [26]int
	var correctCount = 0
	wordLen := len(answer)
	for i := 0; i < wordLen; i++ {
		if input[i] != answer[i] {
			charCount[answer[i]-'a']++
		}
	}
	for i := 0; i < wordLen; i++ {
		if input[i] == answer[i] {
			fmt.Printf("\033[42m%c\033[0m", input[i])
			charConfirmed[input[i]-'a'] = 3
			correctCount++
		} else {
			if charCount[input[i]-'a'] > 0 {
				charCount[input[i]-'a']--
				fmt.Printf("\033[43m%c\u001B[0m", input[i])
				charConfirmed[input[i]-'a'] = max(charConfirmed[input[i]-'a'], 2)
			} else {
				fmt.Printf("\033[100m%c\u001B[0m", input[i])
				charConfirmed[input[i]-'a'] = max(charConfirmed[input[i]-'a'], 1)
			}
		}
	}
	fmt.Printf("\nKnown: ")
	for i := 0; i < 26; i++ {
		switch charConfirmed[i] {
		case 0:
			fmt.Printf("%c", 'a'+i)
		case 1:
			fmt.Printf("\033[100m \u001B[0m")
		case 2:
			fmt.Printf("\033[43m%c\u001B[0m", 'a'+i)
		case 3:
			fmt.Printf("\033[42m%c\u001B[0m", 'a'+i)
		}
	}
	fmt.Printf("\n")
	if correctCount == wordLen {
		return true
	} else {
		return false
	}
}
