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
	"teni"
)

const (
	PropKeyAbout       = "about"
	PropKeyMethodTeni  = "method_teni"
	PropKeyMethodVni   = "method_vni"
	PropKeyMethodTelex = "method_telex"
	PropKeyToneStd     = "tone_std"
	PropKeyToneNew     = "tone_new"
)

var runMode = ""

func GetPropListByConfig(c *Config) *ibus.PropList {
	teniChecked := ibus.PROP_STATE_UNCHECKED
	vniChecked := ibus.PROP_STATE_UNCHECKED
	telexChecked := ibus.PROP_STATE_UNCHECKED
	toneStdChecked := ibus.PROP_STATE_UNCHECKED
	toneNewChecked := ibus.PROP_STATE_UNCHECKED

	switch c.InputMethod {
	case teni.IMTeni:
		teniChecked = ibus.PROP_STATE_CHECKED
	case teni.IMVni:
		vniChecked = ibus.PROP_STATE_CHECKED
	case teni.IMTelex:
		telexChecked = ibus.PROP_STATE_CHECKED
	}
	switch c.ToneType {
	case ConfigToneStd:
		toneStdChecked = ibus.PROP_STATE_CHECKED
	case ConfigToneNew:
		toneNewChecked = ibus.PROP_STATE_CHECKED
	}

	return ibus.NewPropList(
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       "about",
			Type:      ibus.PROP_TYPE_NORMAL,
			Label:     dbus.MakeVariant(ibus.NewText("Bộ gõ " + EngineName + " " + Version + runMode)),
			Tooltip:   dbus.MakeVariant(ibus.NewText("Mở trang chủ")),
			Sensitive: true,
			Visible:   true,
			Icon:      "gtk-about",
			Symbol:    dbus.MakeVariant(ibus.NewText("A")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       "-",
			Type:      ibus.PROP_TYPE_SEPARATOR,
			Label:     dbus.MakeVariant(ibus.NewText("")),
			Tooltip:   dbus.MakeVariant(ibus.NewText("")),
			Sensitive: true,
			Visible:   true,
			Symbol:    dbus.MakeVariant(ibus.NewText("")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       PropKeyMethodTeni,
			Type:      ibus.PROP_TYPE_RADIO,
			Label:     dbus.MakeVariant(ibus.NewText("Kiểu gõ Teni")),
			Tooltip:   dbus.MakeVariant(ibus.NewText("Kết hợp Telex và Vni")),
			Sensitive: true,
			Visible:   true,
			State:     teniChecked,
			Symbol:    dbus.MakeVariant(ibus.NewText("T")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       PropKeyMethodVni,
			Type:      ibus.PROP_TYPE_RADIO,
			Label:     dbus.MakeVariant(ibus.NewText("Kiểu gõ Vni")),
			Tooltip:   dbus.MakeVariant(ibus.NewText("Chỉ kiểu gõ Vni")),
			Sensitive: true,
			Visible:   true,
			State:     vniChecked,
			Symbol:    dbus.MakeVariant(ibus.NewText("V")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       PropKeyMethodTelex,
			Type:      ibus.PROP_TYPE_RADIO,
			Label:     dbus.MakeVariant(ibus.NewText("Kiểu gõ Telex")),
			Tooltip:   dbus.MakeVariant(ibus.NewText("Chỉ kiểu gõ Telex")),
			Sensitive: true,
			Visible:   true,
			State:     telexChecked,
			Symbol:    dbus.MakeVariant(ibus.NewText("TX")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       "-",
			Type:      ibus.PROP_TYPE_SEPARATOR,
			Label:     dbus.MakeVariant(ibus.NewText("")),
			Tooltip:   dbus.MakeVariant(ibus.NewText("")),
			Sensitive: true,
			Visible:   true,
			Symbol:    dbus.MakeVariant(ibus.NewText("")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       PropKeyToneStd,
			Type:      ibus.PROP_TYPE_RADIO,
			Label:     dbus.MakeVariant(ibus.NewText("Dấu thanh chuẩn")),
			Tooltip:   dbus.MakeVariant(ibus.NewText("Cân đối, nên dùng")),
			Sensitive: true,
			Visible:   true,
			State:     toneStdChecked,
			Symbol:    dbus.MakeVariant(ibus.NewText("C")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
		&ibus.Property{
			Name:      "IBusProperty",
			Key:       PropKeyToneNew,
			Type:      ibus.PROP_TYPE_RADIO,
			Label:     dbus.MakeVariant(ibus.NewText("Dấu thanh kiểu mới")),
			Tooltip:   dbus.MakeVariant(ibus.NewText("Lệch bên phải, không nên dùng")),
			Sensitive: true,
			Visible:   true,
			State:     toneNewChecked,
			Symbol:    dbus.MakeVariant(ibus.NewText("M")),
			SubProps:  dbus.MakeVariant(*ibus.NewPropList()),
		},
	)
}
