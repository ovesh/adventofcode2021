package day16

import (
	"encoding/hex"
	"fmt"
)

const packetTypeSum = 0
const packetTypeProduct = 1
const packetTypeMin = 2
const packetTypeMax = 3
const packetTypeGT = 5
const packetTypeLT = 6
const packetTypeEQ = 7

const maxInt = int(^uint(0) >> 1)

func (p packet) value() int {
	res := 0
	switch p.typ {
	case packetTypeLiteral:
		return p.literalVal
	case packetTypeSum:
		for _, sb := range p.subPackets {
			res += sb.value()
		}
		return res
	case packetTypeProduct:
		res := 1
		for _, sb := range p.subPackets {
			res *= sb.value()
		}
		return res
	case packetTypeMin:
		res := maxInt
		for _, sb := range p.subPackets {
			curVal := sb.value()
			if curVal < res {
				res = curVal
			}
		}
		return res
	case packetTypeMax:
		for _, sb := range p.subPackets {
			curVal := sb.value()
			if curVal > res {
				res = curVal
			}
		}
		return res
	case packetTypeGT:
		if len(p.subPackets) != 2 {
			panic(fmt.Sprintf("expected 2 subpackets but saw %d", len(p.subPackets)))
		}
		if p.subPackets[0].value() > p.subPackets[1].value() {
			return 1
		}
		return 0
	case packetTypeLT:
		if len(p.subPackets) != 2 {
			panic(fmt.Sprintf("expected 2 subpackets but saw %d", len(p.subPackets)))
		}
		if p.subPackets[0].value() < p.subPackets[1].value() {
			return 1
		}
		return 0
	case packetTypeEQ:
		if len(p.subPackets) != 2 {
			panic(fmt.Sprintf("expected 2 subpackets but saw %d", len(p.subPackets)))
		}
		if p.subPackets[0].value() == p.subPackets[1].value() {
			return 1
		}
		return 0
	}
	panic(fmt.Sprintf("unknown type %d", p.typ))
}

func Part2() int {
	bDecoded, _ := hex.DecodeString(input)
	decoded := bitSlice{bytes: bDecoded, offset: 0, len: 8 * len(bDecoded)}
	p := parsePacket(decoded)
	return p.value()
}
