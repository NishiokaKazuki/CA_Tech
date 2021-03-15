package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadRequestBody(r *http.Request) (map[string]string, error) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	m := map[string]string{}
	err = json.Unmarshal(buf, &m)
	return m, err
}

// func CreateReaponseToken(token res.Tokens) ([]byte, err) {
// 	buf, err := json.Marshal(token)
// }
