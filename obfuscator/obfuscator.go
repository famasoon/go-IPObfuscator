package obfuscator

import (
	"encoding/binary"
	"fmt"
	"net"
)

func IPtoDecimal(ip net.IP) string {
	return fmt.Sprintf("%d", binary.BigEndian.Uint32(ip))
}

func IPtoHex(ip net.IP) string {
	return fmt.Sprintf("0x%02x.0x%02x.0x%02x.0x%02x", ip[0], ip[1], ip[2], ip[3])
}

func IPtoOct(ip net.IP) string {
	return fmt.Sprintf("0%o.0%o.0%o.0%o", ip[0], ip[1], ip[2], ip[3])
}

func IPtoHexPadding(ip net.IP) string {
	return fmt.Sprintf("0x%010x.0x%010x.0x%010x.0x%010x", ip[0], ip[1], ip[2], ip[3])
}

func IPtoOctPadding(ip net.IP) string {
	return fmt.Sprintf("0%010o.0%010o.0%010o.0%010o", ip[0], ip[1], ip[2], ip[3])
}

func IPtoHexAndDec(ip net.IP) string {
	return fmt.Sprintf("0x%02x.0x%02x.0x%02x.%d", ip[0], ip[1], ip[2], ip[3])
}

func IPtoHexAndDec2(ip net.IP) string {
	return fmt.Sprintf("0x%02x.0x%02x.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func IPtoHexAndDec3(ip net.IP) string {
	return fmt.Sprintf("0x%02x.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func IPtoOctAndDec(ip net.IP) string {
	return fmt.Sprintf("0%o.0%o.0%o.%d", ip[0], ip[1], ip[2], ip[3])
}

func IPtoOctAndDec2(ip net.IP) string {
	return fmt.Sprintf("0%o.0%o.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func IPtoOctAndDec3(ip net.IP) string {
	return fmt.Sprintf("0%o.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func HexAndDecimal(ip net.IP) string {
	low := binary.BigEndian.Uint32(ip) & 0x0000ffff
	return fmt.Sprintf("0x%02x.0x%02x.%d", ip[0], ip[1], low)
}

func OctAndDecimal(ip net.IP) string {
	low := binary.BigEndian.Uint32(ip) & 0x0000ffff
	return fmt.Sprintf("0%o.0%o.%d", ip[0], ip[1], low)
}

func HexOctDec(ip net.IP) string {
	low := binary.BigEndian.Uint32(ip) & 0x0000ffff
	return fmt.Sprintf("0x%x.0%o.%d", ip[0], ip[1], low)
}

func HexAndDec(ip net.IP) string {
	low := binary.BigEndian.Uint32(ip) & 0x00ffffff
	return fmt.Sprintf("0x%x.%d", ip[0], low)
}

func OctAndDec(ip net.IP) string {
	low := binary.BigEndian.Uint32(ip) & 0x00ffffff
	return fmt.Sprintf("0%o.%d", ip[0], low)
}

func Hex2Oct2(ip net.IP) string {
	return fmt.Sprintf("0x%x.0x%x.0%o.0%o", ip[0], ip[1], ip[2], ip[3])
}

func HexOct3(ip net.IP) string {
	return fmt.Sprintf("0x%x.0%o.0%o.0%o", ip[0], ip[1], ip[2], ip[3])
}
