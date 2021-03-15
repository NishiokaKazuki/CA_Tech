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
