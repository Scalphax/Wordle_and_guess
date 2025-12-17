package guessv2

import (
	"math"
)

type Solver struct {
	wordList  []string
	validList []string
	powConst  []int
	wordLen   int
	lastWord  string
}

func NewSolver(wordList []string) *Solver {
	validList := make([]string, len(wordList))
	wordLen := len(wordList[0])
	powConst := make([]int, wordLen)

	for i := 0; i < wordLen; i++ {
		powConst[i] = int(math.Pow(3, float64(i)))
	}
	s := Solver{
		wordList:  wordList,
		validList: validList,
		powConst:  powConst,
		wordLen:   wordLen,
		lastWord:  "",
	}
	s.Reset()
	return &s
}

func (s *Solver) Reset() {
	copy(s.validList, s.wordList)
	s.lastWord = ""
}

func (s *Solver) MakeChoice(charaState []byte) string {
	if len(s.validList) == 1 {
		return s.validList[0]
	}
	answer := "trace"
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
	var counts [364]byte
	for _, word := range s.validList {
		charaState := checkAnswer(candidate, word)
		index := 0
		for i, state := range charaState {
			index += int(state) * s.powConst[s.wordLen-i-1]
		}
		counts[index]++
	}

	var E float64
	for i, count := range counts {
		if i < 121 {
			continue
		}
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
			charState[i] = 3
		} else {
			if charCount[input[i]-'a'] > 0 {
				charCount[input[i]-'a']--
				charState[i] = 2
			} else {
				charState[i] = 1
			}
		}
	}
	return charState
}
