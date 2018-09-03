package main

import (
	"bufio"
	"compress/bzip2"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	dumpUrl = "https://dumps.wikimedia.org/viwiktionary/latest/viwiktionary-latest-pages-articles.xml.bz2"
)

func openDumpReader() (*bufio.Reader, error) {
	resp, err := http.Get(dumpUrl)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(bzip2.NewReader(resp.Body)), nil
}

func dumpVieText() (chan string, error) {

	rd, err := openDumpReader()
	if err != nil {
		return nil, err
	}

	ch := make(chan string)

	go func() {
		entered := false
		var lines []string
		for {
			line, _, err := rd.ReadLine()
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}
				break
			}

			if len(line) == 0 {
				continue
			}

			sline := strings.TrimSpace(string(line))
			if entered {
				lines = append(lines, sline)
				if strings.HasSuffix(sline, "</text>") {
					ch <- strings.Join(lines, "\n")
					lines = nil
					entered = false
				}
			} else if strings.HasPrefix(sline, "<text") &&
				strings.HasSuffix(sline, "{{-vie-}}") {
				entered = true
			}
		}
		close(ch)
	}()

	return ch, nil
}
