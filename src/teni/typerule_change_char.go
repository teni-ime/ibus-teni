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

var changeCharMap = map[rune]rune{
	'w': 'ư',
	'W': 'Ư',
}

var changeCharMapEx = map[rune]rune{
	'[': 'ơ',
	'{': 'Ơ',
	']': 'ư',
	'}': 'Ư',
	'w': 'ư',
	'W': 'Ư',
}

func InChangeCharMap(c rune) bool {
	_, exist := changeCharMap[c]
	return exist
}

func InChangeCharMapEx(c rune) bool {
	_, exist := changeCharMapEx[c]
	return exist
}

var caplockSwitchMap = map[uint32]uint32{
	'[': '{',
	'{': '[',
	']': '}',
	'}': ']',
}

func SwitchCaplock(keyVal uint32) uint32 {
	if v, exist := caplockSwitchMap[keyVal]; exist {
		return v
	}
	return keyVal
}
