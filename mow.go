package main

import (
	"fmt"
	"os"
	"sync"
)

// mow loads a configuration file, mows the given lawn and outputs
// the final mowers positions
func mow(filename string) {
	conf := ParseFile(filename)

	lawn := conf.Lawn
	mowers := conf.Mowers

	var wg sync.WaitGroup
	wg.Add(len(mowers))

	for _, mower := range mowers {
		// Mowers mow in parallel (but cannot mow the same plot).
		go mower.Mow(&lawn, &wg)
	}

	// Wait for the deployed mowers to finish their instructions.
	wg.Wait()

	for _, mower := range mowers {
		fmt.Println(mower.String())
	}
}

func main() {
	if len(os.Args) > 1 {
		mow(os.Args[1])
	} else {
		fmt.Println("no configuration file given")
		os.Exit(1)
	}
}
