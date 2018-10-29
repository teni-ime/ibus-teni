package main

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	topPage            = "https://vi.wiktionary.org"
	firstPage          = "https://vi.wiktionary.org/wiki/Thể_loại:Mục_từ_tiếng_Việt"
	contentFromPattern = `[\(>]Trang sau[\)<]`
	nextPagePattern    = `<a href="([^"]+)" title="Thể loại:Mục từ tiếng Việt">Trang sau</a>`
	pageTitlePattern   = `<li><a href="/wiki/[^"]+" title="[^"]+">([^>]+)</a></li>`
	outputFile         = "dict/vi.wiktionary.org.txt"

	vietChars = `
 áàảãạ
ăắằẳẵặ
âấầẩẫậ
 éèẻẽẹ
êếềểễệ
 íìỉĩị
 óòỏõọ
ôốồổỗộ
ơờớởỡợ
 úùủũụ
ưứừửữự
 ýỳỷỹỵ
đ
`
)

var (
	nextPageRegex      = regexp.MustCompile(nextPagePattern)
	nextPageTitleRegex = regexp.MustCompile(pageTitlePattern)
	contentStartRegex  = regexp.MustCompile(contentFromPattern)
)

func GetHttpText(address string) string {
	uAddr, _ := url.PathUnescape(address)
	log.Println("Get " + uAddr)
	if resp, err := http.Get(address); err != nil {
		log.Println(err)
	} else {
		defer resp.Body.Close()
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			log.Println(err)
		} else {
			return string(b)
		}
	}

	return ""
}

func isVietnamese(s string) bool {
	for _, r := range []rune(s) {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			(r == ' ' || r == '.' || r == ',' || r == ';' || r == '&') ||
			strings.IndexRune(vietChars, r) >= 0 {
			continue
		} else {
			return false
		}
	}

	return true
}

func dumpWiktionary() string {
	mapWord := map[string]bool{}

	s := GetHttpText(firstPage)
	for len(s) > 0 {
		s = html.UnescapeString(s)
		startIndex := contentStartRegex.FindStringIndex(s)
		if len(startIndex) > 0 {
			content := s[startIndex[0]:]
			matches := nextPageTitleRegex.FindAllStringSubmatch(content, -1)
			for _, m := range matches {
				w := m[1]
				if isVietnamese(w) {
					mapWord[w] = true
				}
			}

		}

		nextALink := nextPageRegex.FindAllStringSubmatch(s, 1)
		if len(nextALink) > 0 && len(nextALink[0]) >= 1 {
			nextUrl := nextALink[0][1]
			s = GetHttpText(topPage + nextUrl)
		} else {
			break
		}
	}

	var wordList []string
	for k := range mapWord {
		wordList = append(wordList, k)
	}

	return strings.Join(wordList, "\n")
}
