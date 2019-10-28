package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	remote := flag.String("remote", "http://localhost:7777", "Server location")

	flag.Parse()
	start := time.Now()

	resp, err := http.Get(*remote)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s said: %s %s\n", *remote, bytes, time.Since(start))
}
