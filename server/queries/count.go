package queries

import (
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func CountCharactersInGachaGroupsByGachaGroupId(engine *xorm.Engine, groupId uint64) (uint64, error) {
	var cGroups tables.CharactersInGachaGroups
	res, err := engine.Where(
		"gacha_group_id = ?",
		groupId,
	).Count(cGroups)

	return uint64(res), err
}
