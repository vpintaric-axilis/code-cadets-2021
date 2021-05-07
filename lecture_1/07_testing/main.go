package main

import "flag"

func main() {
	var countFrom, countTo, divisor int

	flag.IntVar(&countFrom, "count-from", 1, "Value (inclusive) from which to start counting")
	flag.IntVar(&countTo, "count-to", 10, "Value (inclusive) to count to")
	flag.IntVar(&divisor, "divisor", 1, "The divisor used for filtering")


}
