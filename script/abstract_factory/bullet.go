package main

type iBullet interface {
	getName() string
	getPower() int
}

type bullet struct {
	name  string
	power int
}

func (b *bullet) getName() string {
	return b.name
}

func (b *bullet) getPower() int {
	return b.power
}
