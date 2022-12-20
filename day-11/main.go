package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

const numRounds = 10000
const reliefModifier = 3

var lcm int = 1

func main() {
	monkeys := monkeysFromInput(pt1In)
	for _, m := range monkeys {
		lcm *= m.test.num
	}

	for i := 0; i < numRounds; i++ {
		monkeyBusiness(monkeys)
		if (i+1)%1000 == 0 {
			fmt.Printf("== After round %d ==\n", (i + 1))

			for _, m := range monkeys {
				fmt.Printf("Monkey %d inspected items %d times.\n", m.id, m.inspects)
			}
			fmt.Println()
		}
	}

	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspects > monkeys[j].inspects
	})

	// for _, m := range monkeys {
	// 	 fmt.Printf("%+v\n", m)
	// }

	total := 1
	for i := 0; i < 2; i++ {
		total *= monkeys[i].inspects
	}
	fmt.Printf("shenanigans: %d\n", total)
}

func monkeyBusiness(monkeys []*monkey) {
	for _, m := range monkeys {
		// fmt.Printf("Monkey %d:\n", m.id)
		for _, worryLevel := range m.items {
			// fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", worryLevel)
			cp := new(big.Int)
			cp.Set(worryLevel)
			worryLevel = m.inspect(cp)
			// worryLevel = worryLevel.Quo(worryLevel, big.NewInt(reliefModifier))
			// fmt.Printf("    Monkey gets bored with item. Worry level is divided by %d to %d.\n", reliefModifier, worryLevel)

			cp = new(big.Int)
			cp.Set(worryLevel)
			nextMonkey := m.testWorry(cp)
			// fmt.Printf("    Item with worry level %s is thrown to monkey %d.\n", worryLevel.String(), nextMonkey)
			monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, worryLevel)
		}
		m.items = []*big.Int{}
	}
}

type monkey struct {
	id                              int
	items                           []*big.Int
	op                              op
	test                            op
	inspects                        int
	nextMonkeyTrue, nextMonkeyFalse int
}

func (m *monkey) inspect(worryLevel *big.Int) *big.Int {
	m.inspects += 1
	delta := big.NewInt(int64(m.op.num))
	if m.op.selfOp {
		delta = new(big.Int)
		delta.Set(worryLevel)
	}
	switch m.op.mathOp {
	case mathOpAdd:
		worryLevel = worryLevel.Add(worryLevel, delta)
		// fmt.Printf("    Worry level increases by %s to %d.\n", delta.String(), worryLevel)
	case mathOpMult:
		worryLevel = worryLevel.Mul(worryLevel, delta)
		// ret := delta.String()
		// fmt.Printf("    Worry level is multiplied by %s to %s.\n", ret, worryLevel.String())
	}

	return worryLevel.Rem(worryLevel, big.NewInt(int64(lcm)))
}

func (m *monkey) testWorry(worryLevel *big.Int) int {
	switch m.test.mathOp {
	case mathOpDiv:
		rem := worryLevel.Rem(worryLevel, big.NewInt(int64(m.test.num))).Int64() == 0
		if rem {
			// fmt.Printf("    Current worry level is divisible by %d.\n", m.test.num)
			return m.nextMonkeyTrue
		} else {
			// fmt.Printf("    Current worry level is not divisible by %d.\n", m.test.num)
			return m.nextMonkeyFalse
		}
	}
	panic("help")
}

func (m monkey) string() string {
	buf := fmt.Sprintf("Monkey %d:\n", m.id)
	buf += "  Starting Items: "
	for _, i := range m.items {
		buf += fmt.Sprintf("%d, ", i)
	}
	buf += "\n"
	var endArg string
	if m.op.selfOp {
		endArg = "self"
	} else {
		endArg = fmt.Sprintf("%d", m.op.num)
	}

	buf += fmt.Sprintf("  Operation: new = old %s %s\n", m.op.mathOp, endArg)
	buf += fmt.Sprintf("  Test: divisible by %d\n", m.test.num)
	buf += fmt.Sprintf("    If true: throw to monkey %d\n", m.nextMonkeyTrue)
	buf += fmt.Sprintf("    If false: throw to monkey %d\n", m.nextMonkeyFalse)
	return buf
}

type mathOp string

