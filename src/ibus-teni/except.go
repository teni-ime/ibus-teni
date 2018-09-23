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
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type ExceptMap struct {
	sync.RWMutex
	m          map[string]bool
	enable     bool
	engineName string
}

func (e *ExceptMap) Contains(ss []string) bool {
	e.RLock()
	defer e.RUnlock()

	if !e.enable || len(e.m) == 0 {
		return false
	}

	for _, s := range ss {
		if _, exist := e.m[s]; exist {
			return true
		}
	}

	return false
}

func (e *ExceptMap) update(exceptFile string) {
	b, err := ioutil.ReadFile(exceptFile)
	if err != nil {
		log.Println(err)
		return
	}

	e.Lock()
	e.m = map[string]bool{}
	for _, s := range strings.Split(string(b), "\n") {
		s = strings.TrimSpace(s)
		if len(s) > 0 && !strings.HasPrefix(s, "#") {
			e.m[s] = true
		}
	}
	e.Unlock()
}

func (e *ExceptMap) Enable() {
	e.Lock()
	e.enable = true

	go func() {
		cont := true
		modTime := time.Now()

		efPath := getExceptListFile(e.engineName)

		for cont {
			if sta, _ := os.Stat(efPath); sta != nil {
				if newModeTime := sta.ModTime(); !newModeTime.Equal(modTime) {
					modTime = newModeTime
					e.update(efPath)
				}
			}
			time.Sleep(time.Second)
			e.RLock()
			cont = e.enable
			e.RUnlock()
		}
	}()

	e.Unlock()
}

func (e *ExceptMap) Disable() {
	e.Lock()
	e.enable = false
	e.m = nil
	e.Unlock()
}
