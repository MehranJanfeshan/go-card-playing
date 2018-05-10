package main

import "fmt"

func main() {
	colorMap := map[string]string{
		"red":   "324234&sfe",
		"green": "342343%#234",
	}

	//color := make(map[string]string)
	//fmt.Print(color)
	//fmt.Println(colorMap)

	printMap(colorMap)
}

func printMap(mymap map[string]string) {
	for color, colorcode := range mymap {
		fmt.Print("Color: ", color, " Code: ", colorcode)
	}
}
