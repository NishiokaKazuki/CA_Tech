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
