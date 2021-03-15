package api

import (
	"log"
	"net/http"
)

func signIn(w http.ResponseWriter, r *http.Request) {
}

func createUser(w http.ResponseWriter, r *http.Request) {
}

func getUser(w http.ResponseWriter, r *http.Request) {
}

func updateUser(w http.ResponseWriter, r *http.Request) {
}

func ListenAndServe(port string) {
	http.HandleFunc("/signIn", signIn)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("starting server", port)
}
