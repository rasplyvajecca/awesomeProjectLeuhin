package main

import (
	"awesomeProjectLeuhin/myHouse/myHouse"
	"fmt"
)

func main() {
	fmt.Println("Дом строится")
	houseObjects := myHouse.HouseObjects()
	myHouse.PrintObjects(houseObjects)
}
