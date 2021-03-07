package main

import "fmt"

func main() {
	ak47Factory, _ := getGunFactory("ak47")
	musketFactory, _ := getGunFactory("musket")

	ak47Bullet := ak47Factory.makeBullet()
	ak47Base := ak47Factory.makeBase()

	musketBullet := musketFactory.makeBullet()
	musketBase := musketFactory.makeBase()

	printBulletDetail(ak47Bullet)
	printBulletDetail(musketBullet)

	printBulletDetail(ak47Base)
	printBulletDetail(musketBase)

}

func printBulletDetail(s iBullet) {
	fmt.Printf("Name: %s\n", s.getName())
	fmt.Printf("Power: %d\n", s.getPower())
}

func printBaseDetail(s iBase) {
	fmt.Printf("Name: %s\n", s.getName())
	fmt.Printf("Power: %d\n", s.getPower())
}
