package main

import "fmt"

func NewGun(gunType string) (iGun, error) {
	if gunType == "AK47" {
		return NewAK47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("Invalid gun type")
}
