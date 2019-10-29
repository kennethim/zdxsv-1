package main

import (
	"log"
	"os"
	"time"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

var (
	// These variables are automatically assigned during release process.
	// `-s -w -X main.releaseVersion ={{.Version}} -X main.releaseCommit={{.ShortCommit}} -X main.releaseDate={{.Date}}
	releaseVersion = "local"
	releaseCommit  = "local"
	releaseDate    = "local"
)

func printReleaseInfo() {
	log.Println("releaseVersion", releaseVersion)
	log.Println("releaseCommit", releaseCommit)
	log.Println("releaseDate", releaseDate)
}

func doSelfUpdate() {
	if releaseVersion == "local" {
		return
	}
	log.Println("開始檢查更新")

	latest, found, err := selfupdate.DetectLatest("kennethim/udp")
	if err != nil {
		log.Println("檢查更新失敗", err)
	}

	log.Println("最新版>", latest.Version)

	v := semver.MustParse(releaseVersion)
	if !found || latest.Version.LTE(v) {
		log.Println("目前已經是最新版本")
		return
	}

	exe, err := os.Executable()
	if err != nil {
		log.Println("檢查更新失敗", err)
		return
	}

	log.Println("更新中...")
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		log.Println("更新失敗", err)
		return
	}

	log.Println("更新成功")
	log.Println("請重新開啟本程式")
	time.Sleep(time.Second)
	os.Exit(0)
}
