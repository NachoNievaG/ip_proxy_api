package country

import (
	"encoding/binary"
	"errors"
	"net"
)

type IP interface {
	ipToInt(ip net.IP) uint32
	intToIP(nn uint32) net.IP
}

type IPAddress struct {
	IP    net.IP `json:"ip"`
	IntIP uint32 `json:"-"`
}

var ErrInvalidIP = errors.New("invalid ip")

func NewIPAddressFromString(ip string) (IPAddress, error) {
	var res IPAddress

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return res, ErrInvalidIP
	}
	res.IP = parsedIP
	res.IntIP = res.ipToInt(parsedIP)

	return res, nil
}

func NewIPAddressFromBigEndian(ip uint32) IPAddress {
	var res IPAddress

	res.IP = res.intToIP(ip)
	res.IntIP = res.ipToInt(res.IP)

	return res
}

func (addr IPAddress) ipToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func (addr IPAddress) intToIP(val uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, val)
	return ip
}
