package main

import (
	"fmt"
	"main/internal/game"
	"main/internal/player"
)

func main() {

	game := game.NewGame(3)
	fmt.Println(game)

	var team1, team2 int
	fmt.Println("За какую команду вы будете играть? (1 - крестики, 2 - нолики)")
	fmt.Scan(&team1)
	if team1 != 1 && team1 != 2 {
		fmt.Println("Вы выбрали неверную команду")
	}
	if team1 == 1 {
		team2 = -1
	} else {
		team1 = -1
		team2 = 1
	}
	player1 := player.NewPlayer(team1)
	player2 := player.NewPlayer(team2)
	for {
		var x, y int
		fmt.Println("Введите координаты следующего хода")
		fmt.Scan(&x, &y)
		theEnd, err := game.NewMove(player1.GetPayerTeam(), x, y)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if theEnd {
			fmt.Println("Игра окончена!")
			break
		} else {
			theEnd := game.GenereteMove(player2.GetPayerTeam())
			if theEnd {
				fmt.Println("Игра окончена!")
				break
			}
		}
		fmt.Println(game.GetPlayingField())
	}
}
