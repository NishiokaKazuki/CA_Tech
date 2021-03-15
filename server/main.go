package main

import (
	"CA_Tech/server/api"
	"CA_Tech/server/model/db"
)

const (
	port = "49400"
)

func main() {
	db := db.XormConnect()
	defer db.Close()

	finish := make(chan bool)

	go api.ListenAndServe(port)

	<-finish
}
