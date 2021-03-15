package main

import (
	"server/api"
	"server/model/db"
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
