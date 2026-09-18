package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"Lab3/core"
)

func runRandomAlgorithm() {
	fmt.Println("[MMU] Executing Random Page Replacement algorithm...")
}

func runClockAlgorithm() {
	fmt.Println("[MMU] Executing Clock (Second Chance) algorithm...")
}

func main() {
	frames := flag.Int("frames", 32, "number of frames")
	algoPtr := flag.String("algo", "Random", "MMU algorithm: Random, Clock")

	flag.Parse()

	algorithm := strings.ToLower(*algoPtr)

	switch algorithm {
	case "random":
		fmt.Println("Selected Algorithm: Random ", frames)
		runRandomAlgorithm()

	case "clock":
		fmt.Println("Selected Algorithm: Clock ", frames)
		runClockAlgorithm()

	default:
		fmt.Fprintf(os.Stderr, "Error: unknown MMU algorithm '%s'\n", *algoPtr)
		flag.Usage() // Prints the automatic help message
		os.Exit(1)   // Exits the application with an error code
	}

	core := core.NewCore()
	core.Start("cool", "cool")
}

