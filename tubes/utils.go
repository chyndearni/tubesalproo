package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pause() {
	fmt.Print("Tekan ENTER untuk kembali ke menu...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
