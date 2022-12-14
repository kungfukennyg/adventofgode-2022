package main

import (
	"fmt"
	"strconv"
	"strings"
)

const testIn string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

const testMaxX int = 6
const testMaxY int = 5

func main() {
	start(input, testMaxX, testMaxY)
}

func start(in string, _ int, _ int) {
	cmds := cmdsFromStr(in)
	// for _, c := range cmds {
	// fmt.Println(c)
	// }

	fmt.Println("== Initial State ==")
	fmt.Println()
	head := &pos{0, 0}
	tail := &pos{0, 0}
	// printGrid(*head, *tail, maxX, maxY)

	visited := map[pos]struct{}{}
	for i, cmd := range cmds {
		fmt.Printf("%d\n", i)
		fmt.Printf("\n== %s ==\n\n", cmd.String())
		for step := 1; step <= cmd.steps; step++ {
			move := pos{x: 0, y: 0}
			switch cmd.dir {
			case Up:
				move.y = 1
			case Down:
				move.y = -1
			case Left:
				move.x = -1
			case Right:
				move.x = 1
			default:
				panic(cmd)
			}

			b4 := *head
			head.walk(move)
			fmt.Printf("walking head (%s) to (%s)\n", b4, head)

			b4 = *tail
			tail.walkToHead(*head)
			fmt.Printf("walking tail (%s) to (%s)\n", b4, tail)

			visited[*tail] = struct{}{}
			// printGrid(*head, *tail, maxX, maxY)
			fmt.Println()
		}
	}

	fmt.Printf("visited %d positions\n", len(visited))
}

type pos struct {
	x, y int
}