type op struct {
	mathOp mathOp
	selfOp bool
	num    int
}

const (
	mathOpAdd  mathOp = "+"
	mathOpMult mathOp = "*"
	mathOpDiv  mathOp = "/"
)

func mathOpFromStr(in string) mathOp {
	switch mathOp(in) {
	case mathOpAdd:
		return mathOpAdd
	case mathOpMult:
		return mathOpMult
	case mathOpDiv:
		return mathOpDiv
	default:
		panic(in)
	}
}

func monkeysFromInput(in string) []*monkey {
	monkeys := []*monkey{}
	var cur *monkey
	for _, line := range strings.Split(in, "\n") {
		line = strings.TrimPrefix(line, "  ")
		// fmt.Printf("parsing '%s'\n", line)
		if strings.HasPrefix(line, "Monkey") {
			id := parseInt(strings.Split(line, " ")[1])
			cur = &monkey{id: id}
			// fmt.Printf("got id %d\n", id)
		} else if strings.HasPrefix(line, "Starting") {
			cur.items = []*big.Int{}
			items := strings.SplitAfter(line, ":")
			items = strings.SplitAfter(items[1], ",")
			for _, item := range items {
				cur.items = append(cur.items, big.NewInt(int64(parseInt(item))))
			}
			// fmt.Printf("got items %+v\n", cur.items)
		} else if strings.HasPrefix(line, "Operation") {
			cur.op = op{}
			parts := strings.Split(line, " ")
			cur.op.mathOp = mathOpFromStr(parts[4])
			if parts[5] == "old" {
				cur.op.selfOp = true
			} else {
				cur.op.num = parseInt(parts[5])
			}
			// fmt.Printf("got op %+v\n", cur.op)
		} else if strings.HasPrefix(line, "Test") {
			cur.test = op{mathOp: mathOpDiv}
			parts := strings.Split(line, " ")
			cur.test.num = parseInt(parts[3])
			// fmt.Printf("got test %+v\n", cur.test)
		} else if strings.HasPrefix(line, "If") {
			parts := strings.Split(line, " ")
			// fmt.Printf("If parts: %+v\n", parts)
			num := parseInt(parts[5])
			if parts[1] == "true:" {
				cur.nextMonkeyTrue = num
				cur.nextMonkeyFalse = num
			} else if parts[1] == "false:" {
				cur.nextMonkeyFalse = num
				// If false: is the end of each logical group
				// fmt.Printf("appending monkey: %+v\n", *cur)
				monkeys = append(monkeys, cur)
			}
		}
	}

	return monkeys
}

func parseInt(in string) int {
	in = strings.ReplaceAll(in, " ", "")
	in = strings.ReplaceAll(in, ":", "")
	in = strings.ReplaceAll(in, ",", "")
	ret, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(ret)
}

const testIn = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

const pt1In = `Monkey 0:
Starting items: 78, 53, 89, 51, 52, 59, 58, 85
Operation: new = old * 3
Test: divisible by 5
  If true: throw to monkey 2
  If false: throw to monkey 7

Monkey 1:
Starting items: 64
Operation: new = old + 7
Test: divisible by 2
  If true: throw to monkey 3
  If false: throw to monkey 6

Monkey 2:
Starting items: 71, 93, 65, 82
Operation: new = old + 5
Test: divisible by 13
  If true: throw to monkey 5
  If false: throw to monkey 4

Monkey 3:
Starting items: 67, 73, 95, 75, 56, 74
Operation: new = old + 8
Test: divisible by 19
  If true: throw to monkey 6
  If false: throw to monkey 0

Monkey 4:
Starting items: 85, 91, 90
Operation: new = old + 4
Test: divisible by 11
  If true: throw to monkey 3
  If false: throw to monkey 1

Monkey 5:
Starting items: 67, 96, 69, 55, 70, 83, 62
Operation: new = old * 2
Test: divisible by 3
  If true: throw to monkey 4
  If false: throw to monkey 1

Monkey 6:
Starting items: 53, 86, 98, 70, 64
Operation: new = old + 6
Test: divisible by 7
  If true: throw to monkey 7
  If false: throw to monkey 0

Monkey 7:
Starting items: 88, 64
Operation: new = old * old
Test: divisible by 17
  If true: throw to monkey 2
  If false: throw to monkey 5`
