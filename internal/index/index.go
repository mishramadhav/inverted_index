package index

import (
	"slices"
	"strings"
	"sync"

	"github.com/mishramadhav/inverted_index/internal/operations"
	"github.com/mishramadhav/inverted_index/internal/set"
)

type wordLocation struct {
	docId int
	pos   int
}

type invertedIndex struct {
	mu   sync.RWMutex
	data map[string][]*wordLocation
}

func NewInvertedIndex() *invertedIndex {
	return &invertedIndex{
		data: make(map[string][]*wordLocation),
	}
}

// GetDocumentIDs returns the list of document IDs that contain the given key.
func (idx *invertedIndex) GetDocumentIDs(word string) []int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if _, ok := idx.data[word]; !ok {
		return nil
	}

	docIds := operations.Map(idx.data[word], func(wordLoc *wordLocation) int {
		return wordLoc.docId
	})

	uniqueDocIDs := set.NewWithValues(docIds...).Values()
	slices.Sort(uniqueDocIDs)
	return uniqueDocIDs
}

// AddDocument adds a document to the index.
func (idx *invertedIndex) AddDocument(docId int, content string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()

	// Split the content into words.
	words := strings.Fields(strings.ToLower(content))
	for i, word := range words {
		// Append the document ID to the list of documents for this word.
		idx.data[word] = append(idx.data[word], &wordLocation{docId: docId, pos: i})
	}
}

// GetWordLocations returns the list of word locations for the given word.
func (idx *invertedIndex) GetWordLocations(word string) [][]int {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if _, ok := idx.data[word]; !ok {
		return nil
	}

	return operations.Map(idx.data[word], func(wordLoc *wordLocation) []int {
		return []int{wordLoc.docId, wordLoc.pos}
	})
}
