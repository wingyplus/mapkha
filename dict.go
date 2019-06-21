package mapkha

import (
	"bufio"
	"os"
	"path"
	"runtime"
)

// Dict is a prefix tree
type Dict struct {
	tree *PrefixTree
}

// LoadDict is for loading a word list from file
func LoadDict(path string) (*Dict, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	wordsWithPayload := WordsWithPayload{
		Word:    []string{},
		Payload: []interface{}{},
		Length:  0,
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) != 0 {
			wordsWithPayload.Word = append(wordsWithPayload.Word, line)
			wordsWithPayload.Payload = append(wordsWithPayload.Payload, true)
			wordsWithPayload.Length++
		}
	}
	tree := MakePrefixTree(wordsWithPayload)
	dix := Dict{tree}
	return &dix, nil
}

func MakeDict(words []string) *Dict {
	wordsWithPayload := WordsWithPayload{
		Word:    []string{},
		Payload: []interface{}{},
		Length:  0,
	}

	for _, word := range words {
		wordsWithPayload.Word = append(wordsWithPayload.Word, word)
		wordsWithPayload.Payload = append(wordsWithPayload.Payload, true)
		wordsWithPayload.Length++
	}
	tree := MakePrefixTree(wordsWithPayload)
	dix := Dict{tree}
	return &dix
}

// LoadDefaultDict - loading default Thai dictionary
func LoadDefaultDict() (*Dict, error) {
	_, filename, _, _ := runtime.Caller(0)
	return LoadDict(path.Join(path.Dir(filename), "tdict-std.txt"))
}

// Lookup - lookup node in a Prefix Tree
func (d *Dict) Lookup(p int, offset int, ch rune) (*PrefixTreePointer, bool) {
	pointer, found := d.tree.Lookup(p, offset, ch)
	return pointer, found
}
