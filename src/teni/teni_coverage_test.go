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
	"testing"
)

func TestTeniCoverage(t *testing.T) {
	cdDictDir()

	e := InitWordTrie(true, "you can't find me")
	if e == nil {
		t.Errorf("e must be not nil here!")
	}

	e = InitWordTrie(true, DictStdList...)
	fe(e)
	pc := NewEngine()
	initResultLen := pc.ResultLen()
	if initResultLen != 0 {
		t.Errorf("initResultLen %d", initResultLen)
	}
	initResult := pc.GetResult()
	if initResult != nil {
		t.Errorf("initResult %+v", initResult)
	}
	initCommitResult := pc.GetResultStr()
	if initCommitResult != "" {
		t.Errorf("initResult %+v", initResult)
	}

	keys := []rune{'t', 'i', 'e', 'e', 's', 'n', 'g'}
	for _, k := range keys {
		pc.AddKey(k)
	}

	resultLen := pc.ResultLen()
	expectedResultLen := uint32(len([]rune("tiếng")))
	if resultLen != expectedResultLen {
		t.Errorf("ResultLen %d, expected %d", resultLen, expectedResultLen)
	}

	rawKeyLen := pc.RawKeyLen()
	expectedRawKeyLen := len(keys)
	if rawKeyLen != expectedRawKeyLen {
		t.Errorf("rawKeyLen %d, expected %d", rawKeyLen, expectedRawKeyLen)
	}

	pc.AddKey('a')
	commitStr := pc.GetResultStr()
	expectedCommitStr := "tieesnga"
	if commitStr != expectedCommitStr {
		t.Errorf("commitStr [%s], expectedCommitStr [%s]", commitStr, expectedCommitStr)
	}

	pc.Backspace()
	commitStr = pc.GetResultStr()
	expectedCommitStr = "tiếng"
	if commitStr != expectedCommitStr {
		t.Errorf("commitStr [%s], expectedCommitStr [%s]", commitStr, expectedCommitStr)
	}

	rawStr := pc.GetRawStr()
	expectedRawStr := "tieesng"
	if rawStr != expectedRawStr {
		t.Errorf("commitStr [%s], expectedCommitStr [%s]", rawStr, expectedRawStr)
	}

	pc.AddStr("aaaaaaaaaaaaaaaaaa")
	commitStr = pc.GetResultStr()
	expectedCommitStr = "tieesngaaaaaaaaaaaaaaaaaa"
	if commitStr != expectedCommitStr {
		t.Errorf("commitStr [%s], expectedCommitStr [%s]", commitStr, expectedCommitStr)
	}
}
