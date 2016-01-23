// Copyright 2016, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mitchellh/cli"
)

const timeFmt = "Mon, 2006-01-02 15:04:05 MST -0700"

var flagUTC bool

func main() {
	err := loadDB()
	if err != nil {
		fmt.Println("failed to load DB", err)
		os.Exit(-1)
	}

	if len(os.Args) == 1 {
		now := time.Now()
		fmt.Println(now.Format(timeFmt))
		for name, offset := range locs {
			loc := time.FixedZone(name, offset)
			fmt.Println(now.In(loc).Format(timeFmt))
		}
		os.Exit(0)
	}

	c := cli.NewCLI("wdate", "1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"add":    addCommandFactory,
		"remove": removeCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
