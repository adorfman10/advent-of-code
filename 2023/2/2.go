package two

import (
	"strconv"
	"strings"

	"github.com/adorfman10/advent-of-code/util"
)

type game struct {
	id    int
	blue  int
	red   int
	green int
}

func A() int {
	defer util.Duration(util.Track("A"))

	input := util.ReadFile("2/input.txt")
	lines := strings.Split(input, "\n")
	// games := make([]game, len(lines))
	sum := 0
	for _, line := range lines {
		game := gameParser(line)
		if game.red > 12 || game.green > 13 || game.blue > 14 {
			continue
		}
		sum += game.id
	}

	return sum
}

func B() int {
	defer util.Duration(util.Track("B"))

	input := util.ReadFile("2/input.txt")
	lines := strings.Split(input, "\n")
	// games := make([]game, len(lines))
	sum := 0
	for _, line := range lines {
		game := gameParser(line)
		sum += game.red * game.green * game.blue
	}
	return sum
}

func gameParser(s string) game {
	var g game
	colonSplit := strings.Split(s, ":")
	g.id, _ = strconv.Atoi(strings.Split(colonSplit[0], " ")[1])

	rounds := colonSplit[1]
	round := strings.Split(rounds, ";")
	for _, entry := range round {
		colorAndValues := strings.Split(entry, ",")
		for _, cv := range colorAndValues {
			c := strings.Split(cv[1:], " ")[1]
			v, _ := strconv.Atoi(strings.Split(cv[1:], " ")[0])
			switch c {
			case "red":
				if v > g.red {
					g.red = v
				}
			case "green":
				if v > g.green {
					g.green = v
				}
			case "blue":
				if v > g.blue {
					g.blue = v
				}
			}
		}
	}
	return g
}
