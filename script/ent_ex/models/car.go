package models

import (
	"context"

	"example.com/ent_ex/ent"
)

func CreateCar(ctx context.Context, model string) (*ent.Car, error) {
	sqlDB := DB.Car.Create().SetModel(model)
	return sqlDB.Save(ctx)
}
