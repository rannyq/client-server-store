package main

import (
	bytes2 "bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"ContactInfo"
)

func main() {

	//create json struct
	contactinfo := ContactInfo.ContactInfo{
		"1",
		"Joe Poe",
		"123 Doheny",
		"Dana Point",
		"92629",
		"3105555555",
	}

	b, err := json.Marshal(contactinfo)

	fmt.Printf("%s\n", b)

	//setup remote connection string
	remote := flag.String("remote", "http://localhost:7777", "Server location")

	flag.Parse()

	//get current time
	start := time.Now()

	//post to URL with Json message
	resp, err := http.Post(*remote, "application/json", bytes2.NewBuffer(b))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//read response from server
	resp2, err := http.Get(*remote)

	if err != nil {
		panic(err)
	}

	defer resp2.Body.Close()

	bytes, err := ioutil.ReadAll(resp2.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s said: %s %s\n", *remote, bytes, time.Since(start))
}
