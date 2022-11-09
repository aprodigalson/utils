package ipaddr

import (
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetLastIpInCidr(t *testing.T) {
	Convey("Get cidr last ip ipv 4", t, func() {
		ip := GetLastIPInCidr(getCidrByString("1.1.1.0/24"))
		So("1.1.1.255", ShouldEqual, ip.String())
		
	})
}

func getCidrByString(cidrStr string)net.IPNet{
	_, ipNet, _:=net.ParseCIDR(cidrStr)
	return *ipNet
}