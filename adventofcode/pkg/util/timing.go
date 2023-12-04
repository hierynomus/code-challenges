package util

import (
	"fmt"
	"time"
)

func Timing(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
