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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

const (
	dataDir    = "dict"
	dataSuffix = ".dict"

	outputDir    = "test-data"
	outputSuffix = ".tdata"
)

var genFuncMap = map[string]func(string) string{
	"telex1": telex1,
	"telex2": telex2,
	"telex3": telex3,
	"telexw": telexw,
	"vni1":   vni1,
	"vni2":   vni2,
	"vni3":   vni3,
}

//fatal error
func fe(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fos, err := ioutil.ReadDir(dataDir)
	fe(err)
	for _, fo := range fos {
		if fo.IsDir() {
			continue
		}
		if strings.HasSuffix(fo.Name(), dataSuffix) {
			fName := fo.Name()
			dictFile := filepath.Join(dataDir, fName)
			lines := readFileLines(dictFile)
			genTestData(fName, lines)
		}
	}
}

func readFileLines(f string) []string {
	data, e := ioutil.ReadFile(f)
	fe(e)
	s := strings.Replace(string(data), "\r", "", -1)
	return strings.Split(s, "\n")
}

func genTestData(inputFileName string, lines []string) {
	for genFuncName, genFunc := range genFuncMap {
		outputFileName := filepath.Join(outputDir, inputFileName+"."+genFuncName+outputSuffix)
		fo, e := os.OpenFile(outputFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		fe(e)
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			inputStr := genFunc(line)
			if len(inputStr) > 0 {
				fmt.Fprintln(fo, line, inputStr)
			}
		}
	}

}

func telex1(line string) string {

	var or []rune
	for _, r := range []rune(line) {
		if bt, has := baseToneTelex[r]; has {
			if bt2, has2 := baseToneTelex[bt.B]; has2 {
				or = append(or, bt2.B, bt2.T, bt.T)
			} else {
				or = append(or, bt.B, bt.T)
			}
		} else {
			or = append(or, r)
		}
	}

	return string(or)
}

func telex2(line string) string {
	var base []rune
	var tone []rune
	for _, r := range []rune(line) {
		if bt, has := baseToneTelex[r]; has {
			if bt2, has2 := baseToneTelex[bt.B]; has2 {
				base = append(base, bt2.B)
				tone = append(tone, bt2.T, bt.T)
			} else {
				base = append(base, bt.B)
				tone = append(tone, bt.T)
			}
		} else {
			base = append(base, r)
		}
	}

	baseStr := string(base)
	baseStr = strings.Replace(baseStr, "oo", "ooo", -1)

	toneStr := string(tone)
	toneStr = strings.Replace(toneStr, "ww", "w", -1)
	return baseStr + toneStr
}

func telex3(line string) string {
	var base []rune
	var tone []rune
	for _, r := range []rune(line) {
		if bt, has := baseToneTelex[r]; has {
			if bt2, has2 := baseToneTelex[bt.B]; has2 {
				base = append(base, bt2.B, bt2.T)
				tone = append(tone, bt.T)
			} else {
				base = append(base, bt.B)
				tone = append(tone, bt.T)
			}
		} else {
			base = append(base, r)
		}
	}

	baseStr := string(base)

	toneStr := string(tone)

	if strings.Contains(baseStr, "w") && strings.Contains(toneStr, "w") {
		toneStr = strings.Replace(toneStr, "w", "", 1)
	} else {
		toneStr = strings.Replace(toneStr, "ww", "w", 1)
	}

	return baseStr + toneStr
}

func telexw(line string) string {
	var or []rune
	for _, r := range []rune(line) {
		if bt, has := baseToneTelex[r]; has {
			if bt2, has2 := baseToneTelex[bt.B]; has2 {
				or = append(or, bt2.B, bt2.T, bt.T)
			} else {
				or = append(or, bt.B, bt.T)
			}
		} else {
			or = append(or, r)
		}
	}

	s := string(or)
	if strings.Contains(s, "ow") {
		return strings.Replace(s, "ow", "[", -1)
	} else if strings.Contains(s, "uw") {
		if rand.Int()%2 == 0 {
			return strings.Replace(s, "uw", "]", -1)
		} else {
			return strings.Replace(s, "uw", "w", -1)
		}
	}

	return ""
}

func vni1(line string) string {

	var or []rune
	for _, r := range []rune(line) {
		if bt, has := baseToneVni[r]; has {
			if bt2, has2 := baseToneVni[bt.B]; has2 {
				or = append(or, bt2.B, bt2.T, bt.T)
			} else {
				or = append(or, bt.B, bt.T)
			}
		} else {
			or = append(or, r)
		}
	}

	return string(or)
}

func vni2(line string) string {
	var base []rune
	var tone []rune
	for _, r := range []rune(line) {
		if bt, has := baseToneVni[r]; has {
			if bt2, has2 := baseToneVni[bt.B]; has2 {
				base = append(base, bt2.B)
				tone = append(tone, bt2.T, bt.T)
			} else {
				base = append(base, bt.B)
				tone = append(tone, bt.T)
			}
		} else {
			base = append(base, r)
		}
	}

	baseStr := string(base)
	toneStr := string(tone)
	toneStr = strings.Replace(toneStr, "77", "7", -1)
	return baseStr + toneStr
}

func vni3(line string) string {
	var base []rune
	var tone []rune
	for _, r := range []rune(line) {
		if bt, has := baseToneVni[r]; has {
			if bt2, has2 := baseToneVni[bt.B]; has2 {
				base = append(base, bt2.B, bt2.T)
				tone = append(tone, bt.T)
			} else {
				base = append(base, bt.B)
				tone = append(tone, bt.T)
			}
		} else {
			base = append(base, r)
		}
	}

	baseStr := string(base)
	toneStr := string(tone)
	if strings.Contains(baseStr, "7") && strings.Contains(toneStr, "7") {
		toneStr = strings.Replace(toneStr, "7", "", 1)
	}
	toneStr = strings.Replace(toneStr, "77", "7", -1)
	return baseStr + toneStr
}

