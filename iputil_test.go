package utils

import "testing"

func TestIpToint(t *testing.T) {
	ip := "192.168.0.1"
	i2i := IP2Long(ip)
	t.Log(i2i)
	l2i := Long2IP(i2i)
	t.Log(l2i)
}
