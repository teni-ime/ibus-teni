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
	"sort"
	"strings"
)

const MaxWordLength = 15

type InputMethod int

const (
	IMTeni  InputMethod = iota << 0
	IMVni   InputMethod = iota
	IMTelex InputMethod = iota
)

type Engine struct {
	rawKeys        []rune
	resultStack    [][]rune
	completedStack []bool
	InputMethod    InputMethod
}

type resultCase struct {
	value      []rune
	findResult uint8
	revertMode bool
}

func (pc *resultCase) better(pc2 *resultCase) bool {
	return pc.findResult > pc2.findResult ||
		(pc.findResult == pc2.findResult && pc.revertMode && !pc2.revertMode)
}

type resultCases []*resultCase

func (p resultCases) Len() int { return len(p) }
func (p resultCases) Less(i, j int) bool {
	return p[i].better(p[j])
}
func (p resultCases) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func NewEngine() *Engine {
	return &Engine{
		rawKeys:        nil,
		resultStack:    nil,
		completedStack: nil,
		InputMethod:    IMTeni,
	}
}

func (pc *Engine) RawKeyLen() int {
	return len(pc.rawKeys)
}

func (pc *Engine) ResultLen() uint32 {
	if l := len(pc.resultStack); l > 0 {
		return uint32(len(pc.resultStack[l-1]))
	}

	return 0
}

func (pc *Engine) Reset() {
	if len(pc.rawKeys) > 0 {
		pc.rawKeys = pc.rawKeys[:0]
	}

	if len(pc.resultStack) > 0 {
		pc.resultStack = pc.resultStack[:0]
	}
}

func (pc *Engine) GetResult() []rune {
	if l := len(pc.resultStack); l > 0 {
		return pc.resultStack[l-1]
	}

	return nil
}

func (pc *Engine) GetResultStr() string {
	return string(pc.GetResult())
}

func (pc *Engine) HasToneChar() bool {
	for _, c := range pc.GetResult() {
		if toneCharset[c] {
			return true
		}
	}

	return false
}

func (pc *Engine) isCompleted() bool {
	if l := len(pc.completedStack); l > 0 {
		return pc.completedStack[l-1]
	}
	return false
}

func (pc *Engine) GetCommitResultStr() string {
	if pc.isCompleted() || !pc.HasToneChar() {
		return pc.GetResultStr()
	}

	return string(pc.rawKeys)
}

func (pc *Engine) GetRawStr() string {
	return string(pc.rawKeys)
}

func copyRunes(r []rune) []rune {
	t := make([]rune, len(r))
	copy(t, r)

	return t
}

func (pc *Engine) getCopyResult() []rune {
	if l := len(pc.resultStack); l > 0 {
		return copyRunes(pc.resultStack[l-1])
	}

	return nil
}

func (pc *Engine) Backspace() {
	if l := len(pc.rawKeys); l > 0 {
		pc.rawKeys = pc.rawKeys[:l-1]
	}
	if l := len(pc.resultStack); l > 0 {
		pc.resultStack = pc.resultStack[:l-1]
	}
	if l := len(pc.completedStack); l > 0 {
		pc.completedStack = pc.completedStack[:l-1]
	}
}

func (pc *Engine) AddStr(s string) {
	for _, c := range []rune(s) {
		pc.AddKey(c)
	}
}

func (pc *Engine) AddKey(key rune) {
	resultRunes := pc.getCopyResult()
	var isCompleted bool

	if len(pc.rawKeys) > MaxWordLength ||
		(pc.InputMethod == IMVni && (key < '0' || key > '9')) ||
		(pc.InputMethod == IMTelex && (key >= '0' && key <= '9')) ||
		(len(resultRunes) == 0 && (pc.InputMethod != IMTelex || !InChangeCharMap(key))) ||
		(replaceCharMap[key] == nil && replaceStrMap[key] == nil && (pc.InputMethod == IMTelex && !InChangeCharMap(key))) {
		appendCase := pc.appendChar(key, resultRunes)
		resultRunes = appendCase.value
		isCompleted = appendCase.findResult == FindResultMatchFull

		if appendCase.findResult == FindResultNotMatch {
			if pc.HasToneChar() {
				resultRunes = append(pc.rawKeys, key)
			} else {
				resultRunes = append(pc.getCopyResult(), key)
			}
		}
	} else {
		finalCase := pc.changeChar(key, resultRunes)

		if finalCase == nil || finalCase.findResult != FindResultMatchFull {
			replaceStrCase := pc.replaceStr(key, resultRunes)
			if replaceStrCase != nil &&
				(replaceStrCase.findResult != FindResultNotMatch || replaceStrCase.revertMode) &&
				(finalCase == nil || replaceStrCase.better(finalCase)) {
				finalCase = replaceStrCase
			}
		}

		if finalCase == nil || finalCase.findResult != FindResultMatchFull {
			replaceCharCase := pc.replaceChar(key, resultRunes)
			if replaceCharCase != nil &&
				(replaceCharCase.findResult != FindResultNotMatch || replaceCharCase.revertMode) &&
				(finalCase == nil || replaceCharCase.better(finalCase)) {
				finalCase = replaceCharCase
			}
		}

		if finalCase == nil || finalCase.findResult != FindResultMatchFull {
			appendCase := pc.appendChar(key, resultRunes)
			if finalCase == nil || appendCase.better(finalCase) {
				finalCase = appendCase
			}
		}

		resultRunes = finalCase.value
		isCompleted = finalCase.findResult == FindResultMatchFull

		if !finalCase.revertMode &&
			finalCase.findResult == FindResultNotMatch {
			if pc.HasToneChar() {
				resultRunes = append(pc.rawKeys, key)
			} else {
				resultRunes = append(pc.getCopyResult(), key)
			}
		}
	}

	pc.rawKeys = append(pc.rawKeys, key)
	pc.resultStack = append(pc.resultStack, resultRunes)
	pc.completedStack = append(pc.completedStack, isCompleted)
}

