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
	"fmt"
	"github.com/godbus/dbus"
	"github.com/sarim/goibus/ibus"
	"log"
	"os/exec"
	"runtime/debug"
	"sync"
	"teni"
	"time"
)

const (
	DiffNumpadKeypad = IBUS_KP_0 - IBUS_0
)

type IBusTeniEngine struct {
	sync.Mutex
	ibus.Engine
	preediter      *teni.Engine
	excepted       bool
	zeroLocation   bool
	capSurrounding bool
	engineName     string
	config         *Config
	propList       *ibus.PropList
	exceptMap      *ExceptMap
	display        CDisplay
}

var (
	DictStdList = []string{DictVietnameseCm, DictVietnameseSp, DictVietnameseStd}
	DictNewList = []string{DictVietnameseCm, DictVietnameseSp, DictVietnameseNew}
)

func IBusTeniEngineCreator(conn *dbus.Conn, engineName string) dbus.ObjectPath {

	objectPath := dbus.ObjectPath(fmt.Sprintf("/org/freedesktop/IBus/Engine/Teni/%d", time.Now().UnixNano()))

	var config = LoadConfig(engineName)
	if config.ToneType == ConfigToneStd {
		teni.InitWordTrie(DictStdList...)
	} else {
		teni.InitWordTrie(DictNewList...)
	}

	engine := &IBusTeniEngine{
		Engine:     ibus.BaseEngine(conn, objectPath),
		preediter:  teni.NewEngine(),
		engineName: engineName,
		config:     config,
		propList:   GetPropListByConfig(config),
		exceptMap:  &ExceptMap{engineName: engineName},
	}
	engine.preediter.InputMethod = config.InputMethod
	if config.EnableExcept == ibus.PROP_STATE_CHECKED {
		engine.exceptMap.Enable()
	}
	ibus.PublishEngine(conn, objectPath, engine)

	onMouseClick = func() {
		engine.Lock()
		defer engine.Unlock()
		if engine.preediter.RawKeyLen() > 0 {
			//engine.preediter.Reset()
			//engine.HidePreeditText()

			x11Sync(engine.display)
			x11Flush(engine.display)
			engine.ForwardKeyEvent(IBUS_BackSpace, 14, 0)
			x11Sync(engine.display)
			x11Flush(engine.display)
			engine.ForwardKeyEvent(IBUS_BackSpace, 14, IBUS_RELEASE_MASK)
			x11Sync(engine.display)
			x11Flush(engine.display)
		}
	}

	return objectPath
}


func (e *IBusTeniEngine) SendBackSpace(state uint32, n int) {
	log.Println("SendBackSpace", n)
	if n == 0 {
		return
	}

	if e.capSurrounding {
		log.Println("DeleteSurroundingText")
		e.DeleteSurroundingText(-int32(n), uint32(n))
	} else {
		log.Println("ForwardKyEvent")
		x11Flush(e.display)
		x11Sync(e.display)
		for i := 0; i < n; i++ {
			x11Sync(e.display)
			x11Flush(e.display)
			x11Backspace()

			//e.ForwardKeyEvent(IBUS_BackSpace, 14, state)
			//x11Sync(e.display)
			//x11Flush(e.display)
			//e.ForwardKeyEvent(IBUS_BackSpace, 14, state|IBUS_RELEASE_MASK)
			//x11Sync(e.display)
			//x11Flush(e.display)
		}
	}
}


func (e *IBusTeniEngine) SendText(rs []rune) {
	log.Println("Send key", string(rs))

	//x11Sync(e.display)
	//x11Flush(e.display)
	//e.UpdatePreeditText(ibus.NewText(string(rs)), uint32(len(rs)), true)

	//x11Flush(e.display)
	//x11Sync(e.display)
	e.HidePreeditText()

	x11Flush(e.display)
	x11Sync(e.display)
	e.CommitText(ibus.NewText(string(rs)))
}