func printGrid(head, tail pos, maxX, maxY int) {
	for y := maxY - 1; y >= 0; y-- {
		for x := 0; x < maxX; x++ {
			cur := pos{x, y}
			if head == cur {
				fmt.Printf("H ")
			} else if tail == cur {
				fmt.Printf("T ")
			} else if cur.x == 0 && cur.y == 0 {
				fmt.Printf("s ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
}

func (p *pos) walk(o pos) {
	p.x += o.x
	p.y += o.y
}

func (p pos) String() string {
	return fmt.Sprintf("%d, %d", p.x, p.y)
}

func (tail *pos) walkToHead(head pos) {
	// overlapping
	if tail.x == head.x && tail.y == head.y {
		fmt.Println("overlap")
		return
	}

	if tail.y == head.y {
		// in same row
		fmt.Println("row")
		if tail.x < head.x-1 {
			tail.x += 1
		} else if tail.x > head.x+1 {
			tail.x -= 1
		}
	} else if tail.x == head.x {
		// in same column
		fmt.Println("column")
		if tail.y < head.y-1 {
			tail.y += 1
		} else if tail.y > head.y+1 {
			tail.y -= 1
		}
	} else {
		if tail.y > head.y && tail.y-head.y > 1 {
			// tail is higher than head
			tail.y -= 1
			if tail.x > head.x && tail.x-head.x >= 1 {
				// and to the right
				tail.x -= 1
			} else if tail.x < head.x && head.x-tail.x >= 1 {
				// and to the left
				tail.x += 1
			}
			// 0 < 1
		} else if tail.y < head.y && head.y-tail.y > 1 {
			// tail is lower than head
			tail.y += 1
			if tail.x > head.x && tail.x-head.x >= 1 {
				// and to the right
				tail.x -= 1
			} else if tail.x < head.x && head.x-tail.x >= 1 {
				// and to the left
				tail.x += 1
			}
		} else if tail.x > head.x && tail.x-head.x > 1 {
			// tail is to the right
			tail.x -= 1
			if tail.y > head.y && tail.y-head.y >= 1 {
				// and higher than head
				tail.y -= 1
			} else if tail.y < head.y && head.y-tail.y >= 1 {
				// and lower than head
				tail.y += 1
			}
		} else if tail.x < head.x && head.x-tail.x > 1 {
			// tail is to the left
			tail.x += 1
			if tail.y > head.y && tail.y-head.y >= 1 {
				// and higher than head
				tail.y -= 1
			} else if tail.y < head.y && head.y-tail.y >= 1 {
				// and lower than head
				tail.y += 1
			}
		}
	}
}

type dir string

const (
	Up    = "U"
	Down  = "D"
	Left  = "L"
	Right = "R"
)

type cmd struct {
	dir   dir
	steps int
}

func (c cmd) String() string {
	return fmt.Sprintf("%s %d", c.dir, c.steps)
}

func cmdsFromStr(in string) []cmd {
	cmds := []cmd{}
	for _, line := range strings.Split(in, "\n") {
		parts := strings.Split(line, " ")
		cmd := cmd{
			dir: dir(parts[0]),
		}

		steps, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		cmd.steps = int(steps)
		cmds = append(cmds, cmd)
	}

	return cmds
}

const input string = `U 1
L 1
U 2
R 1
U 2
R 2
L 1
R 1
L 1
D 1
L 2
D 2
U 1
R 2
D 1
L 1
D 2
U 2
R 2
D 1
R 1
D 1
U 2
D 2
U 1
R 1
L 1
D 2
R 2
D 1
R 1
D 2
L 1
D 2
R 1
U 2
R 1
L 1
U 2
R 1
L 2
U 1
R 2
D 2
R 2
D 1
L 2
U 1
R 1
L 1
U 2
L 1
U 2
R 2
D 1
U 2
D 1
R 1
D 2
L 2
U 2
L 1
R 1
D 1
U 2
R 1
L 2
U 2
L 2
U 1
L 1
D 1
U 2
D 1
R 1
L 1
D 1
L 2
D 2
L 2
R 1
U 2
L 2
U 2
D 2
L 2
D 2
U 2
R 2
U 2
R 1
D 2
R 2
L 2
D 1
U 1
L 2
R 1
U 1
L 1
D 2
U 2
D 2
L 1
R 1
L 1
U 1
D 2
U 1
D 1
R 2
U 2
R 2
D 1
U 1
L 3
D 3
R 2
L 2
R 2
L 2
D 2
L 3
D 1
L 3
R 3
L 3
U 3
D 3
U 2
L 2
D 1
U 2
R 1
D 1
R 2
L 2
U 1
R 3
D 1
R 2
L 1
U 2
D 1
U 2
R 2
U 1
L 2
U 2
L 1
D 1
L 2
D 3
L 3
D 1
R 3
U 3
R 3
D 1
L 1
R 3
L 3
D 2
U 2
L 3
D 1
L 1
U 3
L 1
D 3
U 1
R 1
L 3
R 2
U 3
L 3
R 1
L 1
D 3
L 2
D 1
U 3
D 1
L 3
U 2
R 3
U 2
L 1
D 2
U 3
D 2
R 1
U 3
R 2
L 3
D 1
R 1
L 2
D 1
U 1
L 3
D 2
U 2
R 1
L 3
R 3
U 1
L 2
D 2
L 1
U 3
D 2
U 1
R 3
L 2
R 2
L 2
R 1
D 1
U 2
L 2
R 3
L 3
U 4
R 3
D 3
R 1
D 4
R 1
D 4
R 1
D 1
R 3
L 4
D 1
L 1
U 1
D 3
U 4
D 2
L 3
R 1
D 3
U 4
D 4
L 2
R 3
U 2
L 1
D 2
U 2
L 2
R 2
D 4
R 4
L 1
R 3
U 2
D 1
U 3
L 2
D 4
U 1
L 1
D 1
U 3
R 4
L 3
R 2
U 2
R 1
U 3
D 1
U 3
L 2
U 2
L 3
D 1
R 3
L 1
U 4
R 1
L 2
R 1
D 2
R 4
L 1
D 1
L 3
U 3
D 2
U 4
L 1
R 4
L 3
D 1
U 2
L 4
D 2
L 3
R 3
L 3
D 2
U 2
L 4
D 4
U 3
L 2
R 1
U 4
D 3
L 4
D 2
L 3
D 2
L 2
U 4
D 4
R 4
D 2
L 1
U 4
D 4
L 1
R 3
L 1
D 1
L 3
D 3
R 4
L 1
R 3
U 3
R 4
L 1
U 3
R 2
U 1
L 5
R 5
D 2
R 1
U 5
L 2
R 1
U 3
D 1
U 4
L 2
U 4
L 2
R 3
U 2
R 2
U 1
L 3
U 2
D 5
R 3
U 1
R 2
U 4
L 1
D 5
R 4
U 3
D 1
U 1
R 2
U 4
R 1
L 3
D 4
U 2
R 1
L 1
R 3
L 5
R 3
L 5
D 1
R 3
D 4
U 2
D 3
L 3
D 2
L 1
R 1
D 4
L 5
U 5
D 1
L 3
R 5
D 5
U 2
L 5
U 1
L 4
D 5
R 5
U 4
D 2
L 1
D 3
L 5
R 3
D 3
L 2
U 5
R 1
L 5
D 5
R 3
L 1
U 3
D 5
R 1
U 1
R 4
L 2
U 3
R 2
D 5
L 1
U 1
L 3
R 3
D 5
U 5
D 4
R 2
L 1
U 3
L 2
R 1
L 1
R 3
D 3
U 1
L 4
D 2
U 4
L 6
D 6
U 2
R 3
U 1
D 5
L 1
D 6
U 2
R 2
L 1
R 4
D 4
L 1
U 4
L 5
R 5
L 2
U 5
R 5
U 1
D 3
L 5
D 1
R 5
L 6
U 3
D 3
U 6
R 4
L 6
R 5
D 5
R 3
L 5
D 1
L 2
R 1
L 4
U 5
R 2
U 3
D 1
L 4
R 3
L 1
R 3
U 3
R 3
L 2
D 2
L 1
R 4
D 1
L 6
D 1
R 1
U 6
L 5
R 3
L 3
D 5
L 1
D 6
R 2
D 3
U 6
D 2
U 5
L 5
D 1
R 4
D 2
U 3
D 5
L 2
D 2
U 4
R 2
L 2
D 5
L 2
R 2
D 5
R 6
D 5
L 1
U 2
R 1
D 4
L 2
R 1
D 4
L 2
R 1
U 5
R 1
L 6
U 4
R 6
U 3
L 5
U 2
R 5
D 2
U 6
D 2
U 6
L 6
U 6
R 2
U 4
D 3
U 6
L 7
R 7
L 7
R 4
L 7
R 1
L 4
R 4
L 2
D 5
U 4
D 2
R 1
L 5
U 2
R 2
U 3
L 6
U 4
D 3
U 5
R 1
U 3
L 2
D 6
L 6
U 3
L 1
D 1
U 1
L 1
D 7
R 1
U 7
R 3
D 7
R 6
L 1
R 5
L 1
D 6
L 1
D 4
L 1
R 3
D 5
U 4
D 4
U 1
D 6
L 4
D 7
R 4
L 2
U 5
R 1
D 2
L 3
R 6
L 1
D 4
R 7
U 4
D 7
L 3
U 4
D 2
R 6
D 3
R 2
L 4
D 2
R 1
L 5
U 1
R 1
U 3
R 6
D 6
R 5
L 5
R 1
U 5
L 5
D 6
L 1
D 1
L 5
U 4
D 3
R 6
U 2
R 3
D 1
L 4
R 3
L 7
R 3
U 1
D 5
L 6
R 7
D 3
L 7
D 5
L 5
D 4
L 1
D 7
L 7
R 5
L 3
D 4
L 6
R 6
U 1
R 3
D 1
L 4
U 7
L 2
R 8
D 8
R 3
L 7
D 2
R 1
L 4
U 4
L 3
D 2
L 4
U 1
R 3
L 1
D 7
R 2
L 4
U 3
L 2
U 6
L 4
R 2
D 2
R 4
U 3
D 5
U 7
R 2
L 3
D 4
U 5
L 5
U 5
D 8
R 3
L 7
U 1
L 8
U 2
D 3
U 1
L 5
D 8
L 2
D 4
U 3
R 4
U 1
R 5
L 8
R 7
U 7
L 5
U 6
D 7
U 1
R 3
U 8
D 6
L 5
R 5
L 2
U 3
D 3
R 5
D 7
R 2
D 8
L 3
U 3
R 7
D 5
U 7
L 7
U 1
R 1
L 6
R 3
L 4
U 2
R 1
L 3
U 6
R 6
D 5
U 2
D 7
U 5
R 5
D 2
R 5
D 7
U 8
D 6
U 8
R 4
D 2
U 1
L 3
U 5
D 7
L 4
R 1
U 7
D 1
U 6
D 3
L 6
U 9
L 9
U 3
L 8
D 9
U 7
D 3
U 8
L 6
D 8
U 3
D 2
L 7
U 2
L 2
R 8
D 1
L 2
R 5
U 8
R 7
L 9
R 9
U 3
D 9
L 2
D 6
L 2
R 3
D 2
U 9
D 6
U 7
D 8
U 9
R 4
D 6
U 6
L 6
D 6
L 4
U 1
D 3
R 3
D 1
R 1
D 9
U 6
R 1
U 4
D 9
L 5
U 4
D 5
U 8
L 2
D 5
L 2
R 4
D 7
U 3
L 5
R 6
D 3
U 5
L 4
U 4
D 8
R 4
D 5
L 1
D 2
U 6
D 4
L 6
U 4
L 9
R 1
D 2
L 5
U 6
R 4
U 9
D 6
R 8
U 6
L 9
R 4
D 3
R 3
U 7
L 3
R 1
L 8
U 7
D 4
L 8
U 5
D 4
R 7
U 5
L 5
D 6
U 4
R 6
D 4
U 10
L 10
U 8
L 9
D 2
U 1
L 1
D 7
U 6
R 5
D 7
R 7
L 6
R 1
D 5
R 6
L 2
R 9
U 6
L 2
R 8
D 5
L 7
R 10
L 4
R 6
D 1
U 1
L 3
D 3
R 7
U 1
L 10
D 3
L 1
U 3
D 2
U 9
D 6
R 8
D 6
L 2
U 8
D 10
U 8
L 4
D 2
U 1
R 3
L 7
U 9
R 9
U 7
L 9
U 1
D 9
U 6
R 7
L 5
R 7
U 6
D 2
L 5
U 7
D 3
L 1
R 3
L 1
D 10
L 10
D 6
U 7
L 1
R 2
U 2
L 9
R 9
U 5
R 2
D 10
U 10
D 6
L 8
U 4
L 6
U 8
D 8
U 1
L 4
D 10
U 6
D 6
L 10
U 1
D 3
L 10
R 2
U 1
D 1
L 8
D 10
L 8
D 4
L 1
D 3
R 4
D 1
R 8
D 9
U 6
D 3
U 3
L 8
U 4
R 10
D 9
U 4
R 6
L 6
D 1
R 8
D 5
L 10
D 7
U 3
L 1
D 5
U 8
D 3
R 11
U 2
L 10
D 1
R 9
U 5
R 11
U 1
D 5
U 8
R 10
U 6
L 6
R 6
U 10
R 1
L 6
R 8
L 6
D 1
L 5
U 5
R 11
U 4
R 9
U 10
R 9
L 3
U 6
R 7
L 10
U 10
L 9
R 4
U 4
R 1
U 6
L 9
R 5
L 9
D 10
U 10
R 3
D 11
R 10
L 2
R 11
L 6
U 1
R 8
D 1
U 4
R 5
D 4
L 5
U 10
R 6
D 6
U 5
D 2
L 2
D 11
R 3
U 11
L 11
U 1
R 9
U 1
D 2
U 5
R 11
D 8
L 9
D 6
L 8
U 7
L 9
U 11
D 3
L 1
U 11
L 5
D 7
L 1
U 10
L 10
D 8
R 7
D 8
U 11
R 12
L 10
U 12
D 5
R 7
U 11
D 12
U 7
L 12
D 7
L 10
U 5
D 4
R 5
D 2
R 3
U 3
L 12
D 9
L 8
D 10
R 9
U 10
R 5
L 4
U 1
D 5
L 8
U 3
R 11
L 4
D 6
L 11
R 11
U 2
R 2
D 1
L 7
U 12
L 12
R 8
U 12
L 8
U 6
R 10
D 2
L 7
U 10
L 9
D 12
U 4
R 7
U 10
L 10
R 3
U 2
L 6
U 12
D 4
R 7
L 5
R 7
D 9
U 1
D 2
L 12
R 11
D 2
R 4
D 5
R 8
L 2
R 8
L 8
U 8
L 1
U 2
R 1
D 1
U 6
L 10
R 12
D 3
U 1
D 10
U 1
R 2
U 12
R 8
U 11
R 11
D 9
L 4
R 11
L 12
U 7
L 1
U 10
R 9
L 3
U 12
L 12
D 6
R 10
D 11
R 3
D 1
U 10
D 3
R 10
D 4
R 8
L 2
R 6
U 8
R 10
U 5
L 9
R 2
D 10
U 1
D 7
R 1
D 13
R 12
D 10
R 11
L 11
U 13
L 2
D 7
U 10
D 2
R 12
U 13
L 8
D 10
R 9
D 6
U 9
L 3
R 5
L 4
U 10
R 9
U 9
D 8
R 12
L 9
D 11
U 8
R 8
D 3
U 9
D 7
U 8
R 1
U 8
L 12
U 13
L 8
U 4
D 7
U 12
L 8
D 11
U 13
L 8
U 9
L 1
R 13
L 10
R 1
L 4
R 9
L 8
R 13
L 5
U 9
D 10
L 2
R 10
L 5
D 1
L 2
D 6
U 4
L 3
U 11
L 3
R 7
U 5
D 10
U 3
D 13
U 11
R 1
U 11
R 4
L 3
U 11
R 11
L 12
R 13
U 10
R 7
L 3
U 3
R 10
U 4
R 4
L 10
R 10
D 8
L 11
D 8
R 8
U 2
R 2
U 13
L 3
D 1
L 11
R 7
U 1
L 6
U 8
R 4
L 12
U 1
L 4
U 14
D 6
R 10
L 12
D 6
R 5
U 6
L 4
D 4
R 12
U 6
R 1
U 4
L 3
U 1
L 4
D 1
U 8
L 9
D 13
U 12
R 2
D 12
R 4
U 8
R 14
U 3
L 1
D 1
U 10
L 6
U 2
L 8
D 5
R 7
U 10
R 2
U 2
R 2
L 7
U 4
R 3
U 3
L 8
D 13
R 12
L 8
D 11
U 7
L 13
D 11
U 13
D 2
L 1
R 4
L 8
R 11
L 4
R 4
L 2
D 3
R 10
D 14
R 8
L 6
D 2
R 4
L 11
D 5
U 9
L 9
D 8
R 3
U 11
R 5
U 6
L 2
D 7
R 6
D 13
R 6
U 14
L 14
R 9
D 8
U 3
R 11
D 13
R 14
L 3
U 6
D 7
R 13
L 7
R 12
D 12
R 12
U 12
R 8
L 1
U 7
R 10
U 7
L 2
D 12
U 8
D 14
U 6
D 15
R 2
U 6
L 14
U 15
R 10
U 4
R 8
D 13
R 5
D 15
U 9
R 2
U 15
R 3
U 6
D 14
L 6
D 13
R 13
L 7
D 15
U 12
L 4
R 7
D 12
U 15
D 12
R 14
U 6
L 9
D 8
L 4
R 13
L 13
D 3
L 12
D 12
U 5
D 15
L 1
U 14
L 7
U 4
L 11
U 11
L 12
R 6
D 1
L 1
U 6
R 9
U 10
L 12
R 6
U 15
D 5
L 9
D 5
L 13
U 11
D 12
R 3
U 2
R 5
D 7
L 9
D 8
R 8
L 11
D 7
R 3
D 14
L 1
R 15
L 1
U 2
L 9
R 7
U 2
R 13
U 6
R 3
L 6
R 15
L 2
D 15
L 2
R 14
L 12
R 14
D 6
L 4
R 1
L 2
U 10
D 4
U 8
R 3
D 8
U 12
L 3
R 12
L 15
R 8
L 2
D 11
U 3
L 1
D 12
L 8
U 9
L 13
U 2
L 13
U 7
L 6
D 12
U 15
L 7
U 1
D 6
R 15
L 3
U 10
R 1
U 4
R 2
U 3
R 3
L 13
R 16
D 1
L 2
U 9
L 3
D 6
L 15
D 11
L 7
U 11
L 3
U 1
L 15
R 5
U 14
L 8
R 16
D 8
U 13
R 4
U 6
L 4
R 12
L 12
U 13
D 5
R 5
L 13
D 15
R 10
D 8
U 12
R 6
L 5
D 3
U 4
L 6
R 3
L 6
D 9
U 4
L 5
R 3
D 16
R 2
U 16
L 9
U 12
R 5
L 14
D 14
R 15
D 6
L 3
D 6
U 9
D 11
L 4
D 2
L 9
D 16
L 7
R 1
L 16
D 16
U 16
R 1
U 13
D 8
R 11
L 2
R 13
D 7
L 14
R 7
U 8
R 5
L 13
R 3
U 13
R 14
U 15
R 2
D 5
R 16
U 15
D 10
U 4
L 16
U 3
D 8
U 17
L 17
R 17
U 6
D 16
U 3
R 1
L 16
U 6
R 17
L 15
D 16
R 11
D 17
L 10
R 11
U 2
D 11
L 3
R 12
U 15
R 10
D 14
U 2
R 12
D 1
U 13
D 15
U 16
L 3
U 10
D 16
U 4
D 3
R 5
U 12
R 17
D 5
R 9
U 13
R 12
L 4
D 8
U 1
D 16
R 7
D 9
U 14
D 5
L 5
U 4
L 5
R 8
L 3
D 5
L 3
R 12
U 7
R 6
D 2
R 4
D 11
L 15
R 2
U 9
L 12
R 17
L 13
U 7
R 1
D 16
U 9
R 17
D 2
L 17
D 10
R 16
L 9
D 5
R 7
D 14
L 10
U 12
L 14
R 7
D 6
R 10
D 9
R 2
L 1
R 3
U 3
D 11
L 16
D 6
R 12
U 9
D 3
R 11
D 16
L 4
R 5
L 11
R 8
D 14
R 13
U 8
D 16
L 5
R 12
U 9
D 18
L 10
R 16
L 18
D 7
L 6
U 6
R 10
L 9
R 3
D 16
R 5
D 5
R 6
D 15
U 15
L 14
D 11
R 12
D 7
R 18
D 10
U 1
D 1
U 6
D 7
U 12
D 3
L 7
U 15
L 8
R 11
L 14
D 5
U 12
R 8
D 2
U 7
D 12
L 15
D 2
U 14
D 11
U 2
D 18
R 16
U 2
D 1
R 16
D 15
R 17
D 13
L 13
U 10
L 11
U 6
R 5
D 14
R 7
L 5
R 14
D 15
R 4
L 9
D 11
L 4
R 9
U 12
L 5
D 10
R 1
D 5
R 1
L 16
D 12
L 8
U 1
D 9
U 14
L 3
R 11
D 4
L 5
R 17
D 1
L 12
R 4
L 15
R 18
U 7
D 4
R 3
U 15
R 10
L 10
U 16
R 15
L 12
D 11
R 13
D 2
R 12
U 17
R 14
L 8
U 1
L 16
D 11
R 13
D 19
U 5
D 19
L 10
R 18
D 8
R 1
L 10
D 6
U 1
L 14
R 13
L 14
U 9
D 6
L 19
U 18
L 6
R 13
D 3
U 8
R 17
U 14
L 10
R 1
L 9
R 3
U 14
R 6
L 5
D 8
R 15
L 12
R 13
U 4
R 1
L 10
U 11
L 3
U 8
R 16
U 1
D 3
U 8
L 1
D 1
L 9
D 2
U 5
L 8
R 1
U 6
L 1
R 13
U 10
D 14
U 2
L 7
R 9
U 1
R 2
L 5
U 17
R 19
L 18
U 19
D 18
L 6
R 15
U 5
R 17
L 9
U 3
L 9
U 10
L 10
R 5
D 14
U 14
R 13
U 16
D 8
L 1
D 15
R 10
D 15
U 9
D 18
L 8
R 8
L 3
U 6
L 12
R 8
D 15
L 17
U 11
R 1
D 2
R 1
U 1
L 10
D 11
U 6
L 10
U 5
R 14`
