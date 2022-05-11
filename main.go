package main

import (
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := []string{"--http-header-fields=Origin:https://ani.gamer.com.tw", "--title=", "--sub-file=", ""}
	arg := os.Args[1][6:]
	arg, err := url.QueryUnescape(arg)
	args[1] += strings.Split(arg, ";")[0]
	args[2] += strings.Split(arg, ";")[1]
	args[3] = strings.Split(arg, ";")[2]
	if err != nil {
		log.Fatal(err)
		return
	}
	exec.Command("mpv.exe", args...).Start()
}
