package main

import "os/exec"

func main() {
	// after making start-pbay-backend.sh executable
	runAll := "./start-pbay-backend.sh"
	exec.Command(runAll)

	return
}
