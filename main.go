package main

import (
		"fmt"
        "net/http"
        "ContactInfo"
)

const port = ":7777"

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Starting server in port %s\n", port)
	fmt.Fprintf(w, "Ranny did it!!!!\n")

	var str = r.Header.Get("")

	StoreData()

	fmt.Fprintf(w, "Stored!!!!\n")

	fmt.Fprintf(w, str)
}

func StoreData() {

	var storage ContactInfo.Storage

	storage = &ContactInfo.FileStorage{}

	contactinfo := ContactInfo.ContactInfo{
		ID:     "1",
		Name:   "Joe Smoe",
		Street: "123 Doheny",
		City:   "Dana Point",
		Zip:    "92629",
		Phone:  "3105555555",
	}
	storage.WriteFile(contactinfo.ID, contactinfo)
}

func main() {

	var h handler

	err := http.ListenAndServe(port, h)

	if err != nil {
		panic(err)
	}
}
