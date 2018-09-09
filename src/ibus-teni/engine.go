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
	"os/exec"
	"teni"
	"time"
)

const (
	DiffNumpadKeypad = IBUS_KP_0 - IBUS_0
)

type IBusTeniEngine struct {
	ibus.Engine
	preediter      *teni.Engine
	enable         bool
	excepted       bool
	capSurrounding bool
	engineName     string
	config         *Config
	propList       *ibus.PropList
	exceptMap      *ExceptMap
	newFocusIn     bool
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
	return objectPath
}

func (e *IBusTeniEngine) updatePreedit() {
	e.UpdatePreeditTextWithMode(ibus.NewText(e.preediter.GetResultStr()), e.preediter.ResultLen(), true, ibus.IBUS_ENGINE_PREEDIT_COMMIT)
}

func (e *IBusTeniEngine) commitPreedit(lastKey uint32) bool {
	var keyAppended = false
	var commitStr string
	if lastKey == IBUS_Escape {
		commitStr = e.preediter.GetRawStr()
	} else {
		commitStr = e.preediter.GetCommitResultStr()
	}
	e.preediter.Reset()

	//Convert num-pad key to normal number
	if (lastKey >= IBUS_KP_0 && lastKey <= IBUS_KP_9) ||
		(lastKey >= IBUS_KP_Multiply && lastKey <= IBUS_KP_Divide) {
		lastKey = lastKey - DiffNumpadKeypad
	}

	if lastKey >= 0x20 && lastKey <= 0xFF {
		//append printable keys
		commitStr += string(lastKey)
		keyAppended = true
	}

	e.HidePreeditText()
	e.CommitText(ibus.NewText(commitStr))

	return keyAppended
}

func (e *IBusTeniEngine) ProcessKeyEvent(keyVal uint32, keyCode uint32, state uint32) (bool, *dbus.Error) {
	if e.config.EnableExcept == ibus.PROP_STATE_CHECKED && e.newFocusIn {
		e.newFocusIn = false
		awc := x11GetFocusWindowClass()
		e.excepted = e.exceptMap.Contains(awc)
	}

	if !e.enable || e.excepted ||
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
			e.preediter.Backspace()
			e.updatePreedit()
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
		e.updatePreedit()
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
	e.RegisterProperties(e.propList)
	e.preediter.Reset()
	e.newFocusIn = true

	return nil
}

func (e *IBusTeniEngine) FocusOut() *dbus.Error {
	e.preediter.Reset()
	e.newFocusIn = true

	return nil
}

func (e *IBusTeniEngine) Reset() *dbus.Error {
	e.preediter.Reset()
	e.newFocusIn = true

	return nil
}

func (e *IBusTeniEngine) Enable() *dbus.Error {
	e.preediter.Reset()
	e.newFocusIn = true

	return nil
}

func (e *IBusTeniEngine) Disable() *dbus.Error {
	e.preediter.Reset()
	e.newFocusIn = true

	return nil
}

func (e *IBusTeniEngine) SetCapabilities(cap uint32) *dbus.Error {
	e.enable = cap&IBUS_CAP_PREEDIT_TEXT != 0
	e.capSurrounding = cap&IBUS_CAP_SURROUNDING_TEXT != 0
	return nil
}

func (e *IBusTeniEngine) SetCursorLocation(x int32, y int32, w int32, h int32) *dbus.Error {
	return nil
}

func (e *IBusTeniEngine) SetContentType(purpose uint32, hints uint32) *dbus.Error {
	e.enable = purpose == IBUS_INPUT_PURPOSE_FREE_FORM ||
		purpose == IBUS_INPUT_PURPOSE_ALPHA ||
		purpose == IBUS_INPUT_PURPOSE_NAME

	return nil
}

//@method(in_signature="su")
func (e *IBusTeniEngine) PropertyActivate(propName string, propState uint32) *dbus.Error {
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
		} else {
			e.exceptMap.Disable()
		}
		return nil
	}

	if propName == PropKeyExceptList {
		OpenExceptListFile(e.engineName)
		return nil
	}

	return nil
}