type BT struct {
	T rune //Tone char
	B rune //Base char
}

var baseToneTelex = map[rune]*BT{
	'ă': {'w', 'a'},
	'â': {'a', 'a'},
	'ê': {'e', 'e'},
	'ô': {'o', 'o'},
	'ơ': {'w', 'o'},
	'ư': {'w', 'u'},
	'đ': {'d', 'd'},

	'á': {'s', 'a'},
	'ắ': {'s', 'ă'},
	'ấ': {'s', 'â'},
	'é': {'s', 'e'},
	'ế': {'s', 'ê'},
	'í': {'s', 'i'},
	'ó': {'s', 'o'},
	'ố': {'s', 'ô'},
	'ớ': {'s', 'ơ'},
	'ú': {'s', 'u'},
	'ứ': {'s', 'ư'},
	'ý': {'s', 'y'},

	'à': {'f', 'a'},
	'ằ': {'f', 'ă'},
	'ầ': {'f', 'â'},
	'è': {'f', 'e'},
	'ề': {'f', 'ê'},
	'ì': {'f', 'i'},
	'ò': {'f', 'o'},
	'ồ': {'f', 'ô'},
	'ờ': {'f', 'ơ'},
	'ù': {'f', 'u'},
	'ừ': {'f', 'ư'},
	'ỳ': {'f', 'y'},

	'ả': {'r', 'a'},
	'ẳ': {'r', 'ă'},
	'ẩ': {'r', 'â'},
	'ẻ': {'r', 'e'},
	'ể': {'r', 'ê'},
	'ỉ': {'r', 'i'},
	'ỏ': {'r', 'o'},
	'ổ': {'r', 'ô'},
	'ở': {'r', 'ơ'},
	'ủ': {'r', 'u'},
	'ử': {'r', 'ư'},
	'ỷ': {'r', 'y'},

	'ã': {'x', 'a'},
	'ẵ': {'x', 'ă'},
	'ẫ': {'x', 'â'},
	'ẽ': {'x', 'e'},
	'ễ': {'x', 'ê'},
	'ĩ': {'x', 'i'},
	'õ': {'x', 'o'},
	'ỗ': {'x', 'ô'},
	'ỡ': {'x', 'ơ'},
	'ũ': {'x', 'u'},
	'ữ': {'x', 'ư'},
	'ỹ': {'x', 'y'},

	'ạ': {'j', 'a'},
	'ặ': {'j', 'ă'},
	'ậ': {'j', 'â'},
	'ẹ': {'j', 'e'},
	'ệ': {'j', 'ê'},
	'ị': {'j', 'i'},
	'ọ': {'j', 'o'},
	'ộ': {'j', 'ô'},
	'ợ': {'j', 'ơ'},
	'ụ': {'j', 'u'},
	'ự': {'j', 'ư'},
	'ỵ': {'j', 'y'},
}

var baseToneVni = map[rune]*BT{
	'â': {'6', 'a'},
	'ê': {'6', 'e'},
	'ô': {'6', 'o'},
	'ơ': {'7', 'o'},
	'ư': {'7', 'u'},
	'ă': {'8', 'a'},
	'đ': {'9', 'd'},

	'á': {'1', 'a'},
	'ắ': {'1', 'ă'},
	'ấ': {'1', 'â'},
	'é': {'1', 'e'},
	'ế': {'1', 'ê'},
	'í': {'1', 'i'},
	'ó': {'1', 'o'},
	'ố': {'1', 'ô'},
	'ớ': {'1', 'ơ'},
	'ú': {'1', 'u'},
	'ứ': {'1', 'ư'},
	'ý': {'1', 'y'},

	'à': {'2', 'a'},
	'ằ': {'2', 'ă'},
	'ầ': {'2', 'â'},
	'è': {'2', 'e'},
	'ề': {'2', 'ê'},
	'ì': {'2', 'i'},
	'ò': {'2', 'o'},
	'ồ': {'2', 'ô'},
	'ờ': {'2', 'ơ'},
	'ù': {'2', 'u'},
	'ừ': {'2', 'ư'},
	'ỳ': {'2', 'y'},

	'ả': {'3', 'a'},
	'ẳ': {'3', 'ă'},
	'ẩ': {'3', 'â'},
	'ẻ': {'3', 'e'},
	'ể': {'3', 'ê'},
	'ỉ': {'3', 'i'},
	'ỏ': {'3', 'o'},
	'ổ': {'3', 'ô'},
	'ở': {'3', 'ơ'},
	'ủ': {'3', 'u'},
	'ử': {'3', 'ư'},
	'ỷ': {'3', 'y'},

	'ã': {'4', 'a'},
	'ẵ': {'4', 'ă'},
	'ẫ': {'4', 'â'},
	'ẽ': {'4', 'e'},
	'ễ': {'4', 'ê'},
	'ĩ': {'4', 'i'},
	'õ': {'4', 'o'},
	'ỗ': {'4', 'ô'},
	'ỡ': {'4', 'ơ'},
	'ũ': {'4', 'u'},
	'ữ': {'4', 'ư'},
	'ỹ': {'4', 'y'},

	'ạ': {'5', 'a'},
	'ặ': {'5', 'ă'},
	'ậ': {'5', 'â'},
	'ẹ': {'5', 'e'},
	'ệ': {'5', 'ê'},
	'ị': {'5', 'i'},
	'ọ': {'5', 'o'},
	'ộ': {'5', 'ô'},
	'ợ': {'5', 'ơ'},
	'ụ': {'5', 'u'},
	'ự': {'5', 'ư'},
	'ỵ': {'5', 'y'},
}
