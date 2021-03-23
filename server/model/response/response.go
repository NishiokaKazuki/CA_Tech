package response

type Tokens struct {
	Token string `json:"token"`
}

type User struct {
	Name string `json:"name"`
}

type CharacterList struct {
	Characters []ICharacters `json:"characters"`
}

type ICharacters struct {
	UserCharacterId string `json:"userCharacterID"`
	CharacterId     string `json:"characterID"`
	Name            string `json:"name"`
}
