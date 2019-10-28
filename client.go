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

const url = "http://localhost:7777"

func main() {

	//setup remote connection string
	remote := flag.String("remote", url, "Server location")

	flag.Parse()

	for i := 0; i < 10; i++ {

		id := i
		//create json struct
		contactinfo := ContactInfo.ContactInfo{
			id,
			"Joe Poe",
			"123 Doheny",
			"Dana Point",
			"92629",
			"3105555555",
		}

		b, err := json.Marshal(contactinfo)

		fmt.Printf("%s\n", b)

		//get current time
		start := time.Now()

		//post to URL with Json message
		resp, err := http.Post(*remote, "application/json", bytes2.NewBuffer(b))

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
}
