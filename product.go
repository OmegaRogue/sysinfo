// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

// Product information.
type Product struct {
	Name    string `json:"name,omitempty" cty:"name" hcl:"name,optional"`
	Vendor  string `json:"vendor,omitempty" cty:"vendor" hcl:"vendor,optional"`
	Version string `json:"version,omitempty" cty:"version" hcl:"version,optional"`
	Serial  string `json:"serial,omitempty" cty:"serial" hcl:"serial,optional"`
}

func (si *SysInfo) getProductInfo() {
	si.Product.Name = slurpFile("/sys/class/dmi/id/product_name")
	si.Product.Vendor = slurpFile("/sys/class/dmi/id/sys_vendor")
	si.Product.Version = slurpFile("/sys/class/dmi/id/product_version")
	si.Product.Serial = slurpFile("/sys/class/dmi/id/product_serial")
}
