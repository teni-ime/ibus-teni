package main

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func vnsort(a []string) {
	vnm := collate.New(language.Vietnamese)
	vnm.SortStrings(a)
}
