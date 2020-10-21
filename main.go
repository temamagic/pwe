package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	str := strings.Join(os.Args, " ")
	str = strings.ToLower(str)

	s := Replace(str)

	fmt.Println("Що ти зробив:")

	args := strings.Split(s, " ")
	cmd := exec.Command("git", args[1:]...)
	path, _ := os.Getwd()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = path
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
