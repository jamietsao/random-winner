package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	target = 1000
)

// backend team
var team = []string{
	"Brandon",
	"Bigo",
	"Jeff",
	"Benny",
	"Andy",
	"Eric",
	"Leilani",
	"Carmen",
}

func main() {
	// seed it!
	rand.Seed(time.Now().UnixNano())

	// pick winner
	winner()
}

func winner() {
	fmt.Printf("First to %d wins!", target)
	fmt.Printf("%s", strings.Repeat("\n", len(team)))

	// run until winner is determined
	winner := ""
	results := make(map[string]int)
	for {
		// random choice
		i := rand.Intn(len(team))
		results[team[i]]++

		// move cursor back to top left
		fmt.Printf("\u001b[%dD\u001b[%dA", 50, len(team))

		// print current progress
		multiplier := target / 100
		for _, name := range team {
			width := results[name] / (2 * multiplier)
			fmt.Printf("[%10v => %s%s]\n", name, strings.Repeat("#", width), strings.Repeat(" ", 50-width))
		}

		// stop once someone hits the target
		if results[team[i]] == target {
			winner = team[i]
			break
		}

		time.Sleep(1 * time.Millisecond)
	}

	// display final results
	fmt.Printf("\nFinal tally:\n")
	for _, name := range team {
		fmt.Printf("%10v: %d\n", name, results[name])
	}

	// drumroll, please
	fmt.Printf("\nThe winner is ....... %s!\n", strings.ToUpper(winner))
}
