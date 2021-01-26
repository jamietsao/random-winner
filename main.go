package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	target = 1000
)

func init() {
	// seed rng
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var nameStr string

	// parse flags
	flag.StringVar(&nameStr, "names", "", "Comma separated list of names")
	flag.Parse()

	names := strings.Split(nameStr, ",")

	if len(names) < 2 {
		fmt.Println("Must supply at least two names via -names flag:")
		flag.Usage()
		os.Exit(0)
	}

	// pick winner
	winner(names)
}

func winner(names []string) {
	nameCount := len(names)
	longest := 0
	for _, names := range names {
		if len(names) > longest {
			longest = len(names)
		}
	}

	fmt.Printf("First to %d wins!\n", target)
	fmt.Printf("%s", strings.Repeat("\n", nameCount))

	// run until winner is determined
	runs := 0
	winner := ""
	results := make(map[string]int)
	for {
		runs++

		// random choice
		i := rand.Intn(nameCount)
		results[names[i]]++

		// move cursor back to top left
		fmt.Printf("\u001b[%dD\u001b[%dA", 50, nameCount)

		// print current progress
		multiplier := target / 100
		for _, name := range names {
			width := results[name] / (2 * multiplier)
			fmt.Printf("[%*v => %s%s]\n", longest+1, name, strings.Repeat("#", width), strings.Repeat(" ", 50-width))
		}

		// stop once someone hits the target
		if results[names[i]] == target {
			winner = names[i]
			break
		}

		time.Sleep(1 * time.Millisecond)
	}

	// display final results
	fmt.Printf("\nFINAL TALLY:\n")
	for _, name := range names {
		fmt.Printf("%*v: %d\n", longest+1, name, results[name])
	}
	fmt.Printf("%*v: %d\n", longest+1, "TOTAL", runs)

	// drumroll, please
	fmt.Printf("\nThe winner is ....... %s!\n", strings.ToUpper(winner))
}
