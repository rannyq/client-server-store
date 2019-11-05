package main

import (
	bytes2 "bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"ContactInfo"
)

const url = "http://localhost:7777"
const contentType = "application/json"
const extension = ".txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sendData(contactinfo ContactInfo.ContactInfo, remote *string) {

	b, err := json.Marshal(contactinfo)

	fmt.Printf("%s\n", b)

	//get current time
	start := time.Now()

	//post to URL with Json message
	resp, err := http.Post(*remote, contentType, bytes2.NewBuffer(b))

	check(err)

	//make sure you close your resource
	defer resp.Body.Close()

	//read the response from server
	bytes, err := ioutil.ReadAll(resp.Body)

	check(err)

	fmt.Printf("%s responded with: %s %s\n", *remote, bytes, time.Since(start))
}

func main() {

	//setup remote connection string
	remote := flag.String("remote", url, "Server location")

	flag.Parse()

	inputFiles, err := ioutil.ReadDir(".")

	check (err)

	fmt.Printf("Got this many files %s\n", len(inputFiles))

	//cycle through all files in directory
	for i := range inputFiles {

		//pull the names of all the files
		filename := inputFiles[i].Name()

		//check if its a file we are interested in
		if strings.Contains(filename, extension) && len(filename) > len(extension) {
			fmt.Printf("Found file %s %d\n", filename, len(filename))
			//Read the file
			byt, err := ioutil.ReadFile(filename)
			check(err)

			var contactinfo ContactInfo.ContactInfo

			//create the map from json string
			err2 := json.Unmarshal(byt, &contactinfo)
			check(err2)

			sendData(contactinfo, remote)
		} else{
			fmt.Printf("Not a file we want to process %s %d\n", filename, len(filename))
		}
	}
}
