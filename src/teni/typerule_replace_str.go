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

type SR struct {
	S string //change to S
	R bool   //is revertMode
}

var replaceStrMap = map[rune]map[string]*SR{
	's': {
		"uơ": {"ướ", false},
		"ưo": {"ướ", false},
	},
	'f': {
		"uơ": {"ườ", false},
		"ưo": {"ườ", false},
	},
	'r': {
		"uơ": {"ưở", false},
		"ưo": {"ưở", false},
	},
	'x': {
		"uơ": {"ưỡ", false},
		"ưo": {"ưỡ", false},
	},
	'j': {
		"uơ": {"ượ", false},
		"ưo": {"ượ", false},
	},
	'w': {
		"uo": {"ươ", false},
		"uó": {"ướ", false},
		"uò": {"ườ", false},
		"uỏ": {"ưở", false},
		"uõ": {"ưỡ", false},
		"uọ": {"ượ", false},

		"úo": {"ướ", false},
		"ùo": {"ườ", false},
		"ủo": {"ưở", false},
		"ũo": {"ưỡ", false},
		"ụo": {"ượ", false},

		"oa": {"oă", false},
		"oá": {"oắ", false},
		"oà": {"oằ", false},
		"oả": {"oẳ", false},
		"oã": {"oẵ", false},
		"oạ": {"oặ", false},
		"óa": {"oắ", false},
		"òa": {"oằ", false},
		"ỏa": {"oẳ", false},
		"õa": {"oẵ", false},
		"ọa": {"oặ", false},

		"ươ": {"uo", true},
		"ướ": {"uó", true},
		"ườ": {"uò", true},
		"ưở": {"uỏ", true},
		"ưỡ": {"uõ", true},
		"ượ": {"uọ", true},
	},
}
