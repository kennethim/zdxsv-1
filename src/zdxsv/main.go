package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang/glog"

	"zdxsv/pkg/config"
	"zdxsv/pkg/db"
)

var cpu = flag.Int("cpu", 1, "setting GOMAXPROCS")
var profile = flag.Int("profile", 1, "0: no profile, 1: enable http pprof, 2: enable blocking profile")

func pprofPort(mode string) int {
	switch mode {
	case "lobby":
		return 16061
	case "battle":
		return 16062
	case "dns":
		return 16063
	case "login":
		return 16064
	case "status":
		return 16065
	default:
		return 16060
	}
}

func stripHost(addr string) string {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		glog.FatalDepth(1, "err in splitPort ", addr, err)
	}
	return ":" + fmt.Sprint(port)
}

func printUsage() {
	log.Println("Usage: ", os.Args[0], "[login, lobby, battle]")
}

func prepareDB() {
	conn, err := sql.Open("sqlite3", config.Conf.DB.Name)
	if err != nil {
		glog.Fatal(err)
	}
	db.DefaultDB = db.SQLiteDB{conn}
}

func prepareOption(command string) {
	runtime.GOMAXPROCS(*cpu)
	if *profile >= 1 {
		go func() {
			port := pprofPort(command)
			addr := fmt.Sprintf(":%v", port)
			log.Println(http.ListenAndServe(addr, nil))
		}()
	}
	if *profile >= 2 {
		runtime.MemProfileRate = 1
		runtime.SetBlockProfileRate(1)
	}
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	args := flag.Args()
	glog.Infoln(args, len(args))

	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	config.LoadConfig()

	command := args[0]
	prepareOption(command)

	switch command {
	case "dns":
		mainDNS()
	case "battle":
		mainBattle()
	case "lobby":
		prepareDB()
		mainLobby()
	case "login":
		prepareDB()
		mainLogin()
	case "status":
		mainStatus()
	case "initdb":
		os.Remove(config.Conf.DB.Name)
		prepareDB()
		db.DefaultDB.Init()
	default:
		printUsage()
		os.Exit(1)
	}
}
