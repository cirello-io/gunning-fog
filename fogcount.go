package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)
	words := strings.Fields(strings.Replace(s, "\n", " ", -1))
	stopsSymbols := []string{".", ",", ";"}
	var phraseSize = 0
	var phraseSizes []int
	var hardWords int

	for _, w := range words {
		phraseSize++
		for _, stop := range stopsSymbols {
			if strings.HasSuffix(w, stop) {
				w = strings.TrimSuffix(w, stop)
				phraseSizes = append(phraseSizes, phraseSize)
				phraseSize = 0
			}
		}
		if len(w) > 6 {
			hardWords++
		}
	}

	var totalSentenceLength int
	for _, size := range phraseSizes {
		totalSentenceLength = totalSentenceLength + size
	}

	var gradeLevel float64 = 0.4 * ((float64(totalSentenceLength) / float64(len(phraseSizes))) + (100 * float64(hardWords) / float64(len(words))))
	fmt.Println(gradeLevel)
}
