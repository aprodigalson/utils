package ipaddr

import (
	"encoding/binary"
	"net"
)

// GetLastIPInCidr 获取cidr对应网段的最后一个地址
//  @param cidr
//  @return net.IP
func GetLastIPInCidr(cidr net.IPNet) net.IP {
	switch len(cidr.IP) {
	case net.IPv4len:
		return getLastIPInCidrV4(cidr)
	case net.IPv6len:
	default:
		return nil
	}
	return nil

}

func getLastIPInCidrV4(cidr net.IPNet) net.IP {
	mask := binary.BigEndian.Uint32(cidr.Mask)
	ipInt32 := binary.BigEndian.Uint32(cidr.IP)

	lastIPInt := (ipInt32 & mask) | (mask ^ 0xffffffff)
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, lastIPInt)
	return ip
}
