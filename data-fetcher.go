package main

import (
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// This function fetches the latest version
// from the cdn of open inventory
// and returns the version as string
func CheckForLatestVersion() string {
	resp, err := http.Get("https://cdn.open-inventory.org/latest")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	return string(body)
}

// This function loads the update script for the
// given version
func LoadUpdateScript(version string) []string {

	rt := runtime.GOOS

	resp, err := http.Get("https://cdn.open-inventory.org/update-scripts/" + rt + "/" + version)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	return strings.Split(string(body), "\n")
}
