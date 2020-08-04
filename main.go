package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// backend team
var team = map[int]string{
	0: "Brandon",
	1: "Bigo",
	2: "Jeff",
	3: "Benny",
	4: "Andy",
	5: "Eric",
	6: "Leilani",
	7: "Carmen",
}

func main() {
	// seed it!
	rand.Seed(time.Now().UnixNano())

	// pick winner
	winner()
}

func winner() {
	results := make(map[string]int)

	// run 10 million times
	for i := 0; i < 10000000; i++ {
		i := rand.Intn(len(team))
		results[team[i]]++
	}

	// find winner
	winner := ""
	top := 0
	fmt.Printf("\n")
	for name, count := range results {
		if count > top {
			winner = name
			top = count
		}
		fmt.Printf("%10v => %d\n", name, count)
	}

	// drumroll, please
	fmt.Printf("\nThe winner is .......\n")

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("🥁\n")
			}
		}
	}()
	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true

	fmt.Printf("\n%s!!\n", strings.ToUpper(winner))
}
