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

/*
#cgo CFLAGS: -std=c99
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/Xatom.h>

inline void ucharfree(unsigned char* uc) {
	XFree(uc);
}

inline void windowfree(Window* w) {
	XFree(w);
}

inline char* uchar2char(unsigned char* uc, unsigned long len) {
	for (int i=0; i<len; i++) {
		if (uc[i] == 0) {
			uc[i] = '\n';
		}
	}
	return (char*)uc;
}

inline unsigned long uchar2long(unsigned char* uc) {
	return *(unsigned long*)(uc);
}

static int ignore_x_error(Display *display, XErrorEvent *error) {
    return 0;
}

void setXIgnoreErrorHandler() {
	XSetErrorHandler(ignore_x_error);
}

*/
import "C"
import (
	"strings"
)

const (
	MaxPropertyLen = 128
	MaxCacheWM     = 16

	WM_CLASS = "WM_CLASS"
)

type CDisplay *C.Display

var cacheWM = NewCacheWM(MaxCacheWM)

func init() {
	C.setXIgnoreErrorHandler()
}

func x11GetUCharProperty(display *C.Display, window C.Window, propName string) (*C.uchar, C.ulong) {
	var actualType C.Atom
	var actualFormat C.int
	var nItems, bytesAfter C.ulong
	var prop *C.uchar

	filterAtom := C.XInternAtom(display, C.CString(propName), C.True)

	status := C.XGetWindowProperty(display, window, filterAtom, 0, MaxPropertyLen, C.False, C.AnyPropertyType, &actualType, &actualFormat, &nItems, &bytesAfter, &prop)

	if status == C.Success {
		return prop, nItems
	}

	return nil, 0
}

func x11GetStringProperty(display *C.Display, window C.Window, propName string) string {
	prop, propLen := x11GetUCharProperty(display, window, propName)
	if prop != nil {
		defer C.ucharfree(prop)
		return C.GoString(C.uchar2char(prop, propLen))
	}

	return ""
}

func x11OpenDisplay() *C.Display {
	return C.XOpenDisplay(nil)
}

func x11GetInputFocus(display *C.Display) C.Window {
	var window C.Window
	var revertTo C.int
	C.XGetInputFocus(display, &window, &revertTo)

	return window
}

func x11GetParentWindow(display *C.Display, w C.Window) (rootWindow, parentWindow C.Window) {
	var childrenWindows *C.Window
	var nChild C.uint
	C.XQueryTree(display, w, &rootWindow, &parentWindow, &childrenWindows, &nChild)
	C.windowfree(childrenWindows)

	return
}

func x11CloseDisplay(d *C.Display) {
	C.XCloseDisplay(d)
}

func x11GetFocusWindowClass(display *C.Display) []string {
	if display != nil {

		focusWindow := x11GetInputFocus(display)
		if v, hit := cacheWM.Get(uint32(focusWindow)); hit {
			return v
		}
		strClass := ""
		w := focusWindow
		for {
			s := x11GetStringProperty(display, w, WM_CLASS)

			if len(s) > 0 {
				strClass += s + "\n"
			}

			rootWindow, parentWindow := x11GetParentWindow(display, w)

			if rootWindow == parentWindow {
				break
			}

			w = parentWindow
		}

		v := strings.Split(strClass, "\n")
		if len(v) > 0 {
			cacheWM.Set(uint32(focusWindow), v)
		}

		return v
	}

	return nil
}
