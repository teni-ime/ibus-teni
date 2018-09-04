package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("BEGIN")
	ch, err := dumpVieText()
	if err != nil {
		log.Fatalln(err)
	}

	m := map[string]bool{}
	i := 0
	for {
		i++
		s, ok := <-ch
		if !ok {
			break
		}
		extractVietWord(s, m)
	}

	var words []string
	for k := range m {
		words = append(words, k)
	}

	vnsort(words)
	for _, w := range words {
		fmt.Println(w)
	}

	fmt.Println("DONE")
}
