package main

import "fmt"

func main() {
	swedishMetalBands := [...]string{0: "Entombed", 1: "In Flames", 2: "Opeth", 3: "Bathory"}

	for _, band := range swedishMetalBands {
		fmt.Println(band)
	}

	myFavorites := swedishMetalBands[1:3]
	yourFavorites := swedishMetalBands[2:4]

	fmt.Println("\nMy Favorites:")

	for _, band := range myFavorites {
		fmt.Println(band)
	}

	fmt.Println("Your Favorites:")

	for _, band := range yourFavorites {
		fmt.Println(band)
	}

	for _, myFav := range myFavorites {
		for _, yourFav := range yourFavorites {
			if myFav == yourFav {
				fmt.Printf("\nHoly beezulbub we diverted a war! You and I both like %s.\n", myFav)
			}
		}
	}

	allSwedishBandsAllTheTime := swedishMetalBands[0:4]
	allSwedishBandsAllTheTime = append(allSwedishBandsAllTheTime, "Soilwork", "Hammerfall", "Dark Tranquility", "At The Gates", "Amon Amarth", "Evergrey", "Arch Enemy", "Candlemass")

	for _, band := range allSwedishBandsAllTheTime {
		fmt.Println(band)
	}
}
