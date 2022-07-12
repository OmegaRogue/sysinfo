// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import "time"

// Meta information.
type Meta struct {
	Version   string    `json:"version" cty:"version" hcl:"version,optional"`
	Timestamp time.Time `json:"timestamp" cty:"timestamp" hcl:"timestamp,optional"`
}

func (si *SysInfo) getMetaInfo() {
	si.Meta.Version = Version
	si.Meta.Timestamp = time.Now()
}
