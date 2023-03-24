package main

import (
	"log"
	"os/exec"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xwindow"
)

func main() {
	var focused *xwindow.Window
	var cx int = 0
	var cy int = 0

	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	focused = xwindow.New(X, X.RootWin())
	focused.Geometry()

	keybind.Initialize(X)
	keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			cmd := exec.Command("alacritty")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
		}).Connect(X, X.RootWin(), "Mod1-k", true)
	keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			focused = xwindow.New(X, e.Child)
			focused.Geometry()
			focused.Stack(xproto.StackModeAbove)
		}).Connect(X, X.RootWin(), "Mod1-r", true)

	buttonCallback := func(X *xgbutil.XUtil, e xevent.ButtonPressEvent) {
		focused = xwindow.New(X, e.Child)
		focused.Geometry()
		focused.Stack(xproto.StackModeAbove)
	}

	mousebind.Initialize(X)
	mousebind.ButtonPressFun(buttonCallback).Connect(X, X.RootWin(), "Mod1-1", false, true)
	mousebind.ButtonPressFun(buttonCallback).Connect(X, X.RootWin(), "Mod1-3", false, true)
	mousebind.ButtonPressFun(buttonCallback).Connect(X, X.RootWin(), "1", false, true)

	mousebind.Drag(X, focused.Id, focused.Id, "Mod1-1", true,
		func(X *xgbutil.XUtil, rx, ry, ex, ey int) (bool, xproto.Cursor) {
			cx = ex - focused.Geom.X()
			cy = ey - focused.Geom.Y()
			return true, 0
		},
		func(X *xgbutil.XUtil, rx, ry, ex, ey int) {
			focused.Move(ex-cx, ey-cy)
		},
		func(X *xgbutil.XUtil, rx, ry, ex, ey int) {})

	mousebind.Drag(X, focused.Id, focused.Id, "Mod1-3", true,
		func(X *xgbutil.XUtil, rx, ry, ex, ey int) (bool, xproto.Cursor) {
			cx = ex - focused.Geom.X()
			cy = ey - focused.Geom.Y()
			return true, 0
		},
		func(X *xgbutil.XUtil, rx, ry, ex, ey int) {
			focused.Resize(cx, cy)
			cx = ex - focused.Geom.X()
			cy = ey - focused.Geom.Y()
		},
		func(X *xgbutil.XUtil, rx, ry, ex, ey int) {})

	xevent.Main(X)
}
