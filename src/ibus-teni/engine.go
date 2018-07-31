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
	capSurrounding bool
	engineName     string
	config         *Config
	propList       *ibus.PropList
}

var emptyText = ibus.NewText("")

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
	}
	engine.preediter.NumberOnly = config.InputMethod == ConfigMethodVni

	ibus.PublishEngine(conn, objectPath, engine)
	return objectPath
}

func (e *IBusTeniEngine) updatePreedit() {
	e.UpdatePreeditTextWithMode(ibus.NewText(e.preediter.GetResultStr()), e.preediter.ResultLen(), true, ibus.IBUS_ENGINE_PREEDIT_COMMIT)
}

func (e *IBusTeniEngine) commitPreedit(lastKey uint32) {
	var commitStr string
	if lastKey == IBUS_Escape {
		commitStr = e.preediter.GetRawStr()
	} else {
		commitStr = e.preediter.GetCommitResultStr()
	}
	e.preediter.Reset()

	//log.Printf("lastKey %x, %s", lastKey, string(lastKey))

	//Convert num-pad key to normal number
	if (lastKey >= IBUS_KP_0 && lastKey <= IBUS_KP_9) ||
		(lastKey >= IBUS_KP_Multiply && lastKey <= IBUS_KP_Divide) {
		lastKey = lastKey - DiffNumpadKeypad
	}

	if lastKey >= 0x20 && lastKey <= 0xFF {
		//append printable keys
		commitStr += string(lastKey)
	}

	//log.Printf("CommitText [%s]\n", commitStr)
	e.CommitText(ibus.NewText(commitStr))
	e.UpdatePreeditText(emptyText, 0, true)
}

func (e *IBusTeniEngine) ProcessKeyEvent(keyVal uint32, keyCode uint32, state uint32) (bool, *dbus.Error) {
	//log.Println("ProcessKeyEvent", keyVal, keyCode, state)

	if !e.enable ||
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
			e.commitPreedit(keyVal)
			return true, nil
		}
	}

	if e.preediter.RawKeyLen() > 2*teni.MaxWordLength {
		e.commitPreedit(keyVal)
		return true, nil
	}

	if (keyVal >= 'a' && keyVal <= 'z') ||
		(keyVal >= 'A' && keyVal <= 'Z') ||
		(keyVal >= '0' && keyVal <= '9' && e.preediter.ResultLen() > 0) {
		keyRune := rune(keyVal)

		e.preediter.AddKey(keyRune)
		e.updatePreedit()
		return true, nil
	} else {
		if e.preediter.ResultLen() > 0 {
			e.commitPreedit(keyVal)
			return true, nil
		}
		//pre-edit empty, just append
		return false, nil
	}
}

func (e *IBusTeniEngine) FocusIn() *dbus.Error {
	//log.Println("FocusIn")
	e.RegisterProperties(e.propList)
	e.preediter.Reset()
	return nil
}

func (e *IBusTeniEngine) FocusOut() *dbus.Error {
	//log.Println("FocusOut")
	e.preediter.Reset()
	return nil
}

func (e *IBusTeniEngine) Reset() *dbus.Error {
	//log.Println("Reset")
	e.preediter.Reset()
	return nil
}

func (e *IBusTeniEngine) Enable() *dbus.Error {
	//log.Println("Enable")
	e.preediter.Reset()
	return nil
}

func (e *IBusTeniEngine) Disable() *dbus.Error {
	//log.Println("Disable")
	e.preediter.Reset()
	return nil
}

func (e *IBusTeniEngine) SetCapabilities(cap uint32) *dbus.Error {
	//log.Println("SetCapabilities", cap)
	e.enable = cap&IBUS_CAP_PREEDIT_TEXT != 0
	e.capSurrounding = cap&IBUS_CAP_SURROUNDING_TEXT != 0
	return nil
}

func (e *IBusTeniEngine) SetCursorLocation(x int32, y int32, w int32, h int32) *dbus.Error {
	//log.Println("SetCursorLocation", x, y, w, h)
	return nil
}

func (e *IBusTeniEngine) SetContentType(purpose uint32, hints uint32) *dbus.Error {
	//log.Println("SetContentType", purpose, hints)

	e.enable = purpose == IBUS_INPUT_PURPOSE_FREE_FORM ||
		purpose == IBUS_INPUT_PURPOSE_ALPHA ||
		purpose == IBUS_INPUT_PURPOSE_NAME

	return nil
}

//@method(in_signature="su")
func (e *IBusTeniEngine) PropertyActivate(propName string, propState uint32) *dbus.Error {
	//log.Println("PropertyActivate", propName, propState)
	if propName == PropKeyAbout {
		exec.Command("xdg-open", HomePage).Start()
		return nil
	}

	oldToneType := e.config.ToneType

	if propState == ibus.PROP_STATE_CHECKED &&
		(propName == PropKeyMethodTeni ||
			propName == PropKeyMethodVni ||
			propName == PropKeyToneStd ||
			propName == PropKeyToneNew) {
		switch propName {
		case PropKeyMethodTeni:
			e.config.InputMethod = ConfigMethodTeni
			e.preediter.NumberOnly = false
		case PropKeyMethodVni:
			e.config.InputMethod = ConfigMethodVni
			e.preediter.NumberOnly = true
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
	}
	return nil
}
