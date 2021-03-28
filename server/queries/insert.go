package queries

import (
	"log"
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

// "SELECT "+
// "?, ?, 1 "+
// "WHERE "+
// "NOT EXISTS ( "+
// "SELECT TOP 1 1 FROM characters_is_in_possessions WHERE "+
// "user_id = ? AND character_id = ? "+
// ")",

func InsertCharactersIsInPossessions(engine *xorm.Engine, possession tables.CharactersIsInPossessions) (bool, error) {
	result, err := engine.Exec(
		"UPDATE characters_is_in_possessions SET "+
			"user_id = ?, character_id = ?, quantity = 1 "+
			"WHERE "+
			"user_id = ? AND character_id = ? "+
			"IF @@ROWCOUNT = 0 "+
			"INSERT INTO characters_is_in_possessions ( "+
			"user_id, character_id, quantity "+
			") "+
			"VALUES( "+
			"?, ?, 1"+
			")",
		possession.UserId,
		possession.CharacterId,
		possession.UserId,
		possession.CharacterId,
		possession.UserId,
		possession.CharacterId,
	)
	log.Println(err)
	affected, err := result.RowsAffected()

	return err == nil && affected > 0, err
}
