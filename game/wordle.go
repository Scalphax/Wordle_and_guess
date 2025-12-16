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
	fmt.Printf("%s", answer)
	turn := 0
	for turn < maxTurn {
		var input string
		fmt.Printf("[%d]Input your answer: ", turn+1)
		fmt.Scanln(&input)
		if valid := check(input, wordList); valid {
			turn++
			charState := checkAnswer(input, answer, &charConfirmed)
			printCharState(charState, input)
			printCharConfirmed(&charConfirmed)
			correct := true
			for _, i := range charState {
				if i != 3 {
					correct = false
					break
				}
			}
			if correct {
				fmt.Printf("Correct!\n")
				return true
			}
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

func checkAnswer(input string, answer string, charConfirmed *[26]int) []byte {
	var charCount [26]int
	var correctCount = 0
	wordLen := len(answer)
	charState := make([]byte, wordLen)
	for i := 0; i < wordLen; i++ {
		if input[i] != answer[i] {
			charCount[answer[i]-'a']++
		}
	}
	for i := 0; i < wordLen; i++ {
		if input[i] == answer[i] {
			//fmt.Printf("\033[42m%c\033[0m", input[i])
			charState[i] = 3
			charConfirmed[input[i]-'a'] = 3
			correctCount++
		} else {
			if charCount[input[i]-'a'] > 0 {
				charCount[input[i]-'a']--
				//fmt.Printf("\033[43m%c\u001B[0m", input[i])
				charState[i] = 2
				charConfirmed[input[i]-'a'] = max(charConfirmed[input[i]-'a'], 2)
			} else {
				//fmt.Printf("\033[100m%c\u001B[0m", input[i])
				charState[i] = 1
				charConfirmed[input[i]-'a'] = max(charConfirmed[input[i]-'a'], 1)
			}
		}
	}
	return charState
}

func printCharState(charState []byte, word string) {
	for i, char := range word {
		switch charState[i] {
		case 1:
			fmt.Printf("\033[100m%c\u001B[0m", char)
		case 2:
			fmt.Printf("\033[43m%c\u001B[0m", char)
		case 3:
			fmt.Printf("\033[42m%c\033[0m", char)
		}
	}
	fmt.Printf("\n")
}

func printCharConfirmed(charConfirmed *[26]int) {
	fmt.Printf("Known: ")
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
}
