package main

import "errors"

const timeFmt = "Monday, 2006-01-02 15:04:05 MST -0700"

var errInvalidOffset = errors.New("Invalid time offset format")
