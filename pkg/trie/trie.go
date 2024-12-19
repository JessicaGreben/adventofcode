package trie

import "fmt"

const (
	wordEnding rune = '*'
	start      rune = '0'
)

type trieLevel map[rune]trieLevel

func (t trieLevel) Print() {
	for k, v := range t {
		fmt.Printf("k=%v\n", string(k))
		v.Print()
	}
}

type Trie struct {
	data trieLevel
}

func New() Trie {
	return Trie{
		data: trieLevel{
			start: trieLevel{},
		},
	}
}

func (t *Trie) Print() {
	t.data.Print()
}

func (t *Trie) Insert(word string) {
	c := t.data[start]
	for _, r := range word {
		if _, ok := c[r]; !ok {
			c[r] = trieLevel{}
		}
		c = c[r]
	}
	if _, ok := c[wordEnding]; !ok {
		c[wordEnding] = trieLevel{}
	}
}

func (t *Trie) Search(word string) bool {
	found, isEnd := t.find(word)
	return found && isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	found, _ := t.find(prefix)
	return found
}

func (t *Trie) find(word string) (found bool, isEnd bool) {
	if len(word) == 0 {
		return true, true
	}
	c := t.data[start]
	for _, r := range word {
		if _, ok := c[r]; !ok {
			return false, false
		}
		c = c[r]
	}
	_, ok := c[wordEnding]
	return true, ok
}
