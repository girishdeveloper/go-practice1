package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen() {
	slc := []string{"clear"}
	if runtime.GOOS == "windows" {
		slc = []string{"cmd", "/c", "cls"}
	}
	cmd := exec.Command(strings.Join(slc, " "))
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	ClearScreen()
}
