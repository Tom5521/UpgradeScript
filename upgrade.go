package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var cmd, cmd2 int

func sys(command string) int {

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		return 1
	}
	return 0
}

func main() {
	var par2, par3 string = "", ""
	vari := os.Args
	par := strings.Join(vari, " ")
	if strings.Contains(par, "noconfirm") {
		par2 = "--noconfirm"
		par3 = "--assumeyes"
	}
	fmt.Println("----------------------Upgrade Tool V1-------------------------------")
	for {
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("----------------------SYSTEM UPGRADES-------------------------------")
		fmt.Println("--------------------------------------------------------------------")

		cmd = sys("yay " + par2)
		if cmd == 1 {
			fmt.Println("ERROR")
		} else {
			break
		}
	}
	for {
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("----------------------FLATPAK UPGRADES------------------------------")
		fmt.Println("--------------------------------------------------------------------")
		cmd2 = sys("flatpak upgrade " + par3)
		if cmd2 == 1 {
			fmt.Println("ERROR")
		} else {
			break
		}
	}
}
