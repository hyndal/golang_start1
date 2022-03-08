package main

import (
	"fmt"
	"main/internal/game"
)

func main() {

	gameID1 := game.CreateGame(3, 3, 2)
	fmt.Println(gameID1)
	game1, err := game.GetGameStatus(gameID1)
	if err == nil {
		fmt.Println(game1)
	} else {
		fmt.Println(err)
	}
	//game.PrintGame(gameID1)
	game.AddPlayer(gameID1, 1, "Player1", false)
	game.AddPlayer(gameID1, 2, "Player2", false)

	game.StartGame(gameID1)
	err = game.MakeMove(gameID1, 1, 1, 1)
	if err != nil {
		fmt.Println(err)
	}
	//game.PrintGame(gameID1)

	// var team1, team2 int
	// fmt.Println("За какую команду вы будете играть? (1 - крестики, 2 - нолики)")
	// fmt.Scan(&team1)
	// if team1 != 1 && team1 != 2 {
	// 	fmt.Println("Вы выбрали неверную команду")
	// }
	// if team1 == 1 {
	// 	team2 = 2
	// } else {
	// 	team2 = 1
	// }
	// game := game.NewGame(size, team1, team2)

	// for {
	// 	for i := 1; i < size; i++ {
	// 		for {
	// 			var x, y int
	// 			fmt.Printf("Игрок %d - Введите координаты следующего хода\n", i)
	// 			fmt.Scan(&x, &y)
	// 			theEnd, moveMade, err := game.NewMove(i, x, y)
	// 			if theEnd {
	// 				game.PrintGame()
	// 				fmt.Println(err)
	// 				return
	// 			}
	// 			if moveMade {
	// 				game.PrintGame()
	// 				break
	// 			} else {
	// 				fmt.Println(err)
	// 			}
	// 		}
	// 	}
	// }
}
