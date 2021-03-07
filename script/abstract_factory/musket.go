package main

type musket struct{}

type musketBullet struct {
	bullet
}

type musketBase struct {
	base
}

func (a *musket) makeBullet() iBullet {
	return &ak47Bullet{
		bullet: bullet{
			name:  "musket",
			power: 16,
		},
	}
}

func (a *musket) makeBase() iBase {
	return &ak47Base{
		base: base{
			name:  "musket",
			power: 2,
		},
	}
}
