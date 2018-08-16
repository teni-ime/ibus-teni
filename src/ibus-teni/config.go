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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"teni"
)

const configFile = "%s/.config/ibus/ibus-%s.config.json"

type ToneType uint8

const (
	ConfigToneStd ToneType = iota << 0
	ConfigToneNew ToneType = iota
)

type Config struct {
	InputMethod teni.InputMethod
	ToneType    ToneType
}

func LoadConfig(engineName string) *Config {
	u, err := user.Current()
	if err == nil {
		data, err := ioutil.ReadFile(fmt.Sprintf(configFile, u.HomeDir, engineName))
		if err == nil {
			c := Config{}
			err = json.Unmarshal(data, &c)
			if err == nil {
				return &c
			}
		}
	}
	return &Config{InputMethod: teni.IMTeni}
}

func SaveConfig(c *Config, engineName string) {
	u, err := user.Current()
	if err != nil {
		return
	}

	data, err := json.Marshal(c)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf(configFile, u.HomeDir, engineName), data, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

}
