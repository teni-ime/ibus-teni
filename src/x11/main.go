package main

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/Xatom.h>
#include <stdlib.h>

char* ucharfree(unsigned char* uc) {
	free(uc);
}

char* uchar2char(unsigned char* uc) {
	return (char*)uc;
}

unsigned long uchar2long(unsigned char* uc) {
	return *(unsigned long*)(uc);
}
*/
import "C"
import (
	"fmt"
)

const MaxPropertyLen = 128

func GetUCharProperty(display *C.Display, window C.Window, propName string) *C.uchar {
	var actualType C.Atom
	var actualFormat C.int
	var nItems, bytesAfter C.ulong
	var prop *C.uchar

	filterAtom := C.XInternAtom(display, C.CString(propName), C.True)

	status := C.XGetWindowProperty(display, window, filterAtom, 0, MaxPropertyLen, C.False, C.AnyPropertyType, &actualType, &actualFormat, &nItems, &bytesAfter, &prop)

	if status == C.Success {
		return prop
	}

	return nil
}

func GetStringProperty(display *C.Display, window C.Window, propName string) string {
	prop := GetUCharProperty(display, window, propName)
	if prop != nil {
		defer C.ucharfree(prop)
		return C.GoString(C.uchar2char(prop))
	}

	return ""
}

func GetLongProperty(display *C.Display, window C.Window, propName string) C.ulong {
	prop := GetUCharProperty(display, window, propName)
	if prop != nil {
		defer C.ucharfree(prop)
		return C.uchar2long(prop)
	}

	return 0
}

func OpenDisplay() *C.Display {
	return C.XOpenDisplay(nil)
}

func GetRootWindow(display *C.Display) C.Window {
	return C.XRootWindow(display, C.XDefaultScreen(display))
}

func CloseDisplay(d *C.Display) {
	C.XCloseDisplay(d)
}

func main() {
	display := OpenDisplay()

	rootWindow := GetRootWindow(display)
	fmt.Println(rootWindow)

	activeWindow := GetLongProperty(display, rootWindow, "_NET_ACTIVE_WINDOW")
	fmt.Printf("_NET_ACTIVE_WINDOW: 0x%x\n", activeWindow)
	fmt.Printf("WM_CLASS: %s\n", GetStringProperty(display, C.Window(activeWindow), "WM_CLASS"))
	fmt.Printf("WM_NAME: %s\n", GetStringProperty(display, C.Window(activeWindow), "WM_NAME"))
	fmt.Printf("_NET_WM_NAME: %s\n", GetStringProperty(display, C.Window(activeWindow), "_NET_WM_NAME"))

	defer CloseDisplay(display)

}
