package main

import (
	"fmt"
	"os/exec"
)

func main() {

	latestVersion := CheckForLatestVersion()
	runningVersion := GetRunningVersion()

	if latestVersion == runningVersion {
		fmt.Println("You already have the latest version installed")
		return
	}

	script := LoadUpdateScript(latestVersion)

	for i := 0; i < len(script); i++ {
		cmd := exec.Command(script[i])
		err := cmd.Run()
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Successfully updated to", latestVersion)
}
