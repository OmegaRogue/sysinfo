// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

// Board information.
type Board struct {
	Name     string `json:"name,omitempty" hcl:"name,optional"`
	Vendor   string `json:"vendor,omitempty" hcl:"vendor,optional"`
	Version  string `json:"version,omitempty" hcl:"version,optional"`
	Serial   string `json:"serial,omitempty" hcl:"serial,optional"`
	AssetTag string `json:"assettag,omitempty" hcl:"assettag,optional"`
}

func (si *SysInfo) getBoardInfo() {
	si.Board.Name = slurpFile("/sys/class/dmi/id/board_name")
	si.Board.Vendor = slurpFile("/sys/class/dmi/id/board_vendor")
	si.Board.Version = slurpFile("/sys/class/dmi/id/board_version")
	si.Board.Serial = slurpFile("/sys/class/dmi/id/board_serial")
	si.Board.AssetTag = slurpFile("/sys/class/dmi/id/board_asset_tag")
}
