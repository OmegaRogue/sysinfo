//go:build darwin
// +build darwin

package sysinfo

// Kernel information.
type Kernel struct {
	Release      string `json:"release,omitempty" hcl:"release,optional"`
	Version      string `json:"version,omitempty" hcl:"version,optional"`
	Architecture string `json:"architecture,omitempty" hcl:"architecture,optional"`
}

func (si *SysInfo) getKernelInfo() {
}
