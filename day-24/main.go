package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InstrCode string

const (
	INP InstrCode = "inp"
	ADD           = "add"
	MUL           = "mul"
	DIV           = "div"
	MOD           = "mod"
	EQL           = "eql"
)

type Operand struct {
	val int64
	ref byte
}

func (o *Operand) String() string {
	if o.ref > 0 {
		return string([]byte{o.ref})
	}
	return strconv.Itoa(int(o.val))
}

type Instr struct {
	code     InstrCode
	op1, op2 *Operand
}

func NewInstr(s string) *Instr {
	instr := &Instr{
		code: InstrCode(s[0:3]),
		op1:  parseOp(s[4:5]),
	}
	if len(s) > 5 {
		instr.op2 = parseOp(s[6:])
	}
	return instr
}

func (i *Instr) String() string {
	res := fmt.Sprintf("(%s %s", i.code, i.op1)
	if i.op2 != nil {
		res += " " + i.op2.String()
	}
	res += ")"
	return res
}

func (i *Instr) Exec(input func() int, regs map[byte]int64) {
	switch i.code {
	case INP:
		regs[i.op1.ref] = int64(input())
	case ADD:
		var v2 int64
		if i.op2.ref > 0 {
			v2 = regs[i.op2.ref]
		} else {
			v2 = i.op2.val
		}
		regs[i.op1.ref] += v2
	case MUL:
		var v2 int64
		if i.op2.ref > 0 {
			v2 = regs[i.op2.ref]
		} else {
			v2 = i.op2.val
		}
		regs[i.op1.ref] *= v2
	case DIV:
		var v2 int64
		if i.op2.ref > 0 {
			v2 = regs[i.op2.ref]
		} else {
			v2 = i.op2.val
		}
		regs[i.op1.ref] /= v2
	case MOD:
		var v2 int64
		if i.op2.ref > 0 {
			v2 = regs[i.op2.ref]
		} else {
			v2 = i.op2.val
		}
		regs[i.op1.ref] %= v2
	case EQL:
		var v2 int64
		if i.op2.ref > 0 {
			v2 = regs[i.op2.ref]
		} else {
			v2 = i.op2.val
		}
		if regs[i.op1.ref] == v2 {
			regs[i.op1.ref] = 1
		} else {
			regs[i.op1.ref] = 0
		}
	default:
		panic(fmt.Sprintf("unknown operation %s", i.code))
	}
}

func compute(input []int, regs map[byte]int64, instrs []*Instr) {
	ix := 0
	in := func() int {
		v := input[ix]
		ix++
		return v
	}
	for _, instr := range instrs {
		instr.Exec(in, regs)
	}
}

func parseOp(s string) *Operand {
	if s[0] >= 'a' && s[0] <= 'z' {
		return &Operand{
			ref: s[0],
		}
	}
	return &Operand{
		val: int64(parseInt(s)),
	}
}

const (
	DIGITS = 10
)

// output: w, x, y, z
func hardcoded(in [14]int) (int64, int64, int64, int64) {
	var w, x, y, z int64

	// digit 0
	w = int64(in[0])
	x = z%26 + 12
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 9
	y *= x
	z += y

	// digit 1
	w = int64(in[1])
	x = z%26 + 12
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 4
	y *= x
	z += y

	// digit 2
	w = int64(in[2])
	x = z%26 + 12
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 2
	y *= x
	z += y

	// digit 3
	w = int64(in[3])
	x = z%26 - 9
	z /= 26
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 5
	y *= x
	z += y

	// digit 4
	w = int64(in[4])
	x = z%26 - 9
	z /= 26
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 1
	y *= x
	z += y

	// digit 5
	w = int64(in[5])
	x = z%26 + 14
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 6
	y *= x
	z += y

	// digit 6
	w = int64(in[6])
	x = z%26 + 14
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 11
	y *= x
	z += y

	// digit 7
	w = int64(in[7])
	x = z%26 - 10
	z /= 26
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 15
	y *= x
	z += y

	// digit 8
	w = int64(in[8])
	x = z%26 + 15
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 7
	y *= x
	z += y

	// digit 9
	w = int64(in[9])
	x = z%26 - 2
	z /= 26
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 12
	y *= x
	z += y

	// digit 10
	w = int64(in[10])
	x = z%26 + 11
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 15
	y *= x
	z += y

	// digit 11
	w = int64(in[11])
	x = z%26 - 15
	z /= 26
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 9
	y *= x
	z += y

	//// digit 12
	w = int64(in[12])
	x = z%26 - 9
	z /= 26
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 12
	y *= x
	z += y

	//// digit 13
	w = int64(in[13])
	x = z%26 - 3
	z /= 26
	if x != w {
		x = 1
	} else {
		x = 0
	}
	y = 25*x + 1
	z *= y
	y = w + 12
	y *= x
	z += y

	return w, x, y, z
}

