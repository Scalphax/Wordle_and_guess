package guess

import (
	"container/list"
	"sort"
)

type Solver struct {
	wordWeight *list.List
	lastWord   string
}

func NewSolver(wordList []string) *Solver {
	frequency := countFreq(wordList)
	wordWeight := calcWeight(wordList, frequency)
	return &Solver{
		wordWeight: wordWeight,
	}
}

func (s *Solver) MakeChoice(charState []byte) string {
	if s.lastWord != "" {
		for i, char := range s.lastWord {
			switch charState[i] {
			case 3:
				s.removeElemIndexNot(char, i)
			case 2:
				s.removeElemIndex(char, i)
			case 1:
				remove := true
				for j, ch := range s.lastWord {
					if ch == char && charState[j] != 1 {
						remove = false
					}
				}
				if remove {
					s.removeElemContain(char)
				} else {
					s.removeElemIndex(char, i)
				}
			}
		}
	}
	answer := s.wordWeight.Front().Value.(kv).key
	s.lastWord = answer
	return answer
}

func (s *Solver) removeElemContain(targetChar rune) {
	for e := s.wordWeight.Front(); e != nil; {
		next := e.Next()
		word := e.Value.(kv).key
		for _, char := range word {
			if char == targetChar {
				s.wordWeight.Remove(e)
				break
			}
		}
		e = next
	}
}

func (s *Solver) removeElemIndex(targetChar rune, index int) {
	for e := s.wordWeight.Front(); e != nil; {
		next := e.Next()
		word := e.Value.(kv).key
		if rune(word[index]) == targetChar {
			s.wordWeight.Remove(e)
		}
		e = next
	}
}

func (s *Solver) removeElemIndexNot(targetChar rune, index int) {
	for e := s.wordWeight.Front(); e != nil; {
		next := e.Next()
		word := e.Value.(kv).key
		if rune(word[index]) != targetChar {
			s.wordWeight.Remove(e)
		}
		e = next
	}
}

func Guess(wordList []string) {
	//for e := wordWeight.Front(); e != nil; e = e.Next() {
	//	pair := e.Value.(kv)
	//	fmt.Printf("%s: %d\n", pair.key, pair.value)
	//}
}

type kv struct {
	key   string
	value int
}

func countFreq(wordList []string) map[rune]int {
	charCount := make(map[rune]int)
	for _, word := range wordList {
		var wordCount [26]bool
		for _, char := range word {
			if wordCount[char-'a'] == false {
				charCount[char]++
				wordCount[char-'a'] = true
			}
		}
	}
	//var list []kv
	//for i := 0; i < 26; i++ {
	//	list = append(list, kv{rune('a' + i), charCount[i]})
	//}
	//sort.Slice(list, func(i int, j int) bool {
	//	return list[i].value < list[j].value // 由小到大排序
	//})

	return charCount
}

func calcWeight(wordList []string, frequency map[rune]int) *list.List {
	var wordWeight []kv

	for _, word := range wordList {
		var wordCount [26]bool
		weight := 0
		for _, char := range word {
			if wordCount[char-'a'] == false {
				weight += frequency[char]
				wordCount[char-'a'] = true
			}
		}
		wordWeight = append(wordWeight, kv{word, weight})
	}

	sort.Slice(wordWeight, func(i int, j int) bool {
		return wordWeight[i].value > wordWeight[j].value
	})

	lst := list.New()
	for _, pair := range wordWeight {
		lst.PushBack(pair)
	}

	return lst
}
