//go:build darwin
// +build darwin

package sysinfo

// Kernel information.
type Kernel struct {
	Release      string `json:"release,omitempty" cty:"release" hcl:"release,optional"`
	Version      string `json:"version,omitempty" cty:"version" hcl:"version,optional"`
	Architecture string `json:"architecture,omitempty" cty:"architecture" hcl:"architecture,optional"`
}

func (si *SysInfo) getKernelInfo() {
}
