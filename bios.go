// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

// BIOS information.
type BIOS struct {
	Vendor  string `json:"vendor,omitempty" cty:"vendor" hcl:"vendor,optional"`
	Version string `json:"version,omitempty" cty:"version" hcl:"version,optional"`
	Date    string `json:"date,omitempty" cty:"date" hcl:"date,optional"`
}

func (si *SysInfo) getBIOSInfo() {
	si.BIOS.Vendor = slurpFile("/sys/class/dmi/id/bios_vendor")
	si.BIOS.Version = slurpFile("/sys/class/dmi/id/bios_version")
	si.BIOS.Date = slurpFile("/sys/class/dmi/id/bios_date")
}
