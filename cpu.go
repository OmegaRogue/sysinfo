// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// CPU information.
type CPU struct {
	Vendor  string `json:"vendor,omitempty" cty:"vendor" hcl:"vendor,optional"`
	Model   string `json:"model,omitempty" cty:"model" hcl:"model,optional"`
	Speed   uint   `json:"speed,omitempty" cty:"speed" hcl:"speed,optional"`       // CPU clock rate in MHz
	Cache   uint   `json:"cache,omitempty" cty:"cache" hcl:"cache,optional"`       // CPU cache size in KB
	Cpus    uint   `json:"cpus,omitempty" cty:"cpus" hcl:"cpus,optional"`          // number of physical CPUs
	Cores   uint   `json:"cores,omitempty" cty:"cores" hcl:"cores,optional"`       // number of physical CPU cores
	Threads uint   `json:"threads,omitempty" cty:"threads" hcl:"threads,optional"` // number of logical (HT) CPU cores
}

var (
	reTwoColumns = regexp.MustCompile("\t+: ")
	reExtraSpace = regexp.MustCompile(" +")
	reCacheSize  = regexp.MustCompile(`^(\d+) KB$`)
)

func (si *SysInfo) getCPUInfo() {
	si.CPU.Threads = uint(runtime.NumCPU())

	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return
	}
	defer f.Close()

	cpu := make(map[string]bool)
	core := make(map[string]bool)

	var cpuID string

	s := bufio.NewScanner(f)
	for s.Scan() {
		if sl := reTwoColumns.Split(s.Text(), 2); sl != nil {
			switch sl[0] {
			case "physical id":
				cpuID = sl[1]
				cpu[cpuID] = true
			case "core id":
				coreID := fmt.Sprintf("%s/%s", cpuID, sl[1])
				core[coreID] = true
			case "vendor_id":
				if si.CPU.Vendor == "" {
					si.CPU.Vendor = sl[1]
				}
			case "model name":
				if si.CPU.Model == "" {
					// CPU model, as reported by /proc/cpuinfo, can be a bit ugly. Clean up...
					model := reExtraSpace.ReplaceAllLiteralString(sl[1], " ")
					si.CPU.Model = strings.Replace(model, "- ", "-", 1)
				}
			case "cache size":
				if si.CPU.Cache == 0 {
					if m := reCacheSize.FindStringSubmatch(sl[1]); m != nil {
						if cache, err := strconv.ParseUint(m[1], 10, 64); err == nil {
							si.CPU.Cache = uint(cache)
						}
					}
				}
			}
		}
	}
	if s.Err() != nil {
		return
	}

	// getNodeInfo() must have run first, to detect if we're dealing with a virtualized CPU! Detecting number of
	// physical processors and/or cores is totally unreliable in virtualized environments, so let's not do it.
	if si.Node.Hostname == "" || si.Node.Hypervisor != "" {
		return
	}

	si.CPU.Cpus = uint(len(cpu))
	si.CPU.Cores = uint(len(core))
}
