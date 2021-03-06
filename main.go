package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ad-sho-loko/bogodb/db"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func showTitle() {
	title := `BogoDb : A toy database management system.`
	fmt.Println(title)
}

func client() {
	showTitle()
	stdin := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		stdin.Scan()
		q := stdin.Text()

		if strings.HasPrefix(q, "quit") {
			return
		}

		var err error
		if strings.HasPrefix(q, "exit") {
			_, err = http.Get("http://localhost:32198/exit")
		} else {
			escapedQry := "http://localhost:32198/execute?query=" + url.PathEscape(q)
			_, err = http.Get(escapedQry)
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}

func server() {
	bdb, err := db.NewBogoDb()
	if err != nil {
		log.Fatal(err)
	}
	bdb.Init()
	db.NewApiServer(bdb).Host()
}

var (
	serverMode = flag.Bool("server", false, "boot the db server")
)

func main() {
	flag.Parse()

	if *serverMode {
		server()
		return
	}

	client()
}
