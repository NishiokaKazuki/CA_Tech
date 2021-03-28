package queries

import (
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func FindCharactersByUserId(engine *xorm.Engine, userId uint64) ([]tables.Characters, error) {
	var characters []tables.Characters
	err := engine.Alias("c").Join(
		"INNER",
		[]string{"characters_is_in_possessions", "p"},
		"p.user_id = ? AND p.quantity >= 1",
		userId,
	).Where(
		"c.id = p.character_id",
	).Find(&characters)

	return characters, err
}

func FindGachaGroups(engine *xorm.Engine) ([]tables.GachaGroups, error) {
	var gachaGroups []tables.GachaGroups
	err := engine.Find(&gachaGroups)

	return gachaGroups, err
}

func FindCharactersInGachaGroups(engine *xorm.Engine, gachaGroupIds uint64) ([]tables.CharactersInGachaGroups, error) {
	var groups []tables.CharactersInGachaGroups
	err := engine.In(
		"gacha_group_id",
		gachaGroupIds,
	).Find(&groups)
	return groups, err
}
