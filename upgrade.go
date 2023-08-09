package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tom5521/MyGolangTools/commands"
)

var sh = commands.Sh{}

func main() {
	var (
		command                   string = "yay "
		only_system, only_flatpak bool
		par2, par3                string = "", ""
		vari                             = os.Args
		par                              = strings.Join(vari, " ")
	)
	if strings.Contains(par, "noconfirm") {
		par2 = "--noconfirm"
		par3 = "--assumeyes"
	}
	if strings.Contains(par, "not-yay") {
		command = "sudo pacman -Syu "
	}
	if strings.Contains(par, "system") {
		only_system = true
	}
	if strings.Contains(par, "flatpak") {
		only_flatpak = true
	}
	fmt.Println("----------------------Upgrade Tool V1.0.1-------------------------------")
	for !only_flatpak {
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("----------------------SYSTEM UPGRADES-------------------------------")
		fmt.Println("--------------------------------------------------------------------")
		err := sh.Cmd(command + par2)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			break
		}
	}
	for !only_system {
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("----------------------FLATPAK UPGRADES------------------------------")
		fmt.Println("--------------------------------------------------------------------")
		err := sh.Cmd("flatpak upgrade " + par3)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			break
		}
	}
}
