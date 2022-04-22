package main

import "fmt"

func main() {
	ak47, _ := NewGun("AK47")
	musket, _ := NewGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.Name())
	fmt.Printf("Power: %d\n", g.Power())
}
