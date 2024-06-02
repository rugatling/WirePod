package main

import (
	"github.com/rugatling/WirePod/cross/podapp"
	cross_win "github.com/rugatling/WirePod/cross/win"
)

func main() {
	podapp.StartWirePod(cross_win.NewWindows())
}
