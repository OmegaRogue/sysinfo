// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

//go:build linux
// +build linux

package sysinfo

import (
	"strings"
	"syscall"
	"unsafe"
)

// Kernel information.
type Kernel struct {
	Release      string `json:"release,omitempty" cty:"release" hcl:"release,optional"`
	Version      string `json:"version,omitempty" cty:"version" hcl:"version,optional"`
	Architecture string `json:"architecture,omitempty" cty:"architecture" hcl:"architecture,optional"`
}

func (si *SysInfo) getKernelInfo() {
	si.Kernel.Release = slurpFile("/proc/sys/kernel/osrelease")
	si.Kernel.Version = slurpFile("/proc/sys/kernel/version")

	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		return
	}

	si.Kernel.Architecture = strings.TrimRight(string((*[65]byte)(unsafe.Pointer(&uname.Machine))[:]), "\000")
}
