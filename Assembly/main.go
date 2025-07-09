package main

import (
	"fmt"
)

func main() {
	fmt.Println("Genome assembly!")
	fmt.Println(ScoreOverlapAlignment("ATACG", "GCTACG", 1, 2, 3))
}
