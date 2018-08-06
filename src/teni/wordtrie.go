/*
 * Teni-IME - A Vietnamese Input method editor
 * Copyright (C) 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com>
 * This file is part of Teni-IME.
 *
 *  Teni-IME is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Teni-IME is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Teni-IME.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package teni

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"unicode"
)

const (
	FindResultNotMatch = iota
	FindResultMatchPrefix
	FindResultMatchFull
)

//Word trie
type W struct {
	F bool        //Full word
	N map[rune]*W // Next characters
}

var rootWordTrie = &W{F: false}

func addTrie(trie *W, s []rune, down bool) {

	if trie.N == nil {
		trie.N = map[rune]*W{}
	}

	//add original char
	s0 := s[0]
	if trie.N[s0] == nil {
		trie.N[s0] = &W{}
	}

	if len(s) == 1 {
		if !trie.N[s0].F {
			trie.N[s0].F = !down
		}
	} else {
		addTrie(trie.N[s0], s[1:], down)
	}

	//add down 1 level char
	if dmap, exist := downLvlMap[s0]; exist {
		for _, r := range dmap {
			if trie.N[r] == nil {
				trie.N[r] = &W{}
			}

			if len(s) == 1 {
				trie.N[r].F = true
			} else {
				addTrie(trie.N[r], s[1:], true)
			}
		}
	}
}

func findWord(t *W, s []rune) (result uint8) {

	if len(s) == 0 {
		if t.F {
			return FindResultMatchFull
		}
		return FindResultMatchPrefix
	}

	c := unicode.ToLower(s[0])

	if t.N[c] != nil {
		r := findWord(t.N[c], s[1:])
		if r != FindResultNotMatch {
			return r
		}
	}

	return FindResultNotMatch
}

func findRootWord(s []rune) (result uint8) {
	return findWord(rootWordTrie, s)
}

func fileExist(p string) bool {
	sta, err := os.Stat(p)
	return err == nil && !sta.IsDir()
}

func InitWordTrie(dataFiles ...string) error {
	rootWordTrie = &W{F: false}

	for _, dataFile := range dataFiles {
		if !fileExist(dataFile) && !filepath.IsAbs(dataFile) {
			dataFile = filepath.Join(filepath.Dir(os.Args[0]), dataFile)
		}
		f, err := os.Open(dataFile)
		if err != nil {
			return err
		}
		rd := bufio.NewReader(f)
		for {
			line, _, _ := rd.ReadLine()
			if len(line) == 0 {
				break
			}
			addTrie(rootWordTrie, []rune(string(line)), false)
		}
		f.Close()
	}
	runtime.GC()
	return nil
}