func solve1(instrs []*Instr) {
	//num := [14]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//num := [14]int{1, 1, 5, 1, 4, 1, 2, 1, 6, 1, 1, 1, 2, 1}
	//num := [14]int{2, 1, 7, 3, 2, 3, 8, 2, 6, 7, 7, 6, 4, 1}
	//num := [14]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9}
	//num := [14]int{1, 1, 1, 5, 1, 4, 1, 2, 1, 6, 1, 1, 1, 2}
	//num := [14]int{1, 6, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//num := [14]int{1, 6, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 7}
	//num := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 7}
	num := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	//runtime.Breakpoint()
	w, x, y, z := hardcoded(num)
	printf("%d %d %d %d", w, x, y, z)
	//num := [14]byte{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7}
	minz := ALOT64
Iter:
	for {
		//printf("inspecting number %v", num)
		//regs := make(map[byte]int64)
		//compute(bytes.NewReader(num[:]), regs, instrs)
		//if regs['z'] == 0 {
		//	printf("valid number %v", num)
		//} else {
		//	if regs['z'] < minz {
		//		minz = regs['z']
		//		printf("new min z: %d for number %v", minz, num)
		//	}
		//	//printf("reg z: %d", regs['z'])
		//}
		//w, x, y, z := hardcoded(num)
		_, _, _, z := hardcoded(num)
		//printf("HRC: %d %d %d %d", w, x, y, z)
		//regs := make(map[byte]int64)
		//compute(num[:], regs, instrs)
		//printf("ALU: %d %d %d %d", regs['w'], regs['x'], regs['y'], regs['z'])
		if z == 0 {
			printf("valid number %v", num)
		} else {
			if z < minz {
				minz = z
				printf("new min z: %d for number %v", minz, num)
			}
		}
		//for i := 0; i < len(num); i++ {
		//	num[i]++
		//	if num[i] > 9 {
		//		break Iter
		//	}
		//}
		carry := 1 // increment by 1
		//mini := 3
		mini := 4
		for i := len(num) - 1; i >= mini; i-- {
			//for i := DIGITS - 1; i >= mini; i-- {
			num[i] -= carry
			if num[i] == 0 {
				if i == mini {
					break Iter
				}
				num[i] += 9
				carry = 1
				continue
			}
			break
		}
	}
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	instrs := make([]*Instr, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "//") {
			// comment
			continue
		}
		instrs = append(instrs, NewInstr(line))
	}
	printf("instructions: %+v", instrs)

	//regs := make(map[byte]int64)
	//compute([]int{1, 3}, regs, instrs)
	//printf("regs: %+v", regs)

	solve1(instrs)

}

/*

       0  1  2  3  4  5  6   7  8  9 10  11 12 13
p1     1  1  1 26 26  1  1  26  1 26  1  26 26 26
p2    12 12 12 -9 -9 14 14 -10 15 -2 11 -15 -9 -3
p3     9  4  2  5  1  6 11  15  7 12 15   9 12 12

p2, p3
inp3 = inp2 + 2 - 9 = inp2 - 7
p1, p4
inp4 = inp1 + 4 - 9 = inp1 - 5
p6, p7
inp7 = inp6 + 11 - 10 = inp6 + 1
p8, p9
inp9 = inp8 + 7 - 2 = inp8 + 5
p10, p11
inp11 = inp10 + 15 - 15 = inp10
p5, p12
inp12 = inp5 + 6 - 9 = inp5 - 3
p0, p13
inp13 = inp0 + 9 - 3 = inp0 + 6

00000000001111
01234567890123
39924989499969

00000000001111
01234567890123
16811412161117
*/
