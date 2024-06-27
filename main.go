package main

import (
	"fmt"

	"github.com/mishramadhav/inverted_index/internal/index"
)

func main() {
	myIndex := index.NewInvertedIndex()

	myIndex.AddDocument(1, "The quick brown fox")
	myIndex.AddDocument(2, "The quick brown dog")
	myIndex.AddDocument(3, "The quick brown cat")
	myIndex.AddDocument(4, "jumped over the lazy dog")
	myIndex.AddDocument(5, "jumped over the lazy cat")
	myIndex.AddDocument(6, "jumped over the lazy fox")

	fmt.Println("quick found in ", myIndex.GetDocumentIDs("quick"))
	fmt.Println("quick found in docs at location ", myIndex.GetWordLocations("quick"))
	fmt.Println("dog found in ", myIndex.GetDocumentIDs("dog"))
	fmt.Println("dog found in docs at location ", myIndex.GetWordLocations("dog"))
	fmt.Println("jumped found in ", myIndex.GetDocumentIDs("jumped"))
	fmt.Println("jumped found in docs at location ", myIndex.GetWordLocations("jumped"))
}
