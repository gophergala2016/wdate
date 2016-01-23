// Copyright 2016, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/mitchellh/cli"
)

type removeCmd struct{}

func (f removeCmd) Synopsis() string {
	return "remove a timezone"
}

func (f removeCmd) Help() string {
	return "remove a timezone from given name"
}

func (f removeCmd) Run(args []string) int {
	if len(args) != 1 {
		log.Println("fail to remove. give a timezone name")
		return -1
	}

	err := removeLocation(args[0])
	if err != nil {
		log.Println("fail to remove:", err)
		return -1
	}
	return 0
}

var removeCommandFactory = func() (cli.Command, error) {
	var cmd removeCmd
	return cmd, nil
}
