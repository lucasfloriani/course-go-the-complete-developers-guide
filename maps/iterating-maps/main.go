package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#00FF00",
		"white": "#FFFFFF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
