package main

type AK47 struct {
	Gun
}

func NewAK47() iGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
