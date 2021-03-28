package queries

import (
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func GetUserByName(engine *xorm.Engine, name string) (tables.AppUsers, error) {
	var user tables.AppUsers

	_, err := engine.Where(
		"name = ? and disabled = false",
		name,
	).Get(&user)

	return user, err
}

func GetUserByToken(engine *xorm.Engine, token string) (tables.AppUsers, error) {
	var user tables.AppUsers

	_, err := engine.Alias("u").Join(
		"INNER",
		[]string{"tokens", "t"},
		"t.token = ?",
		token,
	).Where(
		"u.id = t.user_id and disabled = false",
	).Get(&user)

	return user, err
}

func GetToken(engine *xorm.Engine, token string) (tables.Tokens, error) {
	var tokens tables.Tokens

	_, err := engine.Where(
		"token = ?",
		token,
	).Get(&tokens)

	return tokens, err
}

func GetCharactersInGachaGroups(engine *xorm.Engine, gachaGroupId uint64, ran uint64) (tables.CharactersInGachaGroups, error) {
	var group tables.CharactersInGachaGroups
	_, err := engine.Where(
		"gacha_group_id = ?",
		gachaGroupId,
	).Limit(1, int(ran)).Get(&group)
	return group, err
}

func GetCharacterbyId(engine *xorm.Engine, id uint64) (tables.Characters, error) {
	var character tables.Characters

	_, err := engine.Where(
		"id = ?",
		id,
	).Get(&character)

	return character, err
}
