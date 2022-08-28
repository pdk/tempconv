package main

import (
	"fmt"
	"os"

	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s f\n", os.Args[0])
		os.Exit(1)
	}

	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot parse \"%s\" as number\n", os.Args[1])
		os.Exit(2)
	}

	//[°C] = ([°F] − 32) × 5⁄9 (https://en.wikipedia.org/wiki/Celsius)
	c := (f - 32.0) * 5.0 / 9.0

	fmt.Printf("%.2f\n", c)
}
