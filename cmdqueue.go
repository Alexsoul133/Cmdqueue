package cmdqueue

import (
	"os"
	"os/exec"
	"strings"
)

var run [][]string

//Newcmd - push command to stack run
func Newcmd(s []string) {
	run = append(run, s)
}

//Start - run stack of cmd
func Start() {
	for i := range run {
		cmd := exec.Command("cmd", "/K", strings.Join(run[i], " "))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
