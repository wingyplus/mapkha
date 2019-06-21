package mapkha

import (
	"sort"
)

// WordWithPayload is a pair of word and its payload
type WordWithPayload struct {
	Word    string
	Payload interface{}
}

// PrefixTreeNode represents node in a prefix tree
type PrefixTreeNode struct {
	NodeID int
	Offset int
	Ch     rune
}

// PrefixTreePointer is partial information of edge
type PrefixTreePointer struct {
	ChildID int
	IsFinal bool
	Payload interface{}
}

// PrefixTree is a Hash-based Prefix Tree for searching words
type PrefixTree struct {
	tab map[PrefixTreeNode]*PrefixTreePointer
}

type byWord WordsWithPayload

func (payload byWord) Len() int {
	return payload.Length
}

func (payload byWord) Swap(i, j int) {
	payload.Word[i], payload.Word[j] = payload.Word[j], payload.Word[i]
	payload.Payload[i], payload.Payload[j] = payload.Payload[j], payload.Payload[i]
}

func (payload byWord) Less(i, j int) bool {
	return payload.Word[i] < payload.Word[j]
}

type WordsWithPayload struct {
	Word    []string
	Payload []interface{}
	Length  int
}

// MakePrefixTree is for constructing prefix tree for word with payload list
func MakePrefixTree(wordsWithPayload WordsWithPayload) *PrefixTree {
	sort.Sort(byWord(wordsWithPayload))
	tab := make(map[PrefixTreeNode]*PrefixTreePointer)

	for i := 0; i < wordsWithPayload.Length; i++ {
		word := wordsWithPayload.Word[i]
		payload := wordsWithPayload.Payload[i]
		rowNo := 0

		runes := []rune(word)
		for j, ch := range runes {
			isFinal := ((j + 1) == len(runes))
			node := PrefixTreeNode{rowNo, j, ch}
			child, found := tab[node]

			if !found {
				var thisPayload interface{}
				if isFinal {
					thisPayload = payload
				} else {
					thisPayload = nil
				}
				tab[node] = &PrefixTreePointer{i, isFinal, thisPayload}
				rowNo = i
			} else {
				rowNo = child.ChildID
			}
		}
	}
	return &PrefixTree{tab}
}

// Lookup - look up prefix tree from node-id, offset and a character
func (tree *PrefixTree) Lookup(nodeID int, offset int, ch rune) (*PrefixTreePointer, bool) {
	pointer, found := tree.tab[PrefixTreeNode{nodeID, offset, ch}]
	return pointer, found
}
