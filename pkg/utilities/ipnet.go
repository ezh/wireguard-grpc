package utilities

import (
	"fmt"
	"net"
	"strings"
)

func StringToIPNetSlice(val string) ([]net.IPNet, error) {
	val = strings.Trim(val, "[]")
	// Empty string would cause a slice with one (empty) entry
	if len(val) == 0 {
		return []net.IPNet{}, nil
	}
	ss := strings.Split(val, ",")
	out := make([]net.IPNet, len(ss))
	for i, sval := range ss {
		_, n, err := net.ParseCIDR(strings.TrimSpace(sval))
		if err != nil {
			return nil, fmt.Errorf("invalid string being converted to CIDR: %s", sval)
		}
		out[i] = *n
	}
	return out, nil
}

func IPNetSliceToString(s []net.IPNet) []string {
	ipNetStrSlice := make([]string, len(s))
	for i, n := range s {
		ipNetStrSlice[i] = n.String()
	}
	return ipNetStrSlice
}
