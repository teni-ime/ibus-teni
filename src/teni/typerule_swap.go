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

import "unicode"

type BT struct {
	T rune //Tone char
	B rune //Base char
}

var baseTone = map[rune]*BT{
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

	'ơ': {'w', 'o'},
	'ư': {'w', 'u'},

	//No swap these chars
	//'ă': {'w', 'a'},
	//'â': {'a', 'a'},
	//'ê': {'e', 'e'},
	//'ô': {'o', 'o'},
	//'đ': {'d', 'd'},
}

func init() {
	var keys []rune
	for k := range baseTone {
		keys = append(keys, k)
	}

	for _, k := range keys {
		b := baseTone[k]
		baseTone[unicode.ToUpper(k)] = &BT{T: b.T, B: unicode.ToUpper(b.B)}
	}
}
