package main

import "fmt"

type iGunFactory interface {
	makeBullet() iBullet
	makeBase() iBase
}

func getGunFactory(name string) (iGunFactory, error) {
	if name == "ak47" {
		return &ak47{}, nil
	}
	if name == "musket" {
		return &musket{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
