package main

import (
	"log"
	"os"
)

//sudo chmod +r /dev/input/mice
const (
	DevInputMice = "/dev/input/mice"
)

var onMouseClick func()

func init() {
	go func() {
		pressing := false
		miceDev, err := os.OpenFile(DevInputMice, os.O_RDONLY, 0)
		if err == nil {
			data := make([]byte, 3)
			for {
				n, err := miceDev.Read(data)
				log.Println(data)
				if err == nil && n == 3 && data[0]&0x7 != 0 {
					if !pressing {
						if onMouseClick != nil {
							go onMouseClick()
						}
						pressing = true
					}
				} else if pressing {
					pressing = false
				}
			}
		}
	}()
}
