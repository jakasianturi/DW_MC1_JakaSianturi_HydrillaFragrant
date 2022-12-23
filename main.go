package main

import "fmt"

func main() {

	profile := MakeProfile("Goku")
	fmt.Println("Name :", profile.Name)
	fmt.Println("Health :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)

}

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(karakter string) Profile {
	var profile = Profile{}
	if karakter == "Sasuke" {
		profile.Name = "Sasuke"
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
		return profile
	}
	if karakter == "Goku" {
		profile.Name = "Goku"
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
		return profile
	}
	if karakter == "Naruto" {
		profile.Name = "Naruto"
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
		return profile
	}
	return profile
}
func PowerUp(profile Profile, multipler int) Profile {
	profile.Health = profile.Health + (profile.Health * multipler)
	profile.Power = profile.Power + (profile.Power * multipler)
	profile.Exp = profile.Exp + (profile.Exp * multipler)
	return profile
}
