package main

/*
#cgo CFLAGS: -std=c99
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/Xatom.h>

inline char* ucharfree(unsigned char* uc) {
	XFree(uc);
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
*/
import "C"
import (
	"strings"
)

const (
	MaxPropertyLen = 128

	_NET_ACTIVE_WINDOW = "_NET_ACTIVE_WINDOW"
	WM_CLASS           = "WM_CLASS"
)

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

func x11GetLongProperty(display *C.Display, window C.Window, propName string) C.ulong {
	prop, _ := x11GetUCharProperty(display, window, propName)
	if prop != nil {
		defer C.ucharfree(prop)
		return C.uchar2long(prop)
	}

	return 0
}

func x11OpenDisplay() *C.Display {
	return C.XOpenDisplay(nil)
}

func x11GetRootWindow(display *C.Display) C.Window {
	return C.XRootWindow(display, C.XDefaultScreen(display))
}

func x11CloseDisplay(d *C.Display) {
	C.XCloseDisplay(d)
}

func x11GetActiveWindowClass() []string {
	defer func() {
		recover()
	}()
	display := x11OpenDisplay()
	if display != nil {
		defer x11CloseDisplay(display)

		rootWindow := x11GetRootWindow(display)
		if rootWindow != 0 {
			activeWindow := x11GetLongProperty(display, rootWindow, _NET_ACTIVE_WINDOW)
			if activeWindow != 0 {
				strClass := x11GetStringProperty(display, C.Window(activeWindow), WM_CLASS)
				return strings.Split(strClass, "\n")
			}
		}
	}

	return nil
}
