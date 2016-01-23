package main

import (
	"flag"
	"fmt"
	"time"
)

const timeFmt = "Monday, 2006-01-02 15:04:05 MST -0700"

var flagUTC bool

func init() {
	flag.BoolVar(&flagUTC, "u", false, "print UTC")
}

func main() {
	flag.Parse()

	now := time.Now()
	fmt.Println(now.Format(timeFmt))

	if flagUTC {
		fmt.Println(now.UTC().Format(timeFmt))
	}

	loc := time.FixedZone("singapore", +8*60*60+0*60)
	fmt.Println(now.In(loc).Format(timeFmt))
}
