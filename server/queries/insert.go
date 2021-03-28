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

func InsertOrUpDateCharactersIsInPossessions(engine *xorm.Engine, possession tables.CharactersIsInPossessions) (bool, error) {
	var p tables.CharactersIsInPossessions
	_, err := engine.Where(
		"user_id = ? AND character_id = ?",
		possession.UserId,
		possession.CharacterId,
	).Get(&p)
	if err != nil {
		return false, err
	}

	if p.UserId == 0 {
		_, err = engine.Insert(&possession)
	} else {
		result, _ := engine.Exec(
			"UPDATE characters_is_in_possessions SET "+
				"user_id = ?, character_id = ?, quantity = quantity + 1 "+
				"WHERE "+
				"user_id = ? AND character_id = ? ",
			possession.UserId,
			possession.CharacterId,
			possession.UserId,
			possession.CharacterId,
		)
		_, err = result.RowsAffected()
	}
	return err == nil, err
}
