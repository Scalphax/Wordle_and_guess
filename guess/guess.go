package guess

import (
	"container/list"
	"fmt"
	"sort"
)

func Guess(wordList []string) {
	frequency := countFreq(wordList)
	wordWeight := calcWeight(wordList, frequency)
	for e := wordWeight.Front(); e != nil; e = e.Next() {
		pair := e.Value.(kv)
		fmt.Printf("%s: %d\n", pair.key, pair.value)
	}
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
