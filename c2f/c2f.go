package main

import (
	"fmt"
	"os"

	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s c\n", os.Args[0])
		os.Exit(1)
	}

	c, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: cannot parse \"%s\" as number\n", os.Args[1])
		os.Exit(2)
	}

	// [°F] = [°C] × 9⁄5 + 32 (https://en.wikipedia.org/wiki/Celsius)
	f := c*9.0/5.0 + 32.0

	fmt.Printf("%.2f\n", f)
}
