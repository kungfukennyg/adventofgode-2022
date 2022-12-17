package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// err := exec(input1)
	// if err != nil {
	// 	panic(err)
	// }

	// err := exec(input2)
	// if err != nil {
	// 	panic(err)
	// }

	err := exec(input3)
	if err != nil {
		panic(err)
	}
}

func exec(in string) error {
	instr, err := instrsFromStr(in)
	if err != nil {
		return err
	}

	cpu := &cpu{1, 0, 0, 0, 0, ""}

	for _, i := range instr {
		cpu.run(i)
		// fmt.Printf("cpu: %+v\n", cpu.string())
		fmt.Println()
	}

	fmt.Println(cpu.draw)

	return nil
}

type cpu struct {
	regX           int
	pc             int
	signalStrength int
	x, y           int
	draw           string
}

func (c cpu) string() string {
	return fmt.Sprintf("cpu:{regX: %d, pc: %d, drawX: %d, drawY: %d}", c.regX, c.pc, c.x, c.y)
}

func (c *cpu) run(op instr) {
	// fmt.Println(op.string())
	for i := 1; i <= op.cycles; i++ {
		c.pc += 1
		fmt.Printf("Start cycle   %d: begin executing %s\n", c.pc, op.string())
		fmt.Printf("During cycle %d: CRT draws pixel in position: (%d, %d)\n", c.pc, c.x, c.y)
		fmt.Println()

		// draw
		if c.x > 0 && c.x%39 == 0 {
			c.x = 0
			c.y += 1
			c.draw += "\n"
		} else {
			c.x += 1
			if c.x%6 == 0 {
				c.draw += " "
			}
		}

		// check pos for drawing, drawing is 3 pixels wide
		// and regX represents the center of it (0)
		var match bool
		for i := range []int{-1, 0, 1} {
			pos := c.regX + i
			if pos == c.x {
				c.draw += "#"
				match = true
			}
		}
		if !match {
			c.draw += "."
		}
	}

	switch op.cmd {
	case instrNoOp:
		// noop
		break
	case instrAddX:
		c.regX += op.mod
	default:
		panic(fmt.Sprintf("unrecognized instruction %+v", op))
	}

	fmt.Printf("End of cycle  %d: finish executing %s (Register X is now %d)\n", c.pc, op.string(), c.regX)
}

func instrsFromStr(in string) ([]instr, error) {
	ret := []instr{}
	for _, line := range strings.Split(in, "\n") {
		parts := strings.Split(line, " ")
		var i instr
		switch parts[0] {
		case instrNoOp:
			i = instr{
				cmd:    instrNoOp,
				cycles: 1,
			}
		case instrAddX:
			i = instr{
				cmd:    instrAddX,
				cycles: 2,
			}
			n, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return ret, err
			}

			i.mod = int(n)
		default:
			panic(fmt.Sprintf("unrecognized instruction '%s'", parts[0]))
		}

		ret = append(ret, i)
	}
	return ret, nil
}

type instr struct {
	cmd    string
	mod    int
	cycles int
}

func (i instr) string() string {
	switch i.cmd {
	case instrNoOp:
		return i.cmd
	case instrAddX:
		return fmt.Sprintf("%s %d", i.cmd, i.mod)
	default:
		panic(fmt.Sprintf("unrecognized instruction '%s'", i.cmd))
	}
}

var (
	instrNoOp = "noop"
	instrAddX = "addx"
)

const input1 string = `noop
addx 3
addx -5`

const input2 string = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

const input3 string = `noop
noop
noop
addx 3
addx 7
noop
noop
noop
noop
addx 6
noop
addx -1
noop
addx 5
addx 1
noop
addx 4
noop
noop
noop
noop
addx 6
addx -1
noop
addx 3
addx -13
addx -22
noop
noop
addx 3
addx 2
addx 11
addx -4
addx 11
addx -10
addx 2
addx 5
addx 2
addx -2
noop
addx 7
addx 3
addx -2
addx 2
addx 5
addx 2
addx -2
addx -8
addx -27
addx 5
addx 2
addx 21
addx -21
addx 3
addx 5
addx 2
addx -3
addx 4
addx 3
addx 1
addx 5
noop
noop
noop
noop
addx 3
addx 1
addx 6
addx -31
noop
addx -4
noop
noop
noop
noop
addx 3
addx 7
noop
addx -1
addx 1
addx 5
noop
addx 1
noop
addx 2
addx -8
addx 15
addx 3
noop
addx 2
addx 5
noop
noop
noop
addx -28
addx 11
addx -20
noop
addx 7
addx -2
addx 7
noop
addx -2
noop
addx -6
addx 11
noop
addx 3
addx 2
noop
noop
addx 7
addx 3
addx -2
addx 2
addx 5
addx 2
addx -16
addx -10
addx -11
addx 27
addx -20
noop
addx 2
addx 3
addx 5
noop
noop
noop
addx 3
addx -2
addx 2
noop
addx -14
addx 21
noop
addx -6
addx 12
noop
addx -21
addx 24
addx 2
noop
noop
noop`