func (e *IBusTeniEngine) updatePreedit(newRunes, oldRunes []rune, state uint32) {

	oldLen := len(oldRunes)
	newLen := len(newRunes)
	minLen := oldLen
	if newLen < minLen {
		minLen = newLen
	}

	sameTo := -1
	for i := 0; i < minLen; i++ {
		if oldRunes[i] == newRunes[i] {
			sameTo = i
		} else {
			break
		}
	}
	diffFrom := sameTo + 1

	log.Println(string(oldRunes))
	log.Println(string(newRunes))
	log.Println(diffFrom)

	nBackSpace := 0
	if diffFrom < newLen && diffFrom < oldLen {
		e.SendText([]rune{0x200A}) // https://en.wikipedia.org/wiki/Whitespace_character
		nBackSpace += 1
	}

	if diffFrom < oldLen {
		nBackSpace += oldLen - diffFrom
	}

	e.SendBackSpace(state, nBackSpace)

	e.SendText(newRunes[diffFrom:])
}

func (e *IBusTeniEngine) commitPreedit(lastKey uint32) bool {
	e.preediter.Reset()
	return false
}

func (e *IBusTeniEngine) ProcessKeyEvent(keyVal uint32, keyCode uint32, state uint32) (bool, *dbus.Error) {
	e.Lock()
	defer e.Unlock()
	log.Println(keyVal, keyCode, state)

	if e.zeroLocation || e.excepted ||
		state&IBUS_RELEASE_MASK != 0 || //Ignore key-up event
		(state&IBUS_SHIFT_MASK == 0 && (keyVal == IBUS_Shift_L || keyVal == IBUS_Shift_R)) { //Ignore 1 shift key
		return false, nil
	}

	if state&IBUS_CONTROL_MASK != 0 ||
		state&IBUS_MOD1_MASK != 0 ||
		state&IBUS_IGNORED_MASK != 0 ||
		state&IBUS_SUPER_MASK != 0 ||
		state&IBUS_HYPER_MASK != 0 ||
		state&IBUS_META_MASK != 0 {
		if e.preediter.RawKeyLen() == 0 {
			//No thing left, just ignore
			return false, nil
		} else {
			//while typing, do not process control keys
			return true, nil
		}
	}

	if keyVal == IBUS_BackSpace {
		if e.preediter.RawKeyLen() > 0 {
			oldRunes := e.preediter.GetResult()
			e.preediter.Backspace()
			newRunes := e.preediter.GetResult()
			e.updatePreedit(newRunes, oldRunes, state)
			return true, nil
		}

		//No thing left, just ignore
		return false, nil
	}

	if keyVal == IBUS_Return || keyVal == IBUS_KP_Enter {
		if e.preediter.ResultLen() > 0 {
			e.commitPreedit(keyVal)
			if e.capSurrounding {
				return false, nil
			}
			e.ForwardKeyEvent(keyVal, keyCode, state)
			return true, nil
		} else {
			return false, nil
		}
	}

	if keyVal == IBUS_Escape {
		if e.preediter.RawKeyLen() > 0 {
			e.commitPreedit(keyVal)
			return true, nil
		}
	}

	if keyVal == IBUS_space || keyVal == IBUS_KP_Space {
		if e.preediter.ResultLen() > 0 {
			e.commitPreedit(0)
			if e.capSurrounding {
				return false, nil
			}
			e.ForwardKeyEvent(keyVal, keyCode, state)
			return true, nil
		}
	}

	if e.preediter.RawKeyLen() > 2*teni.MaxWordLength {
		e.commitPreedit(keyVal)
		return true, nil
	}

	if (keyVal >= 'a' && keyVal <= 'z') ||
		(keyVal >= 'A' && keyVal <= 'Z') ||
		(keyVal >= '0' && keyVal <= '9' && e.preediter.ResultLen() > 0) ||
		(e.preediter.InputMethod == teni.IMTelex && teni.InChangeCharMap(rune(keyVal))) {
		if e.preediter.InputMethod == teni.IMTelex && state&IBUS_LOCK_MASK != 0 {
			keyVal = teni.SwitchCaplock(keyVal)
		}
		keyRune := rune(keyVal)
		e.preediter.AddKey(keyRune)
		newRunes := e.preediter.GetResult()
		oldRunes := e.preediter.GetPrevResult()
		e.updatePreedit(newRunes, oldRunes, state)
		return true, nil
	} else {
		if e.preediter.ResultLen() > 0 {
			if e.commitPreedit(keyVal) {
				//lastKey already appended to commit string
				return true, nil
			} else {
				//forward lastKey
				if e.capSurrounding {
					return false, nil
				}
				e.ForwardKeyEvent(keyVal, keyCode, state)
				return true, nil
			}
		}
		//pre-edit empty, just forward key
		return false, nil
	}
}

