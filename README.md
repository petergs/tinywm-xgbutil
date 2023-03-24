# tinywm-xgbutil
> A port of [tinywm](https://github.com/mackstann/tinywm) using [xgbutil](https://github.com/BurntSushi/xgbutil)

## Setup
Clone this repo, `cd` into the directory, and run `go get ./...` then `go build`.

Use any option you'd like for X [session configuration](https://wiki.archlinux.org/title/display_manager#Session_configuration). Resource (2) below has an example of creating a `.desktop` file for your display manager. In any case, **make sure you launch a terminal**.

Alternatively, you can test it out in [Xephyr](https://www.freedesktop.org/wiki/Software/Xephyr/) with something like this:

```
#!/bin/bash
Xephyr :1 -ac -softCursor -screen 1400x1000 &
sleep 1
DISPLAY=:1 ./tinywm-xgbutil &
DISPLAY=:1 xterm &
```

## Usage
- Alt+Button1+Drag - move window
- Alt+Button3+Drag - resize window
- Alt+R or Button1 - raise and focus the window under the cursor

## Resources
1. the original TinyWM on mackstann's [Github]() or [here](http://incise.org/tinywm.html)
2. [collinglass](https://github.com/collinglass/tinywm) ported TinyWM using `cgo` 