package day16

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
)

func byte2BinString(b byte) string {
	res := ""
	for i := 0; i < 8; i++ {
		res = strconv.Itoa((int(b>>i))&1) + res
	}
	return res
}

type packet struct {
	version    int
	typ        int
	lenTypeID  int
	literalVal int
	bitLen     int
	subPackets []packet
}

func (p packet) totalVersions() int {
	res := p.version
	for _, sp := range p.subPackets {
		res += sp.totalVersions()
	}
	return res
}

type bitSlice struct {
	bytes  []byte
	offset int
	len    int
}

func (bs bitSlice) subSlice(start, length int) bitSlice {
	shiftRight := (8 - (bs.offset+start+length)%8) % 8
	byteNum := length / 8
	if length%8 != 0 {
		byteNum++
	}
	res := make([]byte, len(bs.bytes))
	copy(res, bs.bytes)
	for i := len(bs.bytes) - 1; i > 0; i-- {
		res[i] = res[i]>>shiftRight | res[i-1]<<(8-shiftRight)
	}
	res[0] >>= shiftRight

	firstByteIdx := (bs.offset + start + shiftRight) / 8
	lastByteIdx := (bs.offset + start + length + shiftRight) / 8
	if (bs.offset+start+length+shiftRight)%8 != 0 {
		lastByteIdx++
	}
	if lastByteIdx >= len(res) {
		res = res[firstByteIdx:]
	} else {
		res = res[firstByteIdx:lastByteIdx]
	}

	offset := (8 - (length % 8)) % 8
	res[0] &= 255 >> offset

	return bitSlice{bytes: res, offset: offset, len: length}
}

func (bs bitSlice) append(bs2 bitSlice) bitSlice {
	res := make([]byte, len(bs.bytes)+len(bs2.bytes))
	for i := len(bs2.bytes) - 1; i >= 0; i-- {
		res[i+len(bs.bytes)] = bs2.bytes[i]
	}
	for i := len(bs.bytes) - 1; i >= 0; i-- {
		res[i] = bs.bytes[i]
	}
	shiftRight := bs2.offset
	i := len(bs.bytes)
	res[i] = res[i] | res[i-1]<<(8-shiftRight)
	for i := len(bs.bytes) - 1; i > 0; i-- {
		res[i] = res[i]>>shiftRight | res[i-1]<<(8-shiftRight)
	}
	res[0] >>= shiftRight
	for res[0] == 0 {
		res = res[1:]
	}
	return bitSlice{bytes: res, offset: bs.offset + shiftRight, len: bs.len + bs2.len}
}

func parsePacket(bs bitSlice) packet {
	packetType := int((binary.BigEndian.Uint16(bs.bytes[:2]) >> (10 - bs.offset)) & 7)
	version := int((binary.BigEndian.Uint16(bs.bytes[:2]) >> (13 - bs.offset)) & 7)
	if packetType == packetTypeLiteral {
		packetBitLen := 6
		packetRest := bs.subSlice(6, bs.len-6)
		readOn := 1
		assembledVal := bitSlice{}
		for readOn != 0 {
			next4 := packetRest.subSlice(0, 5)
			packetBitLen += 5
			if assembledVal.bytes == nil {
				assembledVal = next4.subSlice(1, 4)
			} else {
				assembledVal = assembledVal.append(next4.subSlice(1, 4))
			}
			readOn = int(next4.bytes[0] >> 4)
			if readOn == 0 {
				return packet{version: version, typ: packetType, literalVal: assembledVal.asInt(), bitLen: packetBitLen}
			}
			packetRest = packetRest.subSlice(5, packetRest.len-5)
		}
	}

	// has subpackets
	lenTypeID := int((binary.BigEndian.Uint16(bs.bytes[:2]) >> (9 - bs.offset)) & 1)
	if lenTypeID == 0 {
		lenFieldVal := bs.subSlice(7, 15).asInt()
		offset := 22 // version (3) + type (3) + lentype (1) + lenVal (15)
		res := packet{version: version, typ: packetType, bitLen: offset + lenFieldVal}
		remainingBits := lenFieldVal
		for remainingBits > 0 {
			packetRest := bs.subSlice(offset, remainingBits)
			nextPacket := parsePacket(packetRest)
			res.subPackets = append(res.subPackets, nextPacket)
			remainingBits -= nextPacket.bitLen
			offset += nextPacket.bitLen
		}
		return res
	}

	// lenTypeID == 1
	lenFieldVal := bs.subSlice(7, 11).asInt()
	offset := 18 // version (3) + type (3) + lentype (1) + lenVal (11)
	res := packet{version: version, typ: packetType, bitLen: offset}
	remainingBits := bs.len - offset
	for remainingPackets := lenFieldVal; remainingPackets > 0; remainingPackets-- {
		packetRest := bs.subSlice(offset, remainingBits)
		nextPacket := parsePacket(packetRest)
		remainingBits -= nextPacket.bitLen
		offset += nextPacket.bitLen
		res.bitLen += nextPacket.bitLen
		res.subPackets = append(res.subPackets, nextPacket)
	}
	return res
}

func (bs bitSlice) String() string {
	res := ""
	for i := 0; i < len(bs.bytes); i++ {
		b := bs.bytes[i]
		res += byte2BinString(b)
	}
	return res[bs.offset:]
}

func (bs bitSlice) asInt() int {
	res := int(bs.bytes[0])
	for i := 1; i < len(bs.bytes); i++ {
		res = 256*res + int(bs.bytes[i])
	}
	return res
}

const packetTypeLiteral = 4

func Part1() int {
	bDecoded, _ := hex.DecodeString(input)
	decoded := bitSlice{bytes: bDecoded, offset: 0, len: 8 * len(bDecoded)}
	p := parsePacket(decoded)

	return p.totalVersions()
}
