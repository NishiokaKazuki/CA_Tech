package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"server/auth"
	"server/model/db"
	res "server/model/response"
	"server/model/tables"
	qr "server/queries"
	"server/utils"
	"strconv"
	"time"
)

func signIn(w http.ResponseWriter, r *http.Request) {
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The requested method " + r.Method + " is not allowed for the " + r.Host + r.RequestURI))
		return
	}

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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The requested method " + r.Method + " is not allowed for the " + r.Host + r.RequestURI))
		return
	}

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
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The requested method " + r.Method + " is not allowed for the " + r.Host + r.RequestURI))
		return
	}

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

func GachaDraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The requested method " + r.Method + " is not allowed for the " + r.Host + r.RequestURI))
		return
	}
	var (
		totalPer    uint32
		currentPer  uint32
		groupResult []uint64
		characters  []tables.Characters
		result      res.GachaDrawResult
	)
	groupsCounts := make(map[uint64]uint64)

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

	body, err := utils.ReadRequestBody(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sTimes, exist := body["times"]
	if exist != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	times, err := strconv.Atoi(sTimes)

	gachaGroups, err := qr.FindGachaGroups(db.GetDBConnect())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(gachaGroups)

	for _, g := range gachaGroups {
		totalPer += g.Percentage
		c, err := qr.CountCharactersInGachaGroupsByGachaGroupId(db.GetDBConnect(), g.Id)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		groupsCounts[g.Id] = c
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times; i++ {
		currentPer = 0
		ran := rand.Intn(int(totalPer)) + 1
		for _, g := range gachaGroups {
			currentPer += g.Percentage
			if ran <= int(currentPer) {
				groupResult = append(groupResult, g.Id)
				break
			}
		}
	}

	for _, r := range groupResult {
		ran := rand.Intn(int(groupsCounts[r]))
		ch, err := qr.GetCharactersInGachaGroups(db.GetDBConnect(), r, uint64(ran))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		c, err := qr.GetCharacterbyId(db.GetDBConnect(), ch.CharacterId)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println(c)
		log.Println(user)

		_, err = qr.InsertOrUpDateCharactersIsInPossessions(db.GetDBConnect(), tables.CharactersIsInPossessions{
			UserId:      user.Id,
			CharacterId: c.Id,
			Quantity:    1,
		})
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		characters = append(characters, c)
	}

	for _, c := range characters {
		result.Results = append(result.Results, res.IResult{
			CharacterId: strconv.FormatUint(c.Id, 10),
			Name:        c.Name,
		})
	}

	buf, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

func CharacterList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The requested method " + r.Method + " is not allowed for the " + r.Host + r.RequestURI))
		return
	}

	var iCharacters []res.ICharacters

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

	characters, err := qr.FindCharactersByUserId(db.GetDBConnect(), user.Id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sUserId := strconv.FormatUint(user.Id, 10)
	for _, c := range characters {
		iCharacters = append(iCharacters, res.ICharacters{
			UserCharacterId: sUserId,
			CharacterId:     strconv.FormatUint(c.Id, 10),
			Name:            c.Name,
		})
	}

	buf, err := json.Marshal(res.CharacterList{
		Characters: iCharacters,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

func ListenAndServe(port string) {
	log.Println("starting server", port)

	http.HandleFunc("/signIn", signIn)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)
	http.HandleFunc("/gacha/draw", GachaDraw)
	http.HandleFunc("/character/list", CharacterList)

	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
