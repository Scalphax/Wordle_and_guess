package guess

import (
	_ "embed"
	"sort"
)

func guess(wordList []string) {
	frequency := calc(wordList)
}

type kv struct {
	key   byte
	value int
}

func calc(wordList []string) []kv {
	var count [26]int
	for _, word := range wordList {
		var wordCount [26]int
		for _, char := range word {
			if wordCount[char-'a'] == 0 {
				count[char-'a']++
				wordCount[char-'a']++
			}
		}
	}
	var list []kv
	for i := 0; i < 26; i++ {
		list = append(list, kv{byte('a' + i), count[i]})
	}
	sort.Slice(list, func(i int, j int) bool {
		return list[i].value > list[j].value
	})
	return list
}
