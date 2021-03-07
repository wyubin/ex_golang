package main

type ak47 struct{}

type ak47Bullet struct {
	bullet
}

type ak47Base struct {
	base
}

func (a *ak47) makeBullet() iBullet {
	return &ak47Bullet{
		bullet: bullet{
			name:  "ak47",
			power: 12,
		},
	}
}

func (a *ak47) makeBase() iBase {
	return &ak47Base{
		base: base{
			name:  "ak47",
			power: 4,
		},
	}
}
