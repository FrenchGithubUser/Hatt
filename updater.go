package main

import (
	"fmt"
	"hatt/variables"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func (a *App) CheckForUpdate() selfupdate.Release {
	latest, found, err := selfupdate.DetectLatest("FrenchGithubUser/Hatt")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return selfupdate.Release{}
	}

	v := semver.MustParse(variables.CURRENT_VERSION)
	fmt.Println(latest, found, v)
	// if already on latest or if error
	if !found || latest.Version.LTE(v) {
		return selfupdate.Release{}
	}

	return *latest
}

func (a *App) SelfUpdate() bool {
	latest, _, _ := selfupdate.DetectLatest("FrenchGithubUser/Hatt")

	exe, err := os.Executable()
	if err != nil {
		log.Println("Could not locate executable path")
		return false
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		log.Println("Error occurred while updating binary:", err)
		return false
	}
	// log.Println("Successfully updated to version", latest.Version)

	return true
}
