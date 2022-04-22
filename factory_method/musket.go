package main

type Musket struct {
	Gun
}

func NewMusket() iGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
