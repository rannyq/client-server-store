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

	fmt.Printf("Got a http request on port %s\n", port)

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if len(b) > 0 {

		fmt.Printf("Got Json: %s\n", b)

		defer r.Body.Close()

		var ci ContactInfo.ContactInfo

		err2 := json.Unmarshal(b, &ci)

		if err2 != nil {
			panic(err2)
		}

		StoreData(ci)

		fmt.Fprintf(w, "Wrote File and it took")
	} else {
		fmt.Fprintf(w, "No Data")
	}
}

func StoreData(contactinfo ContactInfo.ContactInfo) {

	storage := &ContactInfo.FileStorage{}

	storage.WriteFile(fmt.Sprint(contactinfo.ID), contactinfo)
}

func main() {

	var h handler

	fmt.Printf("Starting server in port %s\n", port)

	err := http.ListenAndServe(port, h)

	if err != nil {
		panic(err)
	}
}
