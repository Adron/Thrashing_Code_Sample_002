package main

import "fmt"

func main() {
	var someNumbers [5]int
	fmt.Println(someNumbers[0])
	fmt.Println(someNumbers[len(someNumbers)-1])

	someNumbers[0] = 42
	someNumbers[3] = 69
	someNumbers[4] = 666

	for i, v := range someNumbers{
		fmt.Printf("Index: %d Value: %d\n", i, v)
	}

	someStoryTelling := [3]string{"This is a great story", "Maybe not a great story", "Naw it is the best story ever!"}

	for i, v := range someStoryTelling{
		fmt.Printf("Index: %d Value: %s\n", i, v)
	}

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		BYN
		BAM
		KHR
		ISK
		IRR
	)
	symbol := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		BYN: "Br",
		BAM: "Br",
		KHR:"៛",
		ISK:"kr",
		IRR: "﷼",
	}

	for i, v := range symbol{
		fmt.Printf("Index: %d Symbol: %s\n", i, v)
	}
}
