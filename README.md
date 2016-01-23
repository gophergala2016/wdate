# wdate - print current time in several timezone

## Install:

    go get github.com/gophergala2016/wdate.git


## Usage:

It prints current time in local timezone:

    $ wdate
    Sat, 2016-01-23 19:20:06 KST +0900

You can **add** other timezone with name and offset:

    $ wdate add PST -0800
    $ wdate add SGT +0800

    $ wdate
    Sat, 2016-01-23 19:24:28 KST +0900
    Sat, 2016-01-23 02:24:28 PST -0800
    Sat, 2016-01-23 18:26:32 SGT +0800

Or **remove** timezone with name:

    $ wdate remove SGT

    $ wdate
    Sat, 2016-01-23 19:35:05 KST +0900
    Sat, 2016-01-23 02:35:05 PST -0800


## License

Copyright (c) 2016, Homin Lee.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.
