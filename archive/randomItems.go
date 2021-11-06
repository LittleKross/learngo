package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func genItem() string {
	return genFormat(randomType())
}
func randomType() string {
	types := []string{
		"armor",
		"item",
		"weapon",
	}
	return types[rand.Intn(len(types))]
}
func randomSubtype(object string) string {
	result := ""
	subTypeArmor := []string{
		"Gauntlets",
		"Boots",
		"Helmet",
		"Breastplate",
		"Leggings",
	}
	subTypeItem := []string{
		"Compass",
		"Datapad",
		"Map",
		"Backpack",
		"Food",
		"Rope",
	}
	subTypeWeapon := []string{
		"Sword",
		"Knife",
		"Flail",
		"Mace",
		"Pike",
		"Spear",
		"Warhammer",
		"Crossbow",
		"Bow",
		"Catapult",
		"Trebuchet",
		"Cannon",
		"Stick",
		"Pistol",
		"Rifle",
		"Rocket launcher",
		"Grenade launcher",
		"Blaster",
		"Repeater",
		"Energy Sword",
		"Beam rifle",
		"Lightsaber",
		"Lightwhip",
		"Gravity hammer",
		"Flamethrower",
		"Vibroblade",
		"Nuke",
	}
	switch{
	case object == "armor": result = subTypeArmor[rand.Intn(len(subTypeArmor))]
	case object == "item": result = subTypeItem[rand.Intn(len(subTypeArmor))]
	case object == "weapon": result = subTypeWeapon[rand.Intn(len(subTypeWeapon))]
	}
	return result
}
func randomManufacturer() string {
	manufacturers := []string{
		"Spiff Co.",
		"Monsters Inc.",
		"Kuat Drive Yards",
		"Puma Tires",
		"Banking Clan",
		"Baktoid Inovations",
		"Hamanfest",
		"Bowser Oil",
		"Aperature Laberatories",
		"Black Mesa",
		"Nuka Cola",
		"Vault-Tec",
		"ACME",
	}
	return manufacturers[rand.Intn(len(manufacturers))]
}
func genFormat(object string) string {
	result := ""
	switch{
	case object == "armor": result = "Armor is a " + randomSubtype(object) + " made by " + randomManufacturer() //add more info here
	case object == "item": result = "Item is a " + randomSubtype(object) + " made by " + randomManufacturer() //add more info here
	case object == "weapon": result = "Weapon is a " + randomSubtype(object) + " made by " + randomManufacturer() //add more info here
	}
	return result
}

func main() {
	fmt.Println(genItem())
}
