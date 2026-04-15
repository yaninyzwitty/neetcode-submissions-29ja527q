type WordDictionary struct {
    children map[byte]*WordDictionary
    isEnd    bool
}

func Constructor() *WordDictionary {
    return &WordDictionary{children: make(map[byte]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
    node := this
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if _, ok := node.children[ch]; !ok {
            node.children[ch] = &WordDictionary{
                children: make(map[byte]*WordDictionary),
            }
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.searchHelper(word, 0)
}

func (this *WordDictionary) searchHelper(word string, index int) bool {
    if index == len(word) {
        return this.isEnd
    }

    ch := word[index]

    if ch == '.' {
        for _, child := range this.children {
            if child.searchHelper(word, index+1) {
                return true
            }
        }
        return false
    }

    child, ok := this.children[ch]
    if !ok {
        return false
    }

    return child.searchHelper(word, index+1)
}