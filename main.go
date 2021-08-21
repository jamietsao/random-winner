package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	target = 1000
	snail  = "🐌"
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
			// show progress as a portion of 50 marks
			width := results[name] / (2 * multiplier)
			fmt.Printf("[%*v _%s%s%s]\n", longest+1, name, strings.Repeat("_", width), snail, strings.Repeat(" ", 50-width))
		}

		// stop once someone hits the target
		if results[names[i]] == target {
			winner = names[i]
			break
		}

		time.Sleep(1 * time.Millisecond)
	}

	// sort & display final results
	sort.Slice(names, func(i, j int) bool {
		// reverse sort to show highest to lowest
		return results[names[i]] > results[names[j]]
	})

	fmt.Printf("\n  FINAL RESULTS  \n")
	fmt.Printf("------------------\n")
	for _, name := range names {
		fmt.Printf("%*v %5d\n", longest+1, name, results[name])
	}
	fmt.Printf("\n%*v %5d\n", longest+1, "TOTAL", runs)

	// drumroll, please
	fmt.Printf("\nThe winner is ....... %s!\n", strings.ToUpper(winner))
}
