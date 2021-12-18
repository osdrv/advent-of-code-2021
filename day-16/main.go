package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	HEX2BIN = map[byte]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
)

func binToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[len(s)-1-i] == '1' {
			res += (1 << i)
		}
	}
	return res
}

func hexToBin(s string) string {
	var buf bytes.Buffer
	for _, ch := range s {
		buf.WriteString(HEX2BIN[byte(ch)])
	}
	return buf.String()
}

func readInt(s string, off, ll int) (int, int) {
	return binToInt(s[off : off+ll]), off + ll
}

func readLiteral(s string, off int) (int, int) {
	oct := make([]string, 0, 1)
	ix := off
	for {
		readmore := s[ix] == '1'
		ix += 1
		oct = append(oct, s[ix:ix+4])
		ix += 4
		if !readmore {
			break
		}
	}
	return binToInt(strings.Join(oct, "")), ix
}

type packet struct {
	ver    int
	typeId int
	val    int
	subs   []packet
}

func (p packet) String() string {
	s := fmt.Sprintf("p{ver: %d, typeId: %d", p.ver, p.typeId)
	if p.subs != nil {
		subs := make([]string, 0, 1)
		for _, sub := range p.subs {
			subs = append(subs, sub.String())
		}
		s += fmt.Sprintf(", subs=[%s]", strings.Join(subs, ", "))
	} else {
		s += fmt.Sprintf(", val=%d", p.val)
	}
	s += "}"
	return s
}

func parsePacket(binstr string) (packet, int) {
	p := packet{}
	ix := 0
	var version int
	version, ix = readInt(binstr, ix, 3)
	p.ver = version

	var typeId int
	typeId, ix = readInt(binstr, ix, 3)
	p.typeId = typeId

	switch typeId {
	case 4:
		var literal int
		literal, ix = readLiteral(binstr, ix)
		p.val = literal
		//packet = append(packet, literal)
	default:
		// opertor
		var lenTypeId int
		lenTypeId, ix = readInt(binstr, ix, 1)
		if lenTypeId == 0 {
			// total length in bits of the sub-packets contained by this packet
			var bitlen int
			bitlen, ix = readInt(binstr, ix, 15)
			until := ix + bitlen
			for ix < until {
				subpacket, off := parsePacket(binstr[ix:])
				p.subs = append(p.subs, subpacket)
				//packet = append(packet, subpacket...)
				ix += off
			}
		} else {
			var numsubs int
			// number of sub-packets immediately contained by this packet
			numsubs, ix = readInt(binstr, ix, 11)
			for i := 0; i < numsubs; i++ {
				subpacket, off := parsePacket(binstr[ix:])
				//packet = append(packet, subpacket...)
				p.subs = append(p.subs, subpacket)
				ix += off
			}
		}
	}

	return p, ix
}

func sumVersions(p packet) int {
	v := p.ver
	if p.subs != nil {
		for _, sub := range p.subs {
			v += sumVersions(sub)
		}
	}
	return v
}

func computeVal(p packet) int {
	switch p.typeId {
	case 0:
		// sum
		res := 0
		for _, sub := range p.subs {
			res += computeVal(sub)
		}
		return res
	case 1:
		res := 1
		for _, sub := range p.subs {
			res *= computeVal(sub)
		}
		return res
	case 2:
		res := ALOT
		for _, sub := range p.subs {
			if v := computeVal(sub); v < res {
				res = v
			}
		}
		return res
	case 3:
		res := -1
		for _, sub := range p.subs {
			if v := computeVal(sub); v > res {
				res = v
			}
		}
		return res
	case 4:
		return p.val
	case 5:
		v1, v2 := computeVal(p.subs[0]), computeVal(p.subs[1])
		if v1 > v2 {
			return 1
		}
		return 0
	case 6:
		v1, v2 := computeVal(p.subs[0]), computeVal(p.subs[1])
		if v1 < v2 {
			return 1
		}
		return 0
	case 7:
		v1, v2 := computeVal(p.subs[0]), computeVal(p.subs[1])
		if v1 == v2 {
			return 1
		}
		return 0
	}
	return -1
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	hexstr := readFile(f)
	binstr := hexToBin(hexstr)

	printf("binary string: %s", binstr)

	p, _ := parsePacket(binstr)
	printf("packet: %s", p)

	printf("sum versions: %d", sumVersions(p))

	printf("value is: %d", computeVal(p))
}
