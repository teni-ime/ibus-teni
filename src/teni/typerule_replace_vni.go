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

func addVniRule() {
	replaceCharMap['0'] = replaceCharMap['z']
	replaceCharMap['1'] = replaceCharMap['s']
	replaceCharMap['2'] = replaceCharMap['f']
	replaceCharMap['3'] = replaceCharMap['r']
	replaceCharMap['4'] = replaceCharMap['x']
	replaceCharMap['5'] = replaceCharMap['j']
	replaceCharMap['6'] = joinMapCR(replaceCharMap['a'], replaceCharMap['e'], replaceCharMap['o'])
	replaceCharMap['7'] = map[rune]*CR{
		'o': c_ơ,
		'ó': c_ớ,
		'ò': c_ờ,
		'ỏ': c_ở,
		'õ': c_ỡ,
		'ọ': c_ợ,

		'ô': c_ơ,
		'ố': c_ớ,
		'ồ': c_ờ,
		'ổ': c_ở,
		'ỗ': c_ỡ,
		'ộ': c_ợ,

		'u': c_ư,
		'ú': c_ứ,
		'ù': c_ừ,
		'ủ': c_ử,
		'ũ': c_ữ,
		'ụ': c_ự,

		//'ơ': r_o,
		//'ờ': r_ó,
		//'ớ': r_ò,
		//'ở': r_ỏ,
		//'ỡ': r_õ,
		//'ợ': r_ọ,
		//
		//'ư': r_u,
		//'ứ': r_ú,
		//'ừ': r_ù,
		//'ử': r_ủ,
		//'ữ': r_ũ,
		//'ự': r_ụ,
	}
	replaceCharMap['8'] = map[rune]*CR{
		'a': c_ă,
		'á': c_ắ,
		'à': c_ằ,
		'ả': c_ẳ,
		'ã': c_ẵ,
		'ạ': c_ặ,

		//'ă': r_a,
		//'ắ': r_á,
		//'ằ': r_à,
		//'ẳ': r_ả,
		//'ẵ': r_ã,
		//'ặ': r_ạ,
	}
	replaceCharMap['9'] = replaceCharMap['d']

	replaceStrMap['1'] = replaceStrMap['s']
	replaceStrMap['2'] = replaceStrMap['f']
	replaceStrMap['3'] = replaceStrMap['r']
	replaceStrMap['4'] = replaceStrMap['x']
	replaceStrMap['5'] = replaceStrMap['j']
	replaceStrMap['7'] = replaceStrMap['w']
}
