package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tom5521/MyGolangTools/commands"
)

var sh = commands.Sh{}
var (
	args                string = strings.Join(os.Args, ",")
	flatkpakUpgradesTxt string = `
--------------------------------------------------------------------
----------------------FLATPAK UPGRADES------------------------------
--------------------------------------------------------------------
`
	sysUpgradesTxt string = `
--------------------------------------------------------------------
----------------------SYSTEM UPGRADES-------------------------------
--------------------------------------------------------------------
`
	printversion string = `
----------------------Upgrade Tool V2-------------------------------`
	// Format Cmd
	noconfirm, assumeyes string
	sysUpdatecmd         string = "yay -Syu"
)

func updateflatpak() {
	fmt.Print(printversion, flatkpakUpgradesTxt)
	err := sh.Cmd(fmt.Sprintf("flatpak upgrade %v", assumeyes))
	if err != nil {
		updateflatpak()
	}
}
func updatesys() {
	fmt.Print(printversion, sysUpgradesTxt)
	err := sh.Cmd(fmt.Sprintf("%v %v", sysUpdatecmd, noconfirm))
	if err != nil {
		updatesys()
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("defaul mode")
		updatesys()
		updateflatpak()
		return
	}
	if strings.Contains(args, "noconfirm") {
		noconfirm = "--noconfirm"
		assumeyes = "--assumeyes"
	}
	if strings.Contains(args, "not-yay") {
		sysUpdatecmd = "sudo pacman -Syu"
	}
	if strings.Contains(args, "system") {
		updatesys()
		return
	} else if strings.Contains(args, "flatpak") {
		updateflatpak()
	} else {
		updatesys()
		updateflatpak()
		return
	}
}
