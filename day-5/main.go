package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []string

type instr struct {
	numCrates int
	from      int
	to        int
}

func (s *stack) pop() string {
	if len(*s) == 0 {
		return ""
	}

	n := len(*s) - 1
	ret := (*s)[n]
	*s = (*s)[:n]
	return ret
}

func (s *stack) push(str string) {
	*s = append(*s, str)
}

const testInput string = `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func main() {
	stacks := pt1Stacks()
	printStacks(stacks)
	instructions := parseInstrs(pt1Input)

	for _, op := range instructions {
		fmt.Println(op)
		from := (*stacks)[op.from]
		to := (*stacks)[op.to]

		pt2Move(&from, &to, op.numCrates)

		(*stacks)[op.from] = from
		(*stacks)[op.to] = to
		printStacks(stacks)
	}

	fmt.Println("Ending result...")
	printStacks(stacks)
}

func pt1Move(from, to *stack, numCrates int) {
	for i := 0; i < numCrates; i++ {
		crate := from.pop()
		to.push(crate)
	}
}

func pt2Move(from, to *stack, numCrates int) {
	toMove := []string{}
	for i := 0; i < numCrates; i++ {
		crate := from.pop()
		toMove = append(toMove, crate)
	}
	toMove = reverse(toMove)
	for _, m := range toMove {
		to.push(m)
	}
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (i instr) String() string {
	return fmt.Sprintf("move %d crates from crate %d to %d", i.numCrates, (i.from + 1), (i.to + 1))
}

func parseInstrs(input string) []instr {
	ret := []instr{}
	for _, in := range strings.Split(input, "\n") {
		op := parseInstr(in)
		ret = append(ret, op)
	}
	return ret
}

func parseInstr(line string) instr {
	// move 1 from 2 to 1
	line = strings.TrimPrefix(line, "move ")
	// 1 from 2 to 1
	parts := strings.Split(line, " ")
	if len(parts) != 5 {
		panic("weird number of parts")
	}

	return instr{
		numCrates: parse(parts[0]),
		from:      parse(parts[2]) - 1,
		to:        parse(parts[4]) - 1,
	}
}

func parse(in string) int {
	num, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}

func printStacks(stacks *[]stack) {
	tallest := 0
	for _, s := range *stacks {
		if len(s) > tallest {
			tallest = len(s)
		}
	}
	fmt.Println()

	for i := tallest; i > 0; i-- {
		for _, s := range *stacks {
			if len(s) < i {
				fmt.Printf("\t")
				continue
			}

			fmt.Printf("[%s]\t", s[i-1])
		}
		fmt.Println()
	}

	for i := range *stacks {
		fmt.Printf("%d\t", (i + 1))
	}
	fmt.Println()
}

func (s stack) String() string {
	ret := ""
	for _, crate := range s {
		ret += fmt.Sprintf("[%s] ", crate)
	}
	return ret
}

func testStacks() *[]stack {
	stacks := []stack{}
	stacks = append(stacks, stack{"Z", "N"})
	stacks = append(stacks, stack{"M", "C", "D"})
	stacks = append(stacks, stack{"P"})
	return &stacks
}

func pt1Stacks() *[]stack {
	stacks := []stack{}
	stacks = append(stacks, stack{"Q", "S", "W", "C", "Z", "V", "F", "T"})
	stacks = append(stacks, stack{"Q", "R", "B"})
	stacks = append(stacks, stack{"B", "Z", "T", "Q", "P", "M", "S"})
	stacks = append(stacks, stack{"D", "V", "F", "R", "Q", "H"})
	stacks = append(stacks, stack{"J", "G", "L", "D", "B", "S", "T", "P"})
	stacks = append(stacks, stack{"W", "R", "T", "Z"})
	stacks = append(stacks, stack{"H", "Q", "M", "N", "S", "F", "R", "J"})
	stacks = append(stacks, stack{"R", "N", "F", "H", "W"})
	stacks = append(stacks, stack{"J", "Z", "T", "Q", "P", "R", "B"})
	return &stacks
}

const pt1Input string = `move 3 from 8 to 2
move 3 from 1 to 5
move 3 from 1 to 4
move 2 from 7 to 4
move 3 from 7 to 4
move 8 from 5 to 7
move 2 from 1 to 8
move 7 from 3 to 2
move 1 from 5 to 2
move 1 from 6 to 7
move 2 from 5 to 9
move 1 from 9 to 1
move 3 from 9 to 6
move 5 from 6 to 2
move 10 from 7 to 2
move 3 from 8 to 9
move 7 from 9 to 2
move 1 from 1 to 2
move 1 from 9 to 6
move 1 from 4 to 1
move 1 from 8 to 2
move 11 from 4 to 2
move 1 from 7 to 9
move 1 from 4 to 6
move 1 from 9 to 7
move 1 from 1 to 3
move 1 from 7 to 5
move 1 from 4 to 9
move 1 from 5 to 2
move 1 from 3 to 8
move 1 from 6 to 9
move 1 from 8 to 6
move 11 from 2 to 1
move 1 from 6 to 8
move 7 from 2 to 1
move 14 from 2 to 7
move 1 from 6 to 3
move 1 from 8 to 2
move 1 from 3 to 9
move 7 from 7 to 1
move 1 from 6 to 5
move 5 from 7 to 6
move 4 from 2 to 8
move 3 from 6 to 7
move 3 from 7 to 8
move 9 from 1 to 3
move 8 from 3 to 7
move 1 from 3 to 1
move 2 from 2 to 3
move 1 from 6 to 7
move 2 from 1 to 7
move 7 from 1 to 6
move 1 from 3 to 5
move 2 from 5 to 3
move 7 from 6 to 3
move 9 from 7 to 5
move 1 from 9 to 1
move 4 from 8 to 5
move 7 from 1 to 5
move 4 from 7 to 2
move 1 from 7 to 8
move 1 from 6 to 4
move 10 from 5 to 3
move 8 from 5 to 1
move 2 from 8 to 3
move 2 from 8 to 9
move 8 from 2 to 7
move 4 from 9 to 8
move 13 from 3 to 7
move 1 from 5 to 3
move 6 from 3 to 9
move 10 from 1 to 9
move 1 from 3 to 4
move 6 from 9 to 7
move 1 from 5 to 8
move 14 from 7 to 6
move 14 from 6 to 1
move 13 from 1 to 8
move 1 from 1 to 2
move 9 from 8 to 9
move 6 from 8 to 5
move 2 from 4 to 6
move 1 from 8 to 1
move 2 from 2 to 1
move 2 from 8 to 6
move 3 from 1 to 2
move 3 from 3 to 9
move 16 from 9 to 1
move 3 from 2 to 4
move 3 from 7 to 2
move 6 from 5 to 4
move 5 from 7 to 3
move 4 from 6 to 1
move 10 from 2 to 9
move 13 from 9 to 1
move 5 from 7 to 2
move 2 from 4 to 6
move 1 from 9 to 1
move 2 from 9 to 5
move 2 from 6 to 8
move 2 from 5 to 3
move 1 from 8 to 3
move 31 from 1 to 7
move 2 from 1 to 5
move 12 from 7 to 3
move 11 from 3 to 2
move 1 from 8 to 4
move 6 from 4 to 5
move 1 from 3 to 4
move 8 from 3 to 2
move 5 from 5 to 6
move 2 from 6 to 7
move 4 from 7 to 3
move 1 from 6 to 9
move 13 from 7 to 6
move 13 from 2 to 3
move 1 from 4 to 8
move 10 from 2 to 3
move 3 from 7 to 3
move 2 from 2 to 1
move 1 from 8 to 2
move 2 from 4 to 7
move 1 from 9 to 2
move 3 from 7 to 3
move 1 from 5 to 1
move 2 from 5 to 2
move 15 from 6 to 7
move 4 from 1 to 9
move 22 from 3 to 9
move 7 from 3 to 9
move 4 from 3 to 8
move 4 from 9 to 4
move 3 from 2 to 4
move 5 from 7 to 1
move 7 from 4 to 7
move 2 from 8 to 4
move 1 from 4 to 8
move 3 from 1 to 5
move 2 from 1 to 4
move 1 from 2 to 9
move 2 from 5 to 7
move 1 from 5 to 9
move 3 from 8 to 6
move 8 from 7 to 1
move 6 from 7 to 1
move 10 from 1 to 9
move 3 from 6 to 2
move 2 from 1 to 3
move 2 from 3 to 6
move 3 from 7 to 4
move 2 from 7 to 1
move 1 from 2 to 5
move 13 from 9 to 5
move 12 from 9 to 3
move 6 from 5 to 3
move 2 from 9 to 1
move 11 from 9 to 3
move 1 from 4 to 6
move 2 from 5 to 3
move 1 from 1 to 8
move 24 from 3 to 5
move 2 from 9 to 3
move 2 from 2 to 4
move 1 from 9 to 2
move 2 from 6 to 8
move 5 from 3 to 5
move 2 from 8 to 9
move 1 from 9 to 8
move 4 from 1 to 4
move 1 from 9 to 4
move 1 from 8 to 4
move 1 from 8 to 4
move 7 from 4 to 5
move 1 from 1 to 8
move 1 from 6 to 5
move 35 from 5 to 4
move 18 from 4 to 3
move 6 from 4 to 3
move 8 from 5 to 8
move 8 from 8 to 1
move 2 from 4 to 9
move 23 from 3 to 1
move 1 from 8 to 5
move 1 from 9 to 1
move 1 from 5 to 1
move 1 from 9 to 4
move 11 from 1 to 2
move 16 from 4 to 5
move 3 from 3 to 5
move 9 from 2 to 5
move 1 from 4 to 1
move 2 from 2 to 6
move 1 from 2 to 9
move 1 from 6 to 2
move 1 from 3 to 5
move 1 from 3 to 9
move 1 from 2 to 9
move 23 from 1 to 5
move 1 from 6 to 9
move 1 from 9 to 8
move 27 from 5 to 1
move 1 from 9 to 3
move 18 from 5 to 8
move 6 from 5 to 7
move 1 from 5 to 6
move 1 from 9 to 8
move 12 from 8 to 3
move 1 from 1 to 4
move 6 from 7 to 8
move 1 from 6 to 3
move 1 from 4 to 2
move 2 from 1 to 8
move 1 from 2 to 9
move 8 from 3 to 2
move 2 from 9 to 7
move 5 from 2 to 7
move 7 from 7 to 2
move 2 from 8 to 2
move 3 from 1 to 9
move 5 from 1 to 2
move 3 from 9 to 8
move 3 from 8 to 7
move 5 from 2 to 5
move 2 from 7 to 6
move 12 from 8 to 9
move 12 from 1 to 4
move 9 from 9 to 3
move 4 from 5 to 8
move 12 from 3 to 8
move 1 from 7 to 9
move 3 from 9 to 2
move 1 from 4 to 7
move 3 from 1 to 7
move 7 from 4 to 6
move 3 from 6 to 2
move 2 from 7 to 9
move 18 from 8 to 1
move 2 from 4 to 7
move 1 from 2 to 8
move 1 from 8 to 2
move 10 from 2 to 3
move 3 from 9 to 8
move 2 from 6 to 7
move 13 from 3 to 1
move 2 from 8 to 9
move 28 from 1 to 8
move 1 from 5 to 2
move 1 from 4 to 3
move 4 from 7 to 6
move 5 from 6 to 7
move 7 from 2 to 6
move 1 from 9 to 6
move 2 from 2 to 4
move 1 from 9 to 1
move 4 from 1 to 2
move 3 from 2 to 5
move 3 from 4 to 9
move 3 from 5 to 7
move 1 from 1 to 4
move 6 from 7 to 6
move 1 from 2 to 6
move 1 from 4 to 1
move 1 from 1 to 8
move 3 from 9 to 4
move 18 from 6 to 3
move 4 from 3 to 6
move 1 from 7 to 9
move 1 from 6 to 9
move 2 from 3 to 6
move 1 from 9 to 6
move 1 from 9 to 2
move 6 from 6 to 8
move 3 from 4 to 7
move 2 from 7 to 2
move 35 from 8 to 7
move 3 from 3 to 1
move 26 from 7 to 2
move 10 from 3 to 9
move 6 from 9 to 4
move 3 from 1 to 2
move 1 from 4 to 3
move 4 from 4 to 1
move 1 from 3 to 6
move 1 from 8 to 3
move 1 from 6 to 2
move 1 from 3 to 2
move 13 from 7 to 3
move 3 from 1 to 4
move 4 from 3 to 1
move 3 from 1 to 9
move 2 from 1 to 9
move 10 from 2 to 9
move 19 from 2 to 9
move 6 from 3 to 9
move 2 from 3 to 4
move 2 from 2 to 6
move 17 from 9 to 8
move 1 from 2 to 8
move 2 from 9 to 3
move 2 from 6 to 7
move 8 from 9 to 3
move 5 from 4 to 5
move 14 from 9 to 4
move 1 from 2 to 3
move 1 from 7 to 2
move 2 from 9 to 3
move 1 from 2 to 7
move 5 from 5 to 1
move 1 from 2 to 1
move 1 from 3 to 1
move 1 from 9 to 7
move 3 from 7 to 2
move 3 from 3 to 7
move 1 from 2 to 4
move 1 from 3 to 8
move 1 from 2 to 4
move 4 from 3 to 4
move 16 from 8 to 9
move 3 from 1 to 4
move 21 from 4 to 6
move 1 from 7 to 2
move 1 from 8 to 2
move 1 from 1 to 3
move 6 from 6 to 7
move 3 from 1 to 9
move 3 from 7 to 3
move 1 from 4 to 6
move 1 from 4 to 7
move 2 from 2 to 6
move 1 from 8 to 6
move 13 from 6 to 7
move 1 from 2 to 3
move 15 from 9 to 8
move 6 from 6 to 3
move 13 from 8 to 3
move 4 from 9 to 4
move 5 from 4 to 8
move 19 from 3 to 9
move 3 from 3 to 1
move 5 from 8 to 9
move 17 from 9 to 7
move 1 from 1 to 8
move 4 from 9 to 6
move 3 from 3 to 8
move 1 from 1 to 2
move 3 from 3 to 1
move 36 from 7 to 6
move 1 from 1 to 2
move 7 from 8 to 2
move 24 from 6 to 5
move 2 from 6 to 7
move 1 from 3 to 2
move 4 from 6 to 8
move 19 from 5 to 1
move 8 from 6 to 4
move 7 from 2 to 5
move 3 from 2 to 8
move 15 from 1 to 6
move 2 from 9 to 5
move 2 from 7 to 8
move 3 from 4 to 1
move 4 from 5 to 6
move 1 from 9 to 7
move 1 from 8 to 3
move 3 from 6 to 1
move 2 from 4 to 7
move 13 from 1 to 8
move 1 from 3 to 7
move 1 from 4 to 5
move 19 from 8 to 6
move 1 from 7 to 3
move 8 from 5 to 8
move 1 from 6 to 8
move 3 from 5 to 9
move 1 from 6 to 4
move 3 from 4 to 7
move 1 from 3 to 9
move 4 from 7 to 9
move 20 from 6 to 3
move 1 from 8 to 4
move 2 from 9 to 4
move 2 from 9 to 2
move 2 from 9 to 3
move 13 from 6 to 9
move 9 from 9 to 8
move 2 from 6 to 3
move 8 from 8 to 2
move 2 from 7 to 3
move 5 from 9 to 3
move 12 from 3 to 5
move 1 from 4 to 7
move 8 from 2 to 4
move 8 from 4 to 7
move 2 from 2 to 6
move 2 from 8 to 9
move 2 from 6 to 8
move 2 from 9 to 6
move 2 from 6 to 9
move 2 from 4 to 8
move 2 from 9 to 2
move 6 from 3 to 1
move 2 from 2 to 9
move 3 from 9 to 3
move 8 from 7 to 2
move 6 from 1 to 2
move 8 from 3 to 8
move 1 from 7 to 3
move 5 from 3 to 8
move 6 from 2 to 7
move 3 from 7 to 6
move 2 from 7 to 9
move 1 from 7 to 8
move 8 from 5 to 7
move 7 from 2 to 1
move 7 from 1 to 6
move 7 from 7 to 9
move 1 from 7 to 6
move 2 from 3 to 9
move 2 from 8 to 5
move 25 from 8 to 5
move 5 from 5 to 1
move 1 from 6 to 4
move 17 from 5 to 4
move 5 from 5 to 4
move 23 from 4 to 7
move 2 from 5 to 2
move 4 from 6 to 3
move 6 from 3 to 7
move 1 from 5 to 2
move 1 from 1 to 7
move 2 from 2 to 8
move 2 from 2 to 9
move 1 from 5 to 7
move 4 from 1 to 6
move 2 from 8 to 3
move 2 from 9 to 4
move 1 from 4 to 8
move 7 from 9 to 1
move 2 from 3 to 5
move 28 from 7 to 4
move 4 from 6 to 2
move 2 from 6 to 2
move 3 from 7 to 4
move 2 from 5 to 6
move 4 from 2 to 6
move 9 from 6 to 5
move 4 from 1 to 7
move 1 from 6 to 2
move 3 from 2 to 3
move 1 from 8 to 6
move 1 from 7 to 4
move 2 from 3 to 4
move 1 from 7 to 4
move 2 from 1 to 6
move 1 from 7 to 9
move 1 from 7 to 9
move 1 from 6 to 2
move 7 from 5 to 8
move 1 from 3 to 9
move 1 from 5 to 2
move 7 from 8 to 7
move 4 from 4 to 8
move 2 from 8 to 4
move 2 from 2 to 7
move 1 from 1 to 7
move 1 from 5 to 6
move 32 from 4 to 7
move 2 from 6 to 5
move 2 from 8 to 2
move 1 from 2 to 1
move 2 from 5 to 4
move 1 from 2 to 5
move 1 from 1 to 4
move 4 from 4 to 3
move 1 from 6 to 4
move 1 from 5 to 4
move 5 from 9 to 1
move 4 from 3 to 5
move 3 from 1 to 6
move 2 from 9 to 5
move 2 from 1 to 3
move 15 from 7 to 1
move 5 from 5 to 3
move 1 from 5 to 2
move 3 from 4 to 5
move 2 from 5 to 9
move 3 from 3 to 6
move 3 from 3 to 4
move 1 from 3 to 8
move 1 from 9 to 3
move 2 from 4 to 9
move 1 from 5 to 3
move 2 from 9 to 6
move 1 from 8 to 1
move 1 from 3 to 2
move 1 from 4 to 9
move 2 from 9 to 3
move 9 from 1 to 3
move 5 from 3 to 4
move 2 from 1 to 3
move 4 from 1 to 5
move 1 from 2 to 8
move 3 from 4 to 9`
