package models

import (
	"context"

	"example.com/ent_ex/ent"
	"example.com/ent_ex/ent/user"
)

// QueryUsers: return user based name
func QueryUser(ctx context.Context, name string) (*ent.User, error) {
	return DB.User.
		Query().
		Where(user.Name(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
}

func CreateUser(ctx context.Context, name string, age *int) (*ent.User, error) {
	sqlDB := DB.User.Create().SetName(name)
	if age != nil {
		sqlDB = sqlDB.SetAge(*age)
	}
	return sqlDB.Save(ctx)
}
