// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

// Package sysinfo is a Go library providing Linux OS / kernel / hardware system information.
package sysinfo

// SysInfo struct encapsulates all other information structs.
type SysInfo struct {
	Meta    Meta            `json:"sysinfo" hcl:"sysinfo,block"`
	Node    Node            `json:"node" hcl:"node,block"`
	OS      OS              `json:"os" hcl:"os,block"`
	Kernel  Kernel          `json:"kernel" hcl:"kernel,block"`
	Product Product         `json:"product" hcl:"product,block"`
	Board   Board           `json:"board" hcl:"board,block"`
	Chassis Chassis         `json:"chassis" hcl:"chassis,block"`
	BIOS    BIOS            `json:"bios" hcl:"bios,block"`
	CPU     CPU             `json:"cpu" hcl:"cpu,block"`
	Memory  Memory          `json:"memory" hcl:"memory,block"`
	Storage []StorageDevice `json:"storage,omitempty" hcl:"storage,block"`
	Network []NetworkDevice `json:"network,omitempty" hcl:"network,block"`
}

// GetSysInfo gathers all available system information.
func (si *SysInfo) GetSysInfo() {
	// Meta info
	si.getMetaInfo()

	// DMI info
	si.getProductInfo()
	si.getBoardInfo()
	si.getChassisInfo()
	si.getBIOSInfo()

	// SMBIOS info
	si.getMemoryInfo()

	// Node info
	si.getNodeInfo() // depends on BIOS info

	// Hardware info
	si.getCPUInfo() // depends on Node info
	si.getStorageInfo()
	si.getNetworkInfo()

	// Software info
	si.getOSInfo()
	si.getKernelInfo()
}
