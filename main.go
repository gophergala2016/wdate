package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
)

var flagUTC bool

func main() {
	app := cli.NewApp()
	app.Name = "wdate"
	app.Usage = "print world clock"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "utc",
			Usage:       "show UTC time",
			Destination: &flagUTC,
		},
	}
	app.Action = func(c *cli.Context) {
		now := time.Now()
		err := loadDB()
		if err != nil {
			fmt.Println("failed to load DB", err)
			os.Exit(-1)
		}
		fmt.Println(now.Format(timeFmt))
		if flagUTC {
			fmt.Println(now.UTC().Format(timeFmt))
		}
		for name, offset := range locs {
			loc := time.FixedZone(name, offset)
			fmt.Println(now.In(loc).Format(timeFmt))
		}
	}
	app.Commands = []cli.Command{
		{
			Name:        "add",
			Aliases:     []string{"a"},
			Description: "This add a new timezone",
			Action:      scAdd,
		}, {
			Name:        "remove",
			Aliases:     []string{"r"},
			Description: "This remove a timezone",
			Action:      scRemove,
		},
	}

	app.Run(os.Args)
}

func scAdd(c *cli.Context) {
	args := c.Args()
	if len(args) < 2 {
		fmt.Println("usage: add timezone_name timezone_offset")
		os.Exit(-1)
	}
	name := args[0]
	offset, err := timeOffsetStrToInt(args[1])
	if err != nil {
		fmt.Println("fail to add:", err)
		os.Exit(-1)
	}
	// TODO: add offset
	addLocation(name, offset)
}

func scRemove(c *cli.Context) {
	args := c.Args()
	if len(args) != 1 {
		fmt.Println("usage: remove timezone_name")
		os.Exit(-1)
	}
	// TODO: remove a timezone
	removeLocation(args[0])
}

func timeOffsetStrToInt(s string) (int, error) {
	// +0700 -> 07*60*60 + 30*60
	if len(s) != 5 {
		return 0, errInvalidOffset
	}

	if s[0] != '+' && s[0] != '-' {
		return 0, errInvalidOffset
	}

	hStr, mStr := s[1:3], s[3:5]
	h, err := strconv.Atoi(hStr)
	if err != nil {
		return 0, err
	}

	m, err := strconv.Atoi(mStr)
	if err != nil {
		return 0, err
	}
	switch s[0] {
	case '+':
		return h*60*60 + m*60, nil
	case '-':
		return -h*60*60 - m*60, nil
	}
	return 0, errInvalidOffset
}
