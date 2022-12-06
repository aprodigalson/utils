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
		return getLastIPInCidrV6(cidr)
	default:
		return nil
	}
}

func getLastIPInCidrV4(cidr net.IPNet) net.IP {
	mask := binary.BigEndian.Uint32(cidr.Mask)
	ipInt32 := binary.BigEndian.Uint32(cidr.IP)

	lastIPInt := (ipInt32 & mask) | (mask ^ 0xffffffff)
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, lastIPInt)
	return ip
}

func getLastIPInCidrV6(cidr net.IPNet) net.IP {

	maskLow := binary.BigEndian.Uint64(cidr.Mask[:8])
	ipIntLow := binary.BigEndian.Uint64(cidr.IP[:8])

	maskHigh := binary.BigEndian.Uint64(cidr.Mask[8:])
	ipIntHigh := binary.BigEndian.Uint64(cidr.IP[8:])

	ipIntLow = (ipIntLow & maskLow) | (maskLow ^ 0xffffffffffffffff)
	ipIntHigh = (ipIntHigh & maskHigh) | (maskHigh ^ 0xffffffffffffffff)
	ip := make(net.IP, 16)
	binary.BigEndian.PutUint64(ip[:8], ipIntLow)
	binary.BigEndian.PutUint64(ip[8:], ipIntHigh)
	return ip
}
