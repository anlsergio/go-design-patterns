package main

type Gun struct {
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) Name() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) Power() int {
	return g.power
}
