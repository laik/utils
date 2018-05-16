// Copyright 2016 The laik Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package utils

import (
	"net"
)

// 将ip地址转换为长整型
func IP2Long(str string) int64 {
	if ip := net.ParseIP(str); ip != nil {
		var n uint32
		ipBytes := ip.To4()
		for i := uint8(0); i <= 3; i++ {
			n |= uint32(ipBytes[i]) << ((3 - i) * 8)
		}
		return int64(n)
	}
	return 0
}

// 将长整型转换为ip地址
func Long2IP(in int64) string {
	ipBytes := net.IP{}
	for i := uint(0); i <= 3; i++ {
		ipBytes = append(ipBytes, byte(in>>((3-i)*8)))
	}
	return ipBytes.String()
}
