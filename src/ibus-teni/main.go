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
	"github.com/godbus/dbus"
	"github.com/sarim/goibus/ibus"
	"log"
	"os"
)

const (
	ComponentName = "org.freedesktop.IBus.Teni"
	EngineName    = "Teni"
	HomePage      = "https://github.com/teni-ime/ibus-teni"

	DictVietnameseCm  = "dict/vietnamese.cm.dict"
	DictVietnameseSp  = "dict/vietnamese.sp.dict"
	DictVietnameseStd = "dict/vietnamese.std.dict"
	DictVietnameseNew = "dict/vietnamese.new.dict"
)

func main() {
	if isIBusDaemonChild() {
		if len(os.Args) == 3 && os.Args[1] == "cd" {
			os.Chdir(os.Args[2])
		}
		bus := ibus.NewBus()
		bus.RequestName(ComponentName, 0)

		conn := bus.GetDbusConn()
		ibus.NewFactory(conn, IBusTeniEngineCreator)

		select {}
	} else {
		log.Println("Running debug mode")
		runMode = " (debug)"

		bus := ibus.NewBus()
		bus.RegisterComponent(makeDebugComponent())

		conn := bus.GetDbusConn()
		ibus.NewFactory(conn, IBusTeniEngineCreator)

		log.Println("Setting Global Engine to", DebugEngineName)
		bus.CallMethod("SetGlobalEngine", 0, DebugEngineName)

		c := make(chan *dbus.Signal, 10)
		conn.Signal(c)

		select {
		case <-c:
		}
	}
}
