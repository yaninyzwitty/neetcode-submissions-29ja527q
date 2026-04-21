func ladderLength(beginWord string, endWord string, wordList []string) int {
	// make a wordMap for o(1) operations
	wordMap := make(map[string]bool)
	for _, word := range wordList {
		wordMap[word] = true
	}
	// if the endWord is already in the map, early return
	if !wordMap[endWord] {
		return 0
	}

	queue := []string{beginWord}
	steps := 1

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]

			if word == endWord {
				return steps
			}

			for j := 0; j < len(word); j++ {
				for c := 'a'; c <= 'z'; c++ {

					// construct the next word
					next := word[:j] + string(c) + word[j+1:]

					if wordMap[next] {
						queue = append(queue, next)
						// deque
						delete(wordMap, next)
					}
				}
			}

			
		}
		steps++

	}
    return 0
}
