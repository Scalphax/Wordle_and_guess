package guessv2

import (
	"fmt"
	"math"
)

type Solver struct {
	wordList  []string
	validList []string
	powConst  []int
	lastWord  string
	firstWord string
}

const (
	CORRECT = 2
	HAS     = 1
	NONE    = 0
)

func NewSolver(wordList []string) *Solver {
	wordLen := len(wordList[0])
	powConst := make([]int, wordLen)

	for i := 0; i < wordLen; i++ {
		powConst[i] = int(math.Pow(3, float64(i)))
	}
	s := Solver{
		wordList: wordList,
		powConst: powConst,
		lastWord: "",
	}
	s.Reset()
	s.setFirstWord()
	return &s
}

func (s *Solver) setFirstWord() {
	fmt.Printf("Finding first word...\n")
	bestE := -1.0
	bestWord := ""

	for _, candidate := range s.validList {
		E := s.calcE(candidate)
		if E > bestE {
			bestE = E
			bestWord = candidate
		}
	}
	s.firstWord = bestWord
	fmt.Printf("Found first word: %s\n", bestWord)
}

func (s *Solver) Reset() {
	validList := make([]string, len(s.wordList))
	copy(validList, s.wordList)
	s.validList = validList
	s.lastWord = ""
}

func (s *Solver) MakeChoice(charaState []byte) string {
	if len(s.validList) == 1 {
		return s.validList[0]
	}
	answer := s.firstWord
	if s.lastWord != "" {
		s.filterList(charaState)

		bestE := -1.0
		for _, candidate := range s.validList {
			E := s.calcE(candidate)
			if E > bestE {
				bestE = E
				answer = candidate
			}
		}
	}
	s.lastWord = answer
	return answer
}

func (s *Solver) calcE(candidate string) float64 {
	// 统计不同结果频次
	wordLen := len(s.wordList[0])
	counts := make([]byte, s.powConst[wordLen-1]*3)
	for _, word := range s.validList {
		charaState := checkAnswer(candidate, word)
		index := 0
		for i, state := range charaState {
			index += int(state) * s.powConst[wordLen-i-1]
		}
		counts[index]++
	}

	var E float64
	for _, count := range counts {
		//if i < 121 {
		//	continue
		//}
		p := float64(count) / float64(len(s.validList))
		//fmt.Printf("%f\n", p)
		if p > 0 {
			E += -p * math.Log2(p)
		}
	}
	return E
}

func (s *Solver) filterList(charaState []byte) {
	k := 0
	for _, word := range s.validList {
		if s.matchState(word, charaState) {
			s.validList[k] = word
			k++
		}
	}
	s.validList = s.validList[:k]
}

func (s *Solver) matchState(candidate string, charaState []byte) bool {
	cs := checkAnswer(s.lastWord, candidate)
	for i, state := range cs {
		if state != charaState[i] {
			return false
		}
	}
	return true
}

func checkAnswer(input string, answer string) []byte {
	var charCount [26]int
	wordLen := len(answer)
	charState := make([]byte, wordLen)
	for i := 0; i < wordLen; i++ {
		if input[i] != answer[i] {
			charCount[answer[i]-'a']++
		}
	}
	for i := 0; i < wordLen; i++ {
		if input[i] == answer[i] {
			charState[i] = CORRECT
		} else {
			if charCount[input[i]-'a'] > 0 {
				charCount[input[i]-'a']--
				charState[i] = HAS
			} else {
				charState[i] = NONE
			}
		}
	}
	return charState
}
