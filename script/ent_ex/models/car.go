package models

import (
	"context"

	"example.com/ent_ex/ent"
	"example.com/ent_ex/ent/car"
)

func CreateCar(ctx context.Context, model string, plateNumber string) (*ent.Car, error) {
	sqlDB := DB.Car.Create().SetPlateNumber(plateNumber).SetModel(model)
	return sqlDB.Save(ctx)
}

// QueryCar: return car based plateNumber
func QueryCar(ctx context.Context, plateNumber string) (*ent.Car, error) {
	return DB.Car.
		Query().
		Where(car.PlateNumber(plateNumber)).First(ctx)
}