func (pc *Engine) appendChar(key rune, originalRunes []rune) *resultCase {
	originalRunes = append(originalRunes, key)
	if len(originalRunes) > MaxWordLength {
		return &resultCase{
			value:      originalRunes,
			findResult: FindResultNotMatch,
		}
	}

	result := findRootWord(originalRunes)
	return pc.trySwapTone(&resultCase{
		value:      originalRunes,
		findResult: result,
	})
}

func (pc *Engine) replaceStr(key rune, originalRunes []rune) *resultCase {
	rsm := replaceStrMap[key]
	if rsm == nil {
		return nil
	}

	resultText := string(originalRunes)
	for findText, replaceSR := range rsm {
		if foundIndex := strings.Index(resultText, findText); foundIndex >= 0 {
			replacedText := strings.Replace(resultText, findText, replaceSR.S, 1)

			resultRunes := []rune(replacedText)
			if replaceSR.R {
				originalRunes = append(originalRunes, key)
			}

			result := findRootWord(resultRunes)

			return &resultCase{
				value:      resultRunes,
				findResult: result,
				revertMode: replaceSR.R,
			}
		}
	}

	return nil
}

func (pc *Engine) replaceChar(key rune, originalRunes []rune) *resultCase {
	if rcm := replaceCharMap[key]; rcm != nil {
		resultCases := resultCases{}
		for i := len(originalRunes) - 1; i >= 0; i-- {
			c := originalRunes[i]
			if cReplace, found := rcm[c]; found {
				resultRunes := copyRunes(originalRunes)
				resultRunes[i] = cReplace.C
				if cReplace.R {
					resultRunes = append(resultRunes, key)
				}
				result := findRootWord(resultRunes)

				resultCases = append(resultCases, &resultCase{
					value:      resultRunes,
					findResult: result,
					revertMode: cReplace.R,
				})

				if result == FindResultMatchFull {
					break
				}
			}
		}

		if len(resultCases) > 0 {
			sort.Sort(resultCases)
			return resultCases[0]
		}
	}

	return nil
}

func (pc *Engine) changeChar(key rune, originalRunes []rune) *resultCase {
	if changeTo, exist := changeCharMap[key]; exist {
		lr := len(originalRunes)
		lrk := len(pc.rawKeys)
		//revert mode
		if lr > 0 && lrk > 0 && key != originalRunes[lr-1] && pc.rawKeys[lrk-1] == key {
			var resultRunes []rune
			if lrs := len(pc.resultStack); lrs > 1 {
				resultRunes = copyRunes(pc.resultStack[lrs-2])
			}
			resultRunes = append(resultRunes, key)
			return &resultCase{
				value:      resultRunes,
				findResult: FindResultNotMatch,
				revertMode: true,
			}
		}

		resultRunes := copyRunes(originalRunes)
		resultRunes = append(resultRunes, changeTo)

		result := findRootWord(resultRunes)

		return &resultCase{
			value:      resultRunes,
			findResult: result,
			revertMode: false,
		}
	}

	return nil
}

func (pc *Engine) trySwapTone(originalCase *resultCase) *resultCase {
	if originalCase.findResult == FindResultMatchFull {
		return originalCase
	}

	rsCopy := copyRunes(originalCase.value)
	toneKey := rune(0)

	for i, k := range rsCopy {
		if bt, exists := baseTone[k]; exists {
			rsCopy[i] = bt.B
			toneKey = bt.T
			break
		}
	}

	if toneKey == 0 {
		return originalCase
	}

	var replaceStrCase *resultCase
	replaceStrCase = pc.replaceStr(toneKey, rsCopy)

	if replaceStrCase == nil || replaceStrCase.findResult != FindResultMatchFull {
		replaceCharCase := pc.replaceChar(toneKey, rsCopy)
		if replaceCharCase != nil && replaceCharCase.findResult != FindResultNotMatch && (replaceStrCase == nil || replaceCharCase.better(replaceStrCase)) {
			replaceStrCase = replaceCharCase
		}
	}

	if replaceStrCase != nil {
		if replaceStrCase.better(originalCase) {
			return replaceStrCase
		}
	}

	return originalCase
}
