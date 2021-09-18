package models

import (
	"context"

	"example.com/ent_ex/ent"
	"example.com/ent_ex/ent/user"
)

func QueryUser(ctx context.Context, name string) (*ent.User, error) {
	return DB.User.
		Query().
		Where(user.Name(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
}
