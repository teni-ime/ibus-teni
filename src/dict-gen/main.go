package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	mapStdToneWord    = loadWordMap("dict/vietnamese.std.dict")
	mapNewToneWord    = loadWordMap("dict/vietnamese.new.dict")
	mapSpecialWord    = loadWordMap("dict/vietnamese.sp.dict")
	mapCommonWord     = loadWordMap("dict/vietnamese.cm.dict")
	newCommonWordFile = "dict/vietnamese.cm.dict"
)

func loadWordMap(wordListFile string) map[string]bool {
	f, err := os.Open(wordListFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	m := map[string]bool{}
	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		if len(line) == 0 {
			continue
		}
		m[string(line)] = true
	}

	return m
}

func main() {
	fmt.Println("BEGIN")
	log.Println(len(mapCommonWord))
	log.Println(len(mapStdToneWord))
	log.Println(len(mapNewToneWord))

	allWords := dumpWiktionary()

	m := map[string]bool{}
	extractVietWord(allWords, m)

	countNewWord := 0
	for k := range m {
		if _, ok := mapStdToneWord[k]; ok {
			continue
		} else if _, ok := mapNewToneWord[k]; ok {
			continue
		} else if _, ok := mapSpecialWord[k]; ok {
			continue
		} else if _, ok := mapCommonWord[k]; ok {
			continue
		} else {
			mapCommonWord[k] = true
			countNewWord++
		}
	}

	log.Println("countNewWord:", countNewWord)
	var words []string
	for k := range mapCommonWord {
		if len(k) > 0 {
			words = append(words, k)
		}
	}

	vnsort(words)
	ioutil.WriteFile(newCommonWordFile, []byte(strings.Join(words, "\n")), 0777)

	fmt.Println("DONE")
}
