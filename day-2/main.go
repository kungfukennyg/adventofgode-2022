package main

import (
	"errors"
	"fmt"
	"strings"
)

type Outcome int

const (
	Win  = 6
	Loss = 0
	Draw = 3
)

func (o Outcome) String() string {
	switch o {
	case Win:
		return "Win"
	case Loss:
		return "Loss"
	case Draw:
		return "Draw"
	default:
		panic("unrecognized Outcome")
	}
}

type Move int

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func (m Move) String() string {
	switch m {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		panic("unrecognized Move")
	}
}

func main() {
	testInput := "A Y\nB X\nC Z"
	totalPoints, err := doWerk(testInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("testinput, got pt1 points %d\n\n", totalPoints)

	totalPoints, err = doWerkPt2(realInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("realinput, got pt2 points %d\n", totalPoints)

}

func doWerk(input string) (int, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		// A Y
		split := strings.Split(line, " ")
		if len(split) != 2 {
			return 0, errors.New(fmt.Sprintf("expected to have len of %d, got %d", 2, len(split)))
		}
		oppMove := getMove(split[0], false)
		urMove := getMove(split[1], true)

		outcome := getOutcome(oppMove, urMove)

		fmt.Printf("moves: %s : %s, outcome: %s, score change: %d\n",
			fmt.Sprint(oppMove), fmt.Sprint(urMove), fmt.Sprint(outcome), int(outcome)+int(urMove))
		total += int(outcome) + int(urMove)

	}

	return total, nil
}

func doWerkPt2(input string) (int, error) {
	var total int
	for _, line := range strings.Split(input, "\n") {
		// A Y
		split := strings.Split(line, " ")
		if len(split) != 2 {
			return 0, errors.New(fmt.Sprintf("expected to have len of %d, got %d", 2, len(split)))
		}
		oppMove := getMove(split[0], false)
		outcome := getIntendedOutcome(split[1])

		urMove := determineMoveToMeetOutcome(oppMove, outcome)

		fmt.Printf("moves: %s : %s, outcome: %s, score change: %d\n",
			fmt.Sprint(oppMove), fmt.Sprint(urMove), fmt.Sprint(outcome), int(outcome)+int(urMove))
		total += int(outcome) + int(urMove)

	}

	return total, nil
}

func getOutcome(oppMove, playerMove Move) Outcome {
	if oppMove == playerMove {
		return Draw
	}
	if oppMove == Rock {
		if playerMove == Scissors {
			return Loss
		} else {
			return Win
		}
	} else if oppMove == Paper {
		if playerMove == Rock {
			return Loss
		} else {
			return Win
		}
	} else if oppMove == Scissors {
		if playerMove == Paper {
			return Loss
		} else {
			return Win
		}
	}

	panic(fmt.Sprintf("unrecognized match up %d, %d", oppMove, playerMove))
}

func getIntendedOutcome(move string) Outcome {
	switch move {
	case "X":
		return Loss
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		panic(fmt.Sprintf("unrecognized input %s", move))
	}
}

func determineMoveToMeetOutcome(opp Move, intended Outcome) Move {
	if intended == Draw {
		return opp
	}

	if intended == Win {
		if opp == Rock {
			return Paper
		} else if opp == Paper {
			return Scissors
		} else if opp == Scissors {
			return Rock
		}
	} else if intended == Loss {
		if opp == Rock {
			return Scissors
		} else if opp == Scissors {
			return Paper
		} else if opp == Paper {
			return Rock
		}
	}

	panic("shouldn't happen")
}

func getMove(move string, isPlayer bool) Move {
	if isPlayer {
		switch move {
		case "X":
			return Rock
		case "Y":
			return Paper
		case "Z":
			return Scissors
		default:
			panic(fmt.Sprintf("unrecognized input %s", move))
		}
	} else {
		switch move {
		case "A":
			return Rock
		case "B":
			return Paper
		case "C":
			return Scissors
		default:
			panic(fmt.Sprintf("unrecognized input %s", move))
		}
	}
}

const realInput string = "B Y\nA Y\nB Z\nA Z\nA Y\nB Z\nC X\nC X\nC X\nC Y\nC Z\nB Y\nC Y\nC Z\nA Y\nB Y\nC Y\nB Y\nB Y\nB Y\nC X\nB Z\nA X\nA Z\nC Z\nC Y\nC Y\nB Y\nB X\nC Z\nB Y\nB Y\nC Y\nB Y\nB Z\nB Z\nB Y\nA Y\nA Y\nB Z\nB Y\nB Y\nB Y\nC Y\nA Y\nB Y\nC Z\nB Y\nB Y\nA Z\nB Y\nA Y\nB Y\nB Z\nC Y\nC Z\nA Z\nC Z\nB Y\nA X\nC Z\nA X\nA Z\nB Y\nB Y\nA Y\nC Z\nB Y\nB Z\nB Z\nB X\nC Y\nB Y\nA Y\nA Y\nB Y\nA Z\nB X\nB Y\nB Y\nC Y\nC Z\nA Z\nB Y\nA Y\nB Z\nB Y\nB Y\nB Y\nC X\nC Y\nB Y\nB Y\nB Z\nB Y\nC X\nB Y\nB Y\nC Z\nC Z\nA X\nA X\nA X\nA X\nB Y\nC Z\nB Z\nB Y\nC Z\nB Y\nB Z\nB Y\nC X\nB Y\nC X\nC Z\nB Z\nC Z\nC Y\nB Y\nB Y\nA Z\nC Z\nC Y\nC Y\nB Y\nB X\nC Y\nB Z\nC Z\nC Z\nA Z\nB Y\nB Y\nB Z\nB Y\nB Y\nC X\nA Z\nA Y\nB Y\nC Y\nB X\nC Y\nC X\nC Y\nA Y\nB Y\nA Y\nA Z\nC Z\nC Z\nA Z\nC Y\nB Y\nC Z\nC Y\nB Y\nB Y\nC X\nB Y\nB Y\nC Y\nC X\nB X\nC Y\nA Z\nB Y\nA Z\nC Y\nC Z\nB Y\nA Z\nC X\nC X\nB Y\nB Y\nC Z\nB Y\nA Y\nB Y\nA Z\nA X\nA Y\nB Y\nB Y\nA Y\nC Z\nB Z\nB Y\nB Y\nB Y\nB Y\nC X\nC Z\nC Y\nB Y\nA Y\nA Z\nC Y\nB Y\nB Y\nB Z\nB Z\nB Z\nA Y\nB Z\nB Y\nC Y\nB Y\nC X\nC Y\nC Y\nB Y\nA Y\nA Z\nB Y\nB Y\nA Y\nC X\nC X\nB Y\nA Z\nA X\nC Y\nC Y\nC Y\nC X\nC Y\nB Y\nB Z\nB Y\nB X\nA Z\nB Y\nB Z\nA X\nB Y\nC Z\nB Y\nB Z\nB Y\nB Y\nB Y\nB X\nC Z\nC Z\nA Y\nB Y\nC Y\nA Y\nC X\nC Z\nA Z\nC Y\nB Y\nB Y\nC Z\nA Z\nC Y\nC X\nC Z\nB X\nB Y\nC Y\nB Y\nC X\nB Y\nA X\nB Y\nB Y\nA Z\nB Y\nB Y\nC X\nB Y\nA X\nA Y\nB Y\nA Y\nC Y\nC X\nC X\nB Z\nB Y\nC Y\nC Y\nC Y\nB Y\nA Y\nC Y\nB Y\nB Y\nB Y\nB Y\nA Z\nB Y\nB Y\nA Z\nB Y\nC Z\nC Y\nB Y\nB Z\nA Y\nB Y\nA Z\nC Z\nB Y\nC Z\nB Z\nB Z\nA X\nA Z\nC Z\nB Y\nA Y\nC Z\nC Y\nB Y\nB X\nB Y\nC Y\nC Y\nC Y\nC X\nC Y\nB Y\nB Y\nC Y\nB Y\nC Z\nC Z\nB Y\nA Y\nC Y\nB Z\nC X\nC Y\nB Y\nB Y\nB Y\nC Z\nB Z\nB Y\nA Y\nB Y\nA Z\nB Y\nB Z\nB Y\nB Z\nB Y\nB Z\nC Z\nC Z\nA X\nA Y\nB Y\nC Z\nC Y\nB Z\nA Y\nA X\nB Y\nB Y\nB Y\nB Z\nB Y\nB Y\nC Y\nB Y\nB Y\nC Y\nB Y\nA Z\nC Z\nC Z\nC Z\nA Z\nB Y\nB Z\nB Y\nA Z\nB Y\nB Y\nB Y\nB Y\nB Y\nA X\nC Y\nB Y\nC Z\nB Y\nB Y\nB Y\nC Y\nA Z\nB Y\nB Z\nB Y\nB Y\nB Y\nB Y\nB Y\nC Y\nB Z\nA Z\nC Y\nC X\nC X\nB Y\nC Z\nB Z\nB Z\nB Y\nB Y\nB Y\nA Z\nB Z\nA Z\nB Z\nA Z\nB Y\nB X\nB Y\nC Z\nC Y\nB Y\nA Z\nA Z\nB Y\nB Y\nB Y\nB Y\nA Z\nB Z\nB Y\nB Y\nB Z\nB Y\nC Z\nB Y\nC Y\nB Y\nC Y\nB Y\nA Z\nB Y\nA Z\nC Z\nC Y\nA Z\nA Y\nA Y\nA X\nA Y\nB Y\nB Z\nA X\nC X\nC Z\nB Y\nA Z\nA X\nB Y\nB Y\nB Y\nB Z\nC Y\nC X\nB Z\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nA Z\nC Y\nB Z\nB Y\nB Y\nB Y\nA X\nB Z\nA Y\nC Z\nB Z\nB Y\nB Y\nC Y\nA Y\nC Y\nB Y\nB Y\nA X\nC X\nB Z\nC Y\nB Y\nB Z\nC Y\nB Y\nB Z\nB X\nB Y\nA Z\nB Y\nB Z\nB Y\nB Y\nA Z\nA X\nA Z\nB Z\nA X\nC Y\nB Z\nB Y\nB X\nA X\nC Y\nC X\nA Y\nB Z\nB Y\nA X\nA Z\nC X\nC Z\nB Z\nC X\nC X\nB Y\nA Y\nB Y\nB Y\nC Z\nA X\nC X\nB Y\nB Y\nB Y\nB Y\nC X\nB Z\nB Y\nB Y\nC X\nB X\nB Y\nC X\nA Z\nA Y\nB Y\nB Y\nB Y\nC Z\nB Y\nC X\nC X\nB Z\nB Z\nB Z\nB X\nB Y\nC Y\nB Z\nB Y\nB Y\nB Y\nB Z\nB Y\nC Y\nB Y\nB Y\nB Y\nB Y\nB Z\nC Z\nA Y\nC Y\nB Y\nB Y\nA Z\nC Z\nB Y\nB Y\nA Y\nB Y\nC Z\nC Z\nB Y\nA Y\nA Y\nC Y\nB Z\nB Z\nC X\nB Y\nC X\nB Z\nC Z\nA Z\nA Y\nB Z\nC X\nB Y\nA X\nA Z\nA Y\nB Y\nB X\nB Z\nA Z\nB Z\nA X\nA X\nA Y\nB Y\nC Y\nB Y\nB Y\nC Z\nB Y\nA X\nB Y\nC Y\nB Z\nC X\nB Y\nB Y\nB Z\nA Z\nB Y\nC Z\nB Y\nC Y\nB Y\nA Y\nB Y\nB Y\nA Z\nA Y\nC Y\nB Y\nB Z\nC Y\nB Y\nB Y\nC Z\nB Y\nC Z\nA X\nA Z\nB Y\nB Y\nC Z\nB Y\nB Y\nC Y\nA X\nC Z\nB Y\nA Y\nB Y\nA X\nA Z\nB Z\nB Z\nB Y\nB Y\nC Y\nA Y\nC X\nB Y\nA Z\nC Y\nC X\nA Y\nA Z\nC X\nC Z\nA Y\nB Y\nB Y\nC X\nB Y\nA Y\nB Z\nA X\nC Z\nA X\nB Z\nA Z\nA Z\nB Z\nB Y\nB Y\nB Y\nB Y\nC Y\nB Y\nA X\nA Y\nA Y\nB Y\nB Y\nC Z\nA Y\nB Z\nB Y\nC Y\nA X\nB Y\nA X\nA X\nA Z\nC Y\nA Z\nC Z\nB Y\nB Y\nB Y\nA X\nB Z\nC Z\nA X\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nC Y\nB Y\nB Y\nC Y\nC Y\nC Z\nA Y\nC Y\nB Z\nB Y\nC Y\nB Y\nC Y\nA Y\nB Y\nB Y\nB Y\nB Z\nB Z\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nC X\nA Z\nA X\nB Y\nB Z\nA Z\nA Z\nB Y\nB Y\nA Z\nA Y\nA Z\nA Z\nA Y\nB Y\nB Y\nB Y\nB Y\nC X\nB Y\nB Y\nB Y\nB Y\nB Z\nB Z\nB Y\nC Z\nA Z\nA Y\nB Y\nB X\nC Y\nC Y\nA Z\nA X\nB Y\nA X\nB Y\nC Z\nB Y\nC X\nB Y\nB Y\nB Y\nC X\nC Y\nB Y\nB Y\nB Y\nA Y\nB Y\nC X\nB Z\nB Y\nB Y\nC Z\nB Y\nB Y\nA Z\nB Z\nB Y\nB Z\nB X\nA Y\nB Y\nB Y\nB Y\nB Y\nC X\nC Y\nB Y\nA Z\nA Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Z\nC Z\nA X\nB Y\nA Y\nB Y\nB Y\nB Y\nB Y\nB Y\nA Y\nB Z\nB Y\nB Y\nC Y\nA X\nA X\nB X\nC Y\nB Y\nA Y\nB Y\nB Z\nB Y\nC Y\nB Z\nB Z\nA Y\nA Y\nC Z\nB Y\nC Z\nC X\nA Y\nB Y\nA Y\nA X\nB Y\nC Y\nA Z\nB Y\nB Y\nB Y\nA Z\nB Y\nA X\nC X\nA X\nB Z\nC Y\nA X\nC Y\nA Z\nC Y\nB Y\nB Y\nB Y\nB Y\nC Y\nC Z\nA Z\nB Z\nB Z\nC Z\nB Y\nA Y\nB Z\nB Z\nA Y\nB Z\nC Z\nC Y\nC X\nC Z\nC X\nB Y\nA X\nB Z\nB Y\nA X\nB Y\nB Y\nB Z\nC Y\nC Z\nC Y\nB Y\nB Y\nB Z\nA X\nB Y\nA X\nB Z\nA Z\nA Z\nB Y\nC Y\nC Y\nC Y\nB Y\nA Z\nB Z\nB Z\nC Y\nA Y\nA Y\nA Z\nB Y\nA X\nB Z\nC X\nB Z\nB Y\nA Y\nA Y\nC Y\nA Z\nB X\nA X\nC Y\nB Y\nB Y\nC Y\nB Z\nC X\nC Y\nB Y\nC Y\nB Y\nB Y\nC Y\nB X\nA Z\nB Z\nB Y\nA Z\nB Z\nB Y\nA Z\nC Y\nC Y\nB Y\nC Y\nB Y\nB Y\nB Y\nB Y\nB Z\nB Y\nC Y\nC Z\nB Y\nB Y\nA Z\nB Y\nB Z\nA Y\nA Y\nC Y\nA Z\nC Y\nB Z\nB Y\nC Y\nB Z\nC Y\nA Z\nB Y\nB Y\nB Y\nB Y\nC Z\nA Y\nA X\nC X\nA Y\nC Y\nC X\nB Y\nC Y\nA Y\nB Y\nC Y\nB Y\nA X\nB Y\nB Y\nB Z\nC Y\nC Y\nA X\nB Z\nA Z\nB Y\nB Y\nB Y\nB X\nB Y\nB Y\nC Y\nB Y\nC X\nA Y\nB Y\nB Y\nA Z\nC X\nA X\nA X\nC X\nB Y\nC Z\nB Y\nA Z\nC Y\nB Y\nB Y\nC X\nB Z\nC Y\nA X\nA Y\nC Y\nA Z\nB Y\nB Z\nC Y\nA Z\nA Y\nA Y\nB Y\nA X\nB Z\nB X\nB Z\nB Y\nB Y\nA Y\nC Y\nB Y\nB Z\nB Y\nA Y\nC Y\nC Y\nA X\nB Y\nA X\nC X\nA Z\nA Y\nB Z\nB Y\nB Y\nB Y\nC Y\nA X\nB Z\nA X\nC Y\nC Z\nA X\nB Z\nC X\nB Y\nB Y\nA Y\nB Z\nB X\nB Y\nB Y\nA Z\nC Y\nB Y\nB Y\nC Y\nB Y\nA X\nB Y\nB Y\nB Z\nB Y\nA Y\nA Z\nB Y\nB Y\nA Y\nA Z\nA Z\nA Y\nB Y\nA Y\nB Z\nB Y\nB Y\nA Z\nB Y\nB Y\nB X\nA X\nB Y\nB Y\nC Z\nC X\nA Y\nB Z\nB Y\nC Y\nB Y\nB Y\nB Y\nA X\nA Y\nB Y\nC Z\nB Y\nB Y\nB Y\nC Z\nB Y\nB Z\nB Y\nB Y\nB Y\nC Y\nB Y\nB Y\nC Y\nC Y\nA X\nC Y\nA Y\nB Y\nB Y\nB Y\nB Y\nA Y\nA Y\nC X\nB Y\nC Z\nA X\nB X\nC Z\nC Y\nB Y\nB Y\nB X\nB Z\nC Y\nB Z\nB Y\nB Y\nB Y\nB Y\nC Z\nC Y\nC Y\nA Y\nB Y\nB Z\nB Y\nC Z\nB Y\nB Y\nC Y\nB Y\nB Z\nB Z\nB Y\nC Z\nB Y\nB Y\nB Y\nB Y\nB Y\nA Y\nC X\nA Z\nA Y\nB Y\nC X\nB Z\nB Y\nC Y\nB Y\nA Y\nB Y\nB Y\nB Y\nA Z\nA Z\nB Y\nB Y\nB Z\nA Z\nC Y\nC Y\nC Y\nC Z\nB Y\nC Y\nC X\nA Z\nB Z\nB X\nB X\nC Y\nB Y\nA Z\nB Z\nB Y\nB Y\nB Z\nB Y\nB Y\nB X\nB Y\nB Y\nA Z\nA Z\nA Y\nB Z\nB Y\nA Y\nB Y\nA Z\nB Y\nB Y\nA X\nB Z\nC Y\nA Z\nC Y\nC Y\nB Y\nC X\nC Y\nA X\nA Z\nC Z\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nC X\nC Y\nC Y\nC Y\nB Y\nB Y\nC Y\nC Y\nA X\nB X\nC X\nC Z\nB Y\nA Y\nA Z\nB Y\nC Y\nC Z\nB Y\nA Y\nC Y\nB Y\nA Y\nB Y\nA Y\nC Z\nC Y\nA Y\nA Z\nB Y\nC Y\nB Y\nB Y\nB Y\nB Y\nB Z\nC Z\nB Y\nB Y\nC Z\nA Y\nB Y\nC Y\nB Y\nA X\nB Y\nB Y\nA Y\nB Y\nB Y\nA X\nB Y\nB Z\nB Y\nB Y\nB Y\nC Z\nC Y\nA Z\nB Y\nC Y\nC Z\nB Y\nA Y\nB Y\nB Y\nB Y\nC X\nB Z\nC Y\nC X\nA X\nC Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Z\nC Y\nB Y\nC X\nB Y\nB Y\nB Y\nB Z\nA Y\nA Y\nA Z\nB Y\nC Y\nA Y\nA X\nB Z\nC Y\nB Y\nB Y\nB Y\nB Y\nB Z\nA Y\nB Y\nC Y\nC Z\nB X\nB Z\nB Y\nB Y\nC Y\nB X\nA Y\nC Y\nB Y\nC Y\nC X\nA Y\nB Y\nB Z\nC Z\nB Y\nB Y\nC Z\nA Z\nA Y\nB Y\nB Y\nB Y\nC Y\nB Y\nB Y\nB Z\nB Z\nC Y\nB Y\nA Z\nC Y\nC Y\nB Y\nA Z\nB Y\nB Y\nB Y\nA Y\nA X\nC Y\nC Z\nC Y\nB Z\nB Y\nB Y\nB Y\nC Y\nB Y\nB Y\nB Y\nB Z\nB Y\nB Y\nC X\nC Y\nA Z\nC X\nB Y\nB Y\nB Z\nB Y\nA X\nA X\nC X\nB Y\nB Y\nB Y\nA Z\nB Z\nB Z\nB Y\nC Y\nB Y\nB Y\nC Y\nA Z\nC Y\nB Y\nB Z\nA Z\nA Y\nB Z\nA Z\nB Y\nB Y\nA Z\nB Y\nC Z\nB Y\nB Y\nB Z\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nA Z\nB Y\nB Y\nC Y\nB Y\nA Y\nB Y\nC Y\nB Y\nB Y\nB Y\nC Z\nC Y\nC X\nB Y\nA X\nC Y\nB Z\nC Y\nB Y\nC X\nA Y\nB Y\nA Z\nA X\nB Z\nC Y\nC Y\nA Z\nC Z\nB Z\nB Y\nB Y\nB Y\nB Y\nA Y\nA Z\nB Y\nB Y\nA Y\nC X\nB Y\nB Y\nC Y\nC X\nB X\nC Y\nB Y\nA Y\nB Y\nB Y\nB Y\nC Y\nB Y\nB Y\nA Z\nB Y\nA Z\nB Y\nB Y\nB Z\nB Y\nC Y\nB Y\nA Y\nB Y\nA Y\nB Y\nB Z\nB Y\nB Y\nC X\nB Y\nB Y\nC X\nC Z\nC Z\nB Y\nB Z\nB Y\nA Y\nB Y\nA Y\nB Y\nC Y\nA Y\nB Y\nC Y\nB Y\nB Y\nB Y\nA Y\nC Y\nC Z\nC X\nC Y\nC X\nC Y\nB Y\nC Z\nC Z\nB Y\nA Z\nB Z\nB X\nC Y\nA Y\nC Y\nC Y\nB Y\nB Y\nB Y\nB Y\nC X\nB Y\nC Y\nC X\nC Y\nC X\nB Y\nA Y\nA Y\nB Y\nC X\nC X\nC X\nC X\nC X\nB Y\nB Y\nB Y\nC Y\nC Z\nC Y\nB Y\nA Y\nC Z\nB Y\nB Y\nA Y\nB Y\nC Y\nA Z\nB Y\nA Z\nC X\nC Y\nC Y\nA X\nC Y\nC Z\nB Y\nB Y\nC Y\nB Z\nB Z\nB Y\nA Z\nA X\nB Y\nB Y\nA Z\nC Y\nC Z\nA Z\nC Z\nB Y\nC Y\nB Z\nC Z\nC Y\nA X\nA Y\nB Z\nC Z\nB Y\nB Y\nC Z\nC X\nB Y\nB Y\nB Z\nB Y\nA Y\nA Z\nB Y\nA Y\nB Y\nB Y\nA Z\nC X\nC Y\nB Y\nB Y\nA X\nB Y\nB Z\nC Y\nA Y\nC Y\nB Y\nC Y\nB Y\nB Y\nB Y\nC X\nB X\nB Y\nB Y\nB Y\nB Z\nB Z\nB Y\nB X\nB Y\nB X\nB Z\nA X\nC X\nB Y\nB Y\nA Z\nC X\nA Y\nB Y\nB Z\nC Z\nB Y\nB Z\nA Y\nA Z\nB Y\nC Z\nB Y\nA Z\nB Y\nA Z\nA Y\nC Y\nB Y\nB Y\nA X\nC Y\nB Y\nA X\nB Y\nC Z\nA Z\nB Y\nB X\nA Y\nC Y\nB Z\nC Y\nB Y\nB Z\nA Y\nA Z\nB Y\nB Y\nA Z\nA X\nC Y\nB Z\nB Z\nB Y\nC Z\nC Y\nC Z\nB Y\nB Y\nB Z\nB Y\nB Y\nA X\nB Y\nB Y\nC X\nC Z\nB Y\nB Y\nA X\nB Y\nB Z\nC Z\nB Z\nA Y\nB Y\nB Z\nB Y\nC Y\nC Y\nB Y\nB Y\nB Z\nA X\nB Z\nB Y\nC Y\nA X\nB Y\nA Y\nB Z\nA Y\nB Y\nB Y\nB Z\nC Y\nC Y\nB Z\nC Y\nB Y\nC Y\nB Y\nA Z\nA Z\nC Z\nB Y\nB Y\nA Z\nB Z\nB Z\nC Y\nB Y\nB Y\nB Z\nC Y\nA Z\nC Y\nC Z\nC Z\nA Y\nB Y\nC Y\nB Z\nB Y\nC Z\nB Y\nB Y\nA X\nB Y\nB Y\nB Y\nB Y\nA Z\nC Y\nA Y\nB Y\nB Y\nC Z\nB Y\nB Y\nB Y\nB Z\nB Y\nB X\nC Y\nB Y\nB Y\nA X\nC X\nA X\nB Z\nA Z\nB Y\nB Y\nC X\nB Z\nB Y\nB Z\nB Y\nC Z\nA X\nB Z\nB Y\nB Y\nB Z\nB Y\nB Y\nA Y\nB Y\nA X\nB Z\nC X\nB Y\nB Y\nB Y\nB Y\nB Y\nB Z\nB Y\nB Y\nC Z\nB Y\nA Y\nA X\nC Y\nB Z\nB Z\nB Y\nB Y\nA Y\nB Z\nB Y\nB Y\nB Y\nB Y\nB Z\nB X\nB Y\nB Y\nC Y\nB Y\nB Y\nB Y\nB Z\nB Z\nB Y\nB Y\nB Y\nC X\nC Z\nC Z\nA Z\nA Y\nC Y\nB Y\nC Z\nC Y\nC X\nB Y\nB Y\nA X\nA Y\nC X\nB Y\nB Y\nC Y\nB Y\nB Y\nB Y\nC Z\nB Y\nC Z\nC Y\nB Z\nB Y\nB Y\nA X\nC Z\nC Y\nB Y\nB Y\nC Y\nC Y\nB Z\nB Y\nB Z\nB Y\nC Y\nB Y\nC Y\nB Y\nC Y\nC X\nB Y\nC Y\nB Y\nA Z\nB Y\nA Y\nC Y\nC Y\nB Y\nB Y\nC X\nA Z\nA X\nB Z\nB Y\nB Z\nA Z\nC Y\nA Y\nA X\nC Y\nA Y\nA Y\nC Z\nB Y\nB Y\nB Y\nC X\nB Y\nA X\nB Y\nB Y\nB Y\nA Z\nA X\nB Y\nC Y\nB Y\nA X\nB Z\nA Y\nA X\nA Z\nC Z\nB Z\nC X\nB Y\nA X\nB Z\nA Z\nB Y\nA Z\nB Y\nA Z\nC Z\nA Y\nB Y\nB Y\nA X\nB Y\nA X\nB Y\nA Z\nC Z\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nB Y\nA Y\nC X\nC Y\nC Z\nB Y\nB Y\nA X\nC Z\nA Y\nB Y\nB Y\nC Z\nB Z\nB Y\nB Z\nB Y\nB Y\nB Y\nC Y\nA Y\nC Z\nB Y\nA X\nB Y\nC Y\nB Y\nA X\nA X\nB Y\nC Y\nB Y\nB Y\nA Y\nC X\nC Y\nC X\nA Y\nB Y\nB Y\nA Y\nB Y\nC Y\nC Z\nB Z\nB Y\nB Y\nC Y\nB Z\nB Y\nB Y\nB Y\nC Y\nB Y\nC Y\nB Y\nC Z\nB Y\nB Z\nB Y\nB Y\nA Z\nB Y\nB Z\nB Y\nB Z\nA Y\nB Y\nB Y\nB Y\nA Y\nC X\nB Y\nC Y\nB Y\nB Y\nB Y\nC Y\nB Y\nB Y\nC Z\nB Y\nA Z\nC Y\nC Z\nB Y\nB Y\nC Y\nB Z\nB Y\nA Y\nA Y\nB Y\nB Y\nB Y\nB X\nB Z\nA X\nC Z\nB Z\nA Z\nB Y\nB Y\nA X\nA Y\nC Y\nB Y\nB Z\nB Y\nB Y\nB Y\nB Z\nA Y\nB Y\nB Y\nB Y\nC Z\nC Z\nB Y\nB Z\nB Y\nB Y\nB Y\nC Y\nA Z\nB Y\nB Y\nB Z\nB Y\nB Y\nB Y\nA X\nB Y\nB Y\nB Y\nC Y\nB Y\nB Y\nC Z\nB Z\nC Y\nB Y\nB Z\nC Y\nB Y\nB Z\nB Y\nB Y\nA Z\nC Z\nC Y\nB Y\nC X\nB Y\nB Y\nB Z\nB Y\nB Y\nA Y\nB Z\nC Y\nC Y\nA Y\nB Z\nC X\nC Y\nC X\nA Z\nA X\nB Y\nC X\nB Y\nB Y\nB Z\nA Z\nB Z\nB X\nB Y\nB Z\nB Y\nB X\nA Y\nB Y\nC X\nA Z\nA X\nB Y\nB Y\nB Y\nB Z\nB Y\nA X\nC Z\nB Z\nB Z\nB Z\nA Y\nB Y\nB Y\nC X\nB Y\nA Y\nA Z\nB Z\nB Z\nB Y\nB Y\nC Y\nC Y\nC X\nB Y\nC Z\nB Z\nB Y\nB Y\nB Y\nB Z\nC X\nC X\nB Y\nB Y\nB Z\nB Y\nA Y\nB Y\nB Y\nA X\nB Y\nB Y\nB Z\nA Y\nB Y\nA X\nB X\nA Y\nC Y\nB Y\nB X\nA X\nB Y\nB Z\nA Y\nB Z\nB Y\nB Y\nC Y\nB Y\nB Y\nB Y\nB Y\nB Y\nA Y\nB Y\nA Y\nB Y\nA Y\nB Y\nB Y\nB Y\nA X\nA Y\nA Y\nA X\nB Z\nA X\nC Y\nC Y\nB Y\nA X\nB Y\nA Y\nB Y\nB Y\nB Y\nC X\nC Y\nB Y\nC X\nC Y\nB Y\nB Y\nC Y\nB Y\nB Y\nB Y\nB Z\nA Z\nB Y\nB Z\nB Z\nA X\nC Z\nB Y\nC Z\nB Y\nC Z\nB Y\nA X\nC Y\nC Z\nA Z\nA Y\nB Y\nA Y\nB Y\nB Y\nB Y\nB Y\nB Y\nA Y\nC Y\nC Y\nB Z\nB Y\nB Y\nB Z\nA Z\nA X\nC Z\nB Y\nB Y\nA X\nC Y\nA Y\nB Y\nA Y\nB Y\nB Z\nC Y\nC Z\nB Y\nA Y\nB Y\nC Z\nC Z\nA X\nA X\nC Z\nA X\nA Z\nB Y\nB Y\nC Z\nC Y\nC Z\nA X\nA X\nB Z\nB Y\nB Y\nA X\nB Y\nB Z\nC Y\nC X\nB Y\nC Z\nB Y\nC Z\nB Y\nB Z\nB Y\nC Y\nA Z\nB Y\nB Y\nA X\nA Z\nA X\nC Y\nB Y\nB Y\nC Z\nB Z\nC X\nB Y\nC Y\nC Y\nA Z\nC Z\nB Z\nB Y\nB Z\nB Y\nA Z\nA Y\nA Y\nA X\nB Z\nC Z\nC X\nB Y\nA X\nC X\nA Y\nB Y\nB Y\nB Y\nA Z\nB Y\nC X\nB Y\nA Z\nB Z\nB Z\nC Y\nB Y\nC X\nC X\nC Y\nC Y\nC Z\nC X\nB Y\nC Y\nC Y\nB Z\nB Y\nA Y\nB Z\nA X\nB Y\nA Z\nC Y\nB Y\nC Z\nB Y\nA X\nC Z\nB Z\nB Y\nB Y\nB Y\nB Y\nB Y\nC Z\nA Z\nC Y\nA Z\nB Y\nC X\nC X\nA X\nA Y\nB Z\nC Y\nC X\nC X\nC Y\nC X\nB Y\nC X\nA Y\nC Z\nB Y\nB Z\nB Y\nA X\nB Y\nA Z\nA Y\nC X\nA X\nB X\nC X\nC Z\nC X\nA Y\nA Z\nB Y\nA Y"
