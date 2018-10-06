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
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"teni"
)

const (
	configFile           = "%s/.config/ibus/ibus-%s.config.json"
	exceptListFile       = "%s/.config/ibus/ibus-%s.except.txt"
	sampleExceptListFile = "except.tmpl.txt"

	varWmBash  = "${WM.BASH}"
	wmBashFile = "wm.bash"
)

type ToneType uint8

const (
	ConfigToneStd ToneType = iota << 0
	ConfigToneNew ToneType = iota
)

type Config struct {
	InputMethod      teni.InputMethod
	ToneType         ToneType
	EnableExcept     uint32
	EnableLongText   uint32
	EnableForceSpell uint32
}

func LoadConfig(engineName string) *Config {
	c := Config{
		InputMethod:      teni.IMTeni,
		ToneType:         ConfigToneStd,
		EnableExcept:     0,
		EnableLongText:   0,
		EnableForceSpell: 1,
	}

	u, err := user.Current()
	if err == nil {
		data, err := ioutil.ReadFile(fmt.Sprintf(configFile, u.HomeDir, engineName))
		if err == nil {
			json.Unmarshal(data, &c)
		}
	}

	return &c
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

	err = ioutil.WriteFile(fmt.Sprintf(configFile, u.HomeDir, engineName), data, 0644)
	if err != nil {
		log.Println(err)
	}

}

func getExceptListFile(engineName string) string {
	u, err := user.Current()
	if err != nil {
		return fmt.Sprintf(exceptListFile, "~", engineName)
	}

	return fmt.Sprintf(exceptListFile, u.HomeDir, engineName)
}

func getEngineSubFile(fileName string) string {
	if _, err := os.Stat(fileName); err == nil {
		if absPath, err := filepath.Abs(fileName); err == nil {
			return absPath
		}
	}

	return filepath.Join(filepath.Dir(os.Args[0]), fileName)
}

func OpenExceptListFile(engineName string) {
	efPath := getExceptListFile(engineName)
	if _, err := os.Stat(efPath); os.IsNotExist(err) {
		sampleFile := getEngineSubFile(sampleExceptListFile)
		sample, _ := ioutil.ReadFile(sampleFile)
		if len(sample) > 0 {
			wmBashPath := getEngineSubFile(wmBashFile)
			strSample := strings.Replace(string(sample), varWmBash, wmBashPath, 1)
			sample = []byte(strSample)
		}
		ioutil.WriteFile(efPath, sample, 0644)
	}

	exec.Command("xdg-open", efPath).Start()
}
