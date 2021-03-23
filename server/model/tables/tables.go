package tables

type AppUsers struct {
	Id   uint64
	Name string
}

type Tokens struct {
	Id     uint64
	UserId uint64
	Token  string
}

type Characters struct {
	Id   uint64
	Name string
}

type CharactersIsInPossessions struct {
	UserId      uint64
	CharacterId uint64
	Quantity    uint8
}
