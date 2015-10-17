package main

import (
	"fmt"
	"github.com/evalvarez12/language_processing/processing"
)

func main() {

	words := processing.ProcessFile("unlabeledTrainData.tsv")

	for i := range words {
		fmt.Println(words[i])
	}
}
