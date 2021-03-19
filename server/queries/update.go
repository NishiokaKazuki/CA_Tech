package queries

import (
	"server/model/tables"

	"github.com/go-xorm/xorm"
)

func UpdateUserById(engine *xorm.Engine, user tables.AppUsers) error {
	_, err := engine.Where(
		"id = ?",
		user.Id,
	).Update(&user)

	return err
}
