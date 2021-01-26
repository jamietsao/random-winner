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
	count := len(names)

	fmt.Printf("First to %d wins!\n", target)
	fmt.Printf("%s", strings.Repeat("\n", count))

	// run until winner is determined
	winner := ""
	results := make(map[string]int)
	for {
		// random choice
		i := rand.Intn(count)
		results[names[i]]++

		// move cursor back to top left
		fmt.Printf("\u001b[%dD\u001b[%dA", 50, count)

		// print current progress
		multiplier := target / 100
		for _, name := range names {
			width := results[name] / (2 * multiplier)
			fmt.Printf("[%10v => %s%s]\n", name, strings.Repeat("#", width), strings.Repeat(" ", 50-width))
		}

		// stop once someone hits the target
		if results[names[i]] == target {
			winner = names[i]
			break
		}

		time.Sleep(1 * time.Millisecond)
	}

	// display final results
	fmt.Printf("\nFinal tally:\n")
	for _, name := range names {
		fmt.Printf("%10v: %d\n", name, results[name])
	}

	// drumroll, please
	fmt.Printf("\nThe winner is ....... %s!\n", strings.ToUpper(winner))
}
