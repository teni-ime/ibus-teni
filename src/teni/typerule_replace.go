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
	"strings"
	"unicode"
)

func joinMapCR(maps ...map[rune]*CR) map[rune]*CR {
	newMap := map[rune]*CR{}
	for _, m := range maps {
		for k, v := range m {
			newMap[k] = v
		}
	}

	return newMap
}

func upFirstChar(s string) string {
	var Str string
	sRune := []rune(s)
	if len(sRune) >= 1 {
		Str = string(unicode.ToUpper(sRune[0])) + string(sRune[1:])
	}
	return Str
}

func init() {
	addVniRule()

	//BEGIN Add UP-CASE replaceCharMap
	for _, m := range replaceCharMap {
		var keys []rune
		for k := range m {
			keys = append(keys, k)
		}
		for _, k := range keys {
			m[unicode.ToUpper(k)] = &CR{
				C: unicode.ToUpper(m[k].C),
			}
		}
	}
	var lowerKeys []rune
	for k := range replaceCharMap {
		if unicode.IsLower(k) {
			lowerKeys = append(lowerKeys, k)
		}
	}
	for _, k := range lowerKeys {
		replaceCharMap[unicode.ToUpper(k)] = replaceCharMap[k]
	}
	//END Add UP-CASE replaceCharMap

	//BEGIN Add UP-CASE replaceStrMap
	for _, m := range replaceStrMap {
		var keys []string
		for s := range m {
			keys = append(keys, s)
		}
		for _, s := range keys {
			m[strings.ToUpper(s)] = &SR{
				S: strings.ToUpper(m[s].S),
			}
			m[upFirstChar(s)] = &SR{
				S: upFirstChar(m[s].S),
			}
		}
	}
	var lowerStrKey []rune
	for k := range replaceStrMap {
		if unicode.IsLower(k) {
			lowerStrKey = append(lowerStrKey, k)
		}
	}
	for _, k := range lowerStrKey {
		replaceStrMap[unicode.ToUpper(k)] = replaceStrMap[k]
	}
	//END Add UP-CASE replaceStrMap

}
