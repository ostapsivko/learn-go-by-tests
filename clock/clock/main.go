package main

import (
	"clock"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
