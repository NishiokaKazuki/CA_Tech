package queries

import (
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func InsertAppUser(engine *xorm.Engine, appUser tables.AppUsers) (bool, error) {
	affected, err := engine.Insert(&appUser)
	return err == nil && affected > 0, err
}

func InsertToken(engine *xorm.Engine, token tables.Tokens) (bool, error) {
	result, err := engine.Exec("insert into tokens(user_id, token) "+
		" values(?, ?) "+
		"on duplicate key update "+
		"token=?",
		token.UserId,
		token.Token,
		token.Token,
	)
	affected, _ := result.RowsAffected()

	return err == nil && affected > 0, err
}
