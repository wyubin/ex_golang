package main

type iBase interface {
	getName() string
	getPower() int
}

type base struct {
	name  string
	power int
}

func (b *base) getName() string {
	return b.name
}

func (b *base) getPower() int {
	return b.power
}
