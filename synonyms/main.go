package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sky0621/study-gowebchap4/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	fmt.Println(apiKey)

	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました： %v\n", word, err)
			return
		}
		if len(syns) == 0 {
			log.Fatalf("%qに類語はありませんでした\n", word)
			return
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
