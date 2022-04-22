package main

type iGun interface {
	SetName(name string)
	SetPower(power int)
	Name() string
	Power() int
}
