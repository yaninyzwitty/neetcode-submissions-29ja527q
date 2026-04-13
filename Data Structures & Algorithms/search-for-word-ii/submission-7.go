type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

func buildTrie(words []string) *TrieNode {
	root := NewTrieNode()

	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			ch := word[i]

			if node.children[ch] == nil {
				node.children[ch] = NewTrieNode()
			}
			node = node.children[ch]
		}
		node.word = word
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	root := buildTrie(words)

	var result []string
	rows, cols := len(board), len(board[0])

	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		ch := board[i][j]
		nextNode := node.children[ch]

		if nextNode == nil {
			return
		}

		if nextNode.word != "" {
			result = append(result, nextNode.word)
			nextNode.word = "" // avoid duplicates
		}

		board[i][j] = '#'

		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range directions {
			nI, nJ := i+dir[0], j+dir[1] 
			if nI >= 0 && nI < rows && nJ >= 0 && nJ < cols && board[nI][nJ] != '#' {
				dfs(nI, nJ, nextNode)
			}
		}

		board[i][j] = ch
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, root)
		}
	}

	return result
}