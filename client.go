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
const contentType = "application/json"

func sendData(contactinfo ContactInfo.ContactInfo, remote *string) {

	b, err := json.Marshal(contactinfo)

	fmt.Printf("%s\n", b)

	//get current time
	start := time.Now()

	//post to URL with Json message
	resp, err := http.Post(*remote, contentType, bytes2.NewBuffer(b))

	if err != nil {
		panic(err)
	}

	//make sure you close your resource
	defer resp.Body.Close()

	//read the response from server
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s responded with: %s %s\n", *remote, bytes, time.Since(start))
}

func main() {

	//setup remote connection string
	remote := flag.String("remote", url, "Server location")

	flag.Parse()

	//initiate slice of contacts
	contacts := make([]ContactInfo.ContactInfo, 20)

	//fill in the structs
	for i := range contacts {
		contacts[i].ID = i
		contacts[i].Name = "Joe Smoe"
		contacts[i].Street = "123 Doheny"
		contacts[i].City = "Dana Point"
		contacts[i].Zip = "92629"
		contacts[i].Phone = "3105555555"
		sendData(contacts[i], remote)
	}
}
