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

package main

type Runes struct {
	buf []rune
}

func (r *Runes) Len() int {
	return len(r.buf)
}

func (r *Runes) Append(c ...rune) {
	r.buf = append(r.buf, c...)
}

func (r *Runes) AppendRunes(rs Runes) {
	r.buf = append(r.buf, rs.buf...)
}

func (r *Runes) Clear() {
	r.buf = r.buf[:0]
}

func (r *Runes) At(index int) rune {
	if index < len(r.buf) {
		return r.buf[index]
	}
	return 0
}

func (r *Runes) First() rune {
	if len(r.buf) > 0 {
		return r.buf[0]
	}
	return 0
}

func (r *Runes) Last() rune {
	if len(r.buf) > 0 {
		return r.buf[len(r.buf)-1]
	}
	return 0
}
