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

var downLvlMap = map[rune]map[rune]uint8{
	'á': {'a': 2},
	'à': {'a': 2},
	'ả': {'a': 2},
	'ã': {'a': 2},
	'ạ': {'a': 2},

	'ă': {'a': 1},
	'ắ': {'ă': 1, 'á': 1, 'a': 2},
	'ằ': {'ă': 1, 'à': 1, 'a': 2},
	'ẳ': {'ă': 1, 'ả': 1, 'a': 2},
	'ẵ': {'ă': 1, 'ã': 1, 'a': 2},
	'ặ': {'ă': 1, 'ạ': 1, 'a': 2},

	'â': {'a': 1},
	'ấ': {'â': 1, 'á': 1, 'a': 2},
	'ầ': {'â': 1, 'à': 1, 'a': 2},
	'ẩ': {'â': 1, 'ả': 1, 'a': 2},
	'ẫ': {'â': 1, 'ã': 1, 'a': 2},
	'ậ': {'â': 1, 'ạ': 1, 'a': 2},

	'é': {'e': 2},
	'è': {'e': 2},
	'ẻ': {'e': 2},
	'ẽ': {'e': 2},
	'ẹ': {'e': 2},

	'ê': {'e': 1},
	'ế': {'ê': 1, 'é': 1, 'e': 2},
	'ề': {'ê': 1, 'è': 1, 'e': 2},
	'ể': {'ê': 1, 'ẻ': 1, 'e': 2},
	'ễ': {'ê': 1, 'ẽ': 1, 'e': 2},
	'ệ': {'ê': 1, 'ẹ': 1, 'e': 2},

	'í': {'i': 2},
	'ì': {'i': 2},
	'ỉ': {'i': 2},
	'ĩ': {'i': 2},
	'ị': {'i': 2},

	'ó': {'o': 2},
	'ò': {'o': 2},
	'ỏ': {'o': 2},
	'õ': {'o': 2},
	'ọ': {'o': 2},

	'ô': {'o': 2},
	'ố': {'ô': 1, 'ó': 1, 'o': 2},
	'ồ': {'ô': 1, 'ò': 1, 'o': 2},
	'ổ': {'ô': 1, 'ỏ': 1, 'o': 2},
	'ỗ': {'ô': 1, 'õ': 1, 'o': 2},
	'ộ': {'ô': 1, 'ọ': 1, 'o': 2},

	'ơ': {'o': 1},
	'ớ': {'ơ': 1, 'ó': 1, 'o': 2},
	'ờ': {'ơ': 1, 'ò': 1, 'o': 2},
	'ở': {'ơ': 1, 'ỏ': 1, 'o': 2},
	'ỡ': {'ơ': 1, 'õ': 1, 'o': 2},
	'ợ': {'ơ': 1, 'ọ': 1, 'o': 2},

	'ú': {'u': 2},
	'ù': {'u': 2},
	'ủ': {'u': 2},
	'ũ': {'u': 2},
	'ụ': {'u': 2},

	'ư': {'u': 1},
	'ứ': {'ư': 1, 'ú': 1, 'u': 2},
	'ừ': {'ư': 1, 'ù': 1, 'u': 2},
	'ử': {'ư': 1, 'ủ': 1, 'u': 2},
	'ữ': {'ư': 1, 'ũ': 1, 'u': 2},
	'ự': {'ư': 1, 'ụ': 1, 'u': 2},

	'ý': {'y': 2},
	'ỳ': {'y': 2},
	'ỷ': {'y': 2},
	'ỹ': {'y': 2},
	'ỵ': {'y': 2},

	'đ': {'d': 2},
}
