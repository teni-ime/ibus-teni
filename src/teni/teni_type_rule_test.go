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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	testDataDir           = "test-data"
	testDataSuffix        = ".tdata"
	vniTestDataFileSign   = ".vni"
	telexTestDataFileSign = ".telexw"
	newTestDataFileSign   = ".new"
)

const (
	DictDir           = "dict"
	DictVietnameseCm  = "dict/vietnamese.cm.dict"
	DictVietnameseSp  = "dict/vietnamese.sp.dict"
	DictVietnameseStd = "dict/vietnamese.std.dict"
	DictVietnameseNew = "dict/vietnamese.new.dict"
)

var (
	DictStdList = []string{DictVietnameseCm, DictVietnameseSp, DictVietnameseStd}
	DictNewList = []string{DictVietnameseCm, DictVietnameseSp, DictVietnameseNew}
)

//fatal error
func fe(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cdDictDir() {
	if st, e := os.Stat(DictDir); e == nil && st.IsDir() {
		return
	}

	gp := filepath.Dir(os.Args[0])
	p := filepath.Join(gp, DictDir)
	if st, e := os.Stat(p); e == nil && st.IsDir() {
		os.Chdir(p)
		return
	}

	goPath := os.Getenv("GOPATH")
	for _, gp := range strings.Split(goPath, ":") {
		p := filepath.Join(gp, DictDir)
		if st, e := os.Stat(p); e == nil && st.IsDir() {
			os.Chdir(gp)
			return
		}
	}
}

func TestTeniTypeRule(t *testing.T) {
	cdDictDir()

	failedCount := 0
	testCaseCount := 0

	pc := NewEngine()

	fos, err := ioutil.ReadDir(testDataDir)
	fe(err)
	for _, fo := range fos {
		if fo.IsDir() || !strings.HasSuffix(fo.Name(), testDataSuffix) {
			continue
		}
		if strings.Contains(fo.Name(), newTestDataFileSign) {
			e := InitWordTrie(true, DictNewList...)
			fe(e)
		} else {
			e := InitWordTrie(true, DictStdList...)
			fe(e)
		}

		t.Log("Testing ", fo.Name())
		if strings.Contains(fo.Name(), vniTestDataFileSign) {
			pc.InputMethod = IMVni
		} else if strings.Contains(fo.Name(), telexTestDataFileSign) {
			pc.InputMethod = IMTelexEx
		} else {
			pc.InputMethod = IMTeni
		}

		for iLine, line := range readFileLines(filepath.Join(testDataDir, fo.Name())) {
			inout := strings.Split(line, " ")
			if len(inout) != 2 {
				continue
			}
			out := inout[0]
			in := inout[1]

			testCaseCount++
			pc.Reset()
			pc.AddStr(in)
			result := pc.GetResultStr()
			if result != out {
				t.Errorf("\tLine #%d for [%s], expected [%s], got [%s]", iLine, in, out, result)
				failedCount++
			}
		}
	}

	if failedCount > 0 {
		t.Log("Failed count", failedCount, "of", testCaseCount, "total")
	} else {
		t.Log("Passed", testCaseCount, "test cases")
	}

}

func readFileLines(f string) []string {
	data, e := ioutil.ReadFile(f)
	fe(e)
	s := strings.Replace(string(data), "\r", "", -1)
	return strings.Split(s, "\n")
}
