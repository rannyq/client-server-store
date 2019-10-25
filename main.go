package main

import (
	"fmt"
	"net/http"
)

const port = ":8888"

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Starting server in port %s\n", port)
	fmt.Fprintf(w, "Ranny did it!!!!\n")

	var str = r.Header.Get("")

	fmt.Fprintf(w, str)
}

func StoreData() {

	var storage Storage

	storage = &FileStorage{}

	contactinfo := ContactInfo{
		ID:     "1",
		Name:   "RannySue",
		Street: "123 Doheny",
		City:   "Dana Point",
		Zip:    "92629",
		Phone:  "3105555555",
	}
	storage.Add(contactinfo.ID, contactinfo)
}

func main() {

	var h handler

	err := http.ListenAndServe(port, h)

	if err != nil {
		panic(err)
	}
}
