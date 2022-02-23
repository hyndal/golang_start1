package main

import (
	"fmt"
	"main/internal/game"
)

const size int = 3

func main() {

	var team1, team2 int
	fmt.Println("За какую команду вы будете играть? (1 - крестики, 2 - нолики)")
	fmt.Scan(&team1)
	if team1 != 1 && team1 != 2 {
		fmt.Println("Вы выбрали неверную команду")
	}
	if team1 == 1 {
		team2 = 2
	} else {
		team2 = 1
	}
	game := game.NewGame(size, team1, team2)

	for {
		for i := 1; i < size; i++ {
			for {
				var x, y int
				fmt.Printf("Игрок %d - Введите координаты следующего хода\n", i)
				fmt.Scan(&x, &y)
				theEnd, moveMade, err := game.NewMove(i, x, y)
				if theEnd {
					game.PrintGame()
					fmt.Println(err)
					return
				}
				if moveMade {
					game.PrintGame()
					break
				} else {
					fmt.Println(err)
				}
			}
		}
	}
}
