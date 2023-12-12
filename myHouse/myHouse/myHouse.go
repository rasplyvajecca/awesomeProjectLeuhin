package myHouse

import (
	"awesomeProjectLeuhin/myHouse/cats"
	"awesomeProjectLeuhin/myHouse/children"
	"awesomeProjectLeuhin/myHouse/dogs"
	"awesomeProjectLeuhin/myHouse/members"
	"awesomeProjectLeuhin/myHouse/things"
	"fmt"
)

type House struct {
	FirstSize      float64
	SecondSize     float64
	ThirdSize      float64
	CreateMembers  []members.Members
	CreateChildren []children.Children
	CreateDogs     []dogs.Dogs
	CreateCats     []cats.Cats
	CreateThings   []things.Things
}

func HouseObjects() House {
	return House{
		FirstSize:      3.333,
		SecondSize:     3.333,
		ThirdSize:      3.333,
		CreateMembers:  members.CreateObjects(),
		CreateChildren: children.CreateObjects(),
		CreateDogs:     dogs.CreateObjects(),
		CreateCats:     cats.CreateObjects(),
		CreateThings:   things.CreateObjects(),
	}
}

func PrintObjects(house House) {
	fmt.Printf("Мой дом:\n")
	fmt.Printf("Размеры: %.2f м, %.2f м, %.2f м\n", house.FirstSize, house.SecondSize, house.ThirdSize)

	printMembers(house.CreateMembers)
	printChildren(house.CreateChildren)
	printDogs(house.CreateDogs)
	printCats(house.CreateCats)
	printThings(house.CreateThings)
}

func printMembers(Info []members.Members) {
	fmt.Println("Родители:")
	for _, obj := range Info {
		fmt.Printf("- %s: %s %s\n", obj.Member, obj.Name, obj.Surname)
	}
}

func printChildren(Info []children.Children) {
	fmt.Println("Дети:")
	for _, obj := range Info {
		fmt.Printf("- %s, %d\n", obj.Name, obj.Age)
	}
}

func printDogs(Info []dogs.Dogs) {
	fmt.Println("Собаки:")
	for _, obj := range Info {
		fmt.Printf("- %s, %d\n", obj.Name, obj.Age)
	}
}

func printCats(Info []cats.Cats) {
	fmt.Println("Кошки:")
	for _, obj := range Info {
		fmt.Printf("- %s, %d\n", obj.Name, obj.Age)
	}
}

func printThings(Info []things.Things) {
	fmt.Println("Вещи:")
	for _, obj := range Info {
		fmt.Printf("- %s: %0.2f, %d\n", obj.Name, obj.Size, obj.Count)
	}
}