func (e *IBusTeniEngine) FocusIn() *dbus.Error {
	e.Lock()
	if e.display == nil {
		e.display = x11OpenDisplay()
	}
	if e.config.EnableExcept == ibus.PROP_STATE_CHECKED && e.display != nil {
		awc := x11GetFocusWindowClass(e.display)
		log.Println(awc)
		e.excepted = e.exceptMap.Contains(awc)
	}
	e.preediter.Reset()
	e.Unlock()

	e.RegisterProperties(e.propList)

	return nil
}

func (e *IBusTeniEngine) FocusOut() *dbus.Error {
	e.Lock()
	defer e.Unlock()

	e.preediter.Reset()

	return nil
}

func (e *IBusTeniEngine) Reset() *dbus.Error {
	e.Lock()
	defer e.Unlock()

	e.preediter.Reset()

	return nil
}

func (e *IBusTeniEngine) Enable() *dbus.Error {
	return nil
}

func (e *IBusTeniEngine) Disable() *dbus.Error {
	e.Lock()
	defer e.Unlock()

	if e.display != nil {
		x11CloseDisplay(e.display)
		e.display = nil
	}

	return nil
}

func (e *IBusTeniEngine) SetCapabilities(cap uint32) *dbus.Error {
	e.Lock()
	defer e.Unlock()

	e.capSurrounding = cap&IBUS_CAP_SURROUNDING_TEXT != 0
	return nil
}

func (e *IBusTeniEngine) SetCursorLocation(x int32, y int32, w int32, h int32) *dbus.Error {
	return nil
}

func (e *IBusTeniEngine) SetContentType(purpose uint32, hints uint32) *dbus.Error {
	return nil
}

//@method(in_signature="su")
func (e *IBusTeniEngine) PropertyActivate(propName string, propState uint32) *dbus.Error {
	debug.FreeOSMemory()

	if propName == PropKeyAbout {
		exec.Command("xdg-open", HomePage).Start()
		return nil
	}

	oldToneType := e.config.ToneType

	if propState == ibus.PROP_STATE_CHECKED &&
		(propName == PropKeyMethodTeni ||
			propName == PropKeyMethodVni ||
			propName == PropKeyMethodTelex ||
			propName == PropKeyToneStd ||
			propName == PropKeyToneNew) {
		switch propName {
		case PropKeyMethodTeni:
			e.config.InputMethod = teni.IMTeni
			e.preediter.InputMethod = teni.IMTeni
		case PropKeyMethodVni:
			e.config.InputMethod = teni.IMVni
			e.preediter.InputMethod = teni.IMVni
		case PropKeyMethodTelex:
			e.config.InputMethod = teni.IMTelex
			e.preediter.InputMethod = teni.IMTelex
		case PropKeyToneStd:
			e.config.ToneType = ConfigToneStd
		case PropKeyToneNew:
			e.config.ToneType = ConfigToneNew
		}
		SaveConfig(e.config, e.engineName)
		e.propList = GetPropListByConfig(e.config)
		if e.config.ToneType != oldToneType {
			if e.config.ToneType == ConfigToneStd {
				teni.InitWordTrie(DictStdList...)
			} else {
				teni.InitWordTrie(DictNewList...)
			}
		}
		return nil
	}

	if propName == PropKeyExcept {
		e.config.EnableExcept = propState
		SaveConfig(e.config, e.engineName)
		e.propList = GetPropListByConfig(e.config)
		if propState == ibus.PROP_STATE_CHECKED {
			e.exceptMap.Enable()
			e.excepted = e.exceptMap.Contains(x11GetFocusWindowClass(e.display))
		} else {
			e.exceptMap.Disable()
			e.excepted = false
		}
		return nil
	}

	if propName == PropKeyExceptList {
		OpenExceptListFile(e.engineName)
		return nil
	}

	return nil
}
