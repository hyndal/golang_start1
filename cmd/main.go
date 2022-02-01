package main

import "fmt"

func main() {

	var team int
	fmt.Println("За какую команду вы будете играть! (1 - крестики, 2 - нолики)")
	fmt.Scanf("%d\n", &team)
	if team != 1 && team != 2 {
		fmt.Println("Вы выбрали неверную команду")
	}
}
