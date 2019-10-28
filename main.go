package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ContactInfo"
)

const port = ":7777"

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Starting server in port %s\n", port)
	fmt.Fprintf(w, "Ranny did it!!!!\n")

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Got Json: %s\n", b)

	defer r.Body.Close()

	var ci ContactInfo.ContactInfo

	err2 := json.Unmarshal(b, &ci)

	if err2 != nil {
		panic(err2)
	}

	StoreData()

	fmt.Fprintf(w, "Wrote File and it took")
}

func StoreData() {

	contactinfo := ContactInfo.ContactInfo{
		"1",
		"Joe Smoe",
		"123 Doheny",
		"Dana Point",
		"92629",
		"3105555555",
	}

	var storage ContactInfo.Storage

	storage = &ContactInfo.FileStorage{}

	storage.WriteFile(contactinfo.ID, contactinfo)
}

func main() {

	var h handler

	err := http.ListenAndServe(port, h)

	if err != nil {
		panic(err)
	}
}
