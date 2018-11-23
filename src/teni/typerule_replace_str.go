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
}

var replaceStrMap = map[rune]map[string]*SR{
	's': {
		"uơ": {"ướ"},
		"ưo": {"ướ"},
	},
	'f': {
		"uơ": {"ườ"},
		"ưo": {"ườ"},
	},
	'r': {
		"uơ": {"ưở"},
		"ưo": {"ưở"},
	},
	'x': {
		"uơ": {"ưỡ"},
		"ưo": {"ưỡ"},
	},
	'j': {
		"uơ": {"ượ"},
		"ưo": {"ượ"},
	},
	'w': {
		"uo": {"ươ"},
		"uó": {"ướ"},
		"uò": {"ườ"},
		"uỏ": {"ưở"},
		"uõ": {"ưỡ"},
		"uọ": {"ượ"},

		"úo": {"ướ"},
		"ùo": {"ườ"},
		"ủo": {"ưở"},
		"ũo": {"ưỡ"},
		"ụo": {"ượ"},

		"oa": {"oă"},
		"oá": {"oắ"},
		"oà": {"oằ"},
		"oả": {"oẳ"},
		"oã": {"oẵ"},
		"oạ": {"oặ"},
		"óa": {"oắ"},
		"òa": {"oằ"},
		"ỏa": {"oẳ"},
		"õa": {"oẵ"},
		"ọa": {"oặ"},
	},
}
