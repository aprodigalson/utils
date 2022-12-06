package ipaddr

import (
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetLastIpInCidr(t *testing.T) {
	Convey("Get cidr last ip ", t, func() {
		Convey("Get cidr last ip ipv 4", func() {
			ip := GetLastIPInCidr(getCidrByString("1.1.1.0/24"))
			So("1.1.1.255", ShouldEqual, ip.String())
			ip = GetLastIPInCidr(getCidrByString("1.1.1.1/32"))
			So("1.1.1.1", ShouldEqual, ip.String())
			ip = GetLastIPInCidr(getCidrByString("1.1.1.0/0"))
			So("255.255.255.255", ShouldEqual, ip.String())

		})
		Convey("Get cidr last ip ipv 6", func() {
			ip := GetLastIPInCidr(getCidrByString("1111::ffff/128"))
			So("1111::ffff", ShouldEqual, ip.String())

			ip = GetLastIPInCidr(getCidrByString("1111::/127"))
			So("1111::1", ShouldEqual, ip.String())

			ip = GetLastIPInCidr(getCidrByString("1111::/126"))
			So("1111::3", ShouldEqual, ip.String())

			ip = GetLastIPInCidr(getCidrByString("A111::/96"))
			So("a111::ffff:ffff", ShouldEqual, ip.String())

		})
	})
}

func getCidrByString(cidrStr string) net.IPNet {
	_, ipNet, _ := net.ParseCIDR(cidrStr)
	return *ipNet
}
