package main

import (
	cross_mac "github.com/rugatling/WirePod/cross/mac"
	"github.com/rugatling/WirePod/cross/podapp"
)

func main() {
	podapp.StartWirePod(cross_mac.NewMacOS())
}
