package api

import (
	"encoding/json"
	"log"
	"net/http"
	"server/auth"
	"server/model/db"
	res "server/model/response"
	"server/model/tables"
	qr "server/queries"
	"server/utils"
)

func signIn(w http.ResponseWriter, r *http.Request) {
}

func createUser(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ReadRequestBody(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name, exist := body["name"]
	if exist != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := tables.AppUsers{
		Name: name,
	}
	_, err = qr.InsertAppUser(db.GetDBConnect(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err = qr.GetUserByName(db.GetDBConnect(), user.Name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Id == 0 {
		log.Println("users.name has " + user.Name + " was Not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token := auth.CreateToken(user)
	_, err = qr.InsertToken(db.GetDBConnect(), tables.Tokens{
		UserId: user.Id,
		Token:  token,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(res.Tokens{
		Token: token,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := qr.GetUserByToken(db.GetDBConnect(), token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(res.User{
		Name: user.Name,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	body, err := utils.ReadRequestBody(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name, exist := body["name"]
	if exist != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := qr.GetToken(db.GetDBConnect(), r.Header.Get("Authorization"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := tables.AppUsers{
		Id:   token.UserId,
		Name: name,
	}
	err = qr.UpdateUserById(db.GetDBConnect(), user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ListenAndServe(port string) {
	log.Println("starting server", port)

	http.HandleFunc("/signIn", signIn)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
