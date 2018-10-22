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

import (
	"os"
)

//sudo usermod -a -G input $USER
const (
	DevInputMice = "/dev/input/mice"
)

var onMouseClick func()

func init() {
	go func() {
		down := false
		miceDev, err := os.OpenFile(DevInputMice, os.O_RDONLY, 0)
		if err == nil {
			data := make([]byte, 3)
			for {
				n, err := miceDev.Read(data)
				if err == nil && n == 3 && data[0]&0x7 != 0 {
					if data[1] == 0 && data[2] == 0 {
						if !down {
							if onMouseClick != nil {
								go onMouseClick()
							}
							down = true
						}
					}
				} else if down {
					down = false
				}
			}
		}
	}()
}
