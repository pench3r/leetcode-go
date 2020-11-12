package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, queue, dept := getWordList(beginWord, wordList), []string{beginWord}, 0
	for len(queue) > 0 {
		dept++
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			word := queue[0]
			queue = queue[1:]
			candidates := getWordCandidates(word)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						return dept + 1
					}
					delete(wordMap, candidate)
					queue = append(queue, candidate)
				}
			}
		}
	}
	return 0
}

func getWordList(beginWord string, wordList []string) map[string]int {
	res := make(map[string]int)
	for i, word := range wordList {
		if _, ok := res[word]; !ok {
			if word != beginWord {
				res[word] = i
			}
		}
	}
	return res
}

func getWordCandidates(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(byte(int('a')+i))+word[j+1:])
			}
		}
	}
	return res
}
