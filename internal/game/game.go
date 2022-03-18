package game

import (
	"errors"
	"fmt"
)

type position struct {
	x   int `json:"x"`
	y   int `json:"y"`
	plr int `json:"plr"`
}
type player struct {
	ID       int
	name     string
	computer bool
	init     bool
}

type game struct {
	ID           int
	size         int
	winSize      int
	gameStatus   int
	moves        int
	nextPlayer   int
	playingField [][]int
	players      []player
}

var games []game

type GamesStruct struct {
	list []game
}

func (g *game) genereteMove(playerTeam int) (int, int) {

	var retX, retY int

	for x := range g.playingField {
		for y := range g.playingField[x] {
			if g.playingField[x][y] == 0 {
				return x, y
			}
		}
	}

	return retX, retY
}

func (gs *GamesStruct) CreateGame(size int, winSize int, players int) int {

	var gameID int = len(gs.list) + 1
	newGame := game{gameID, size, winSize, 0, 0, 0, make([][]int, size), make([]player, players)}
	for i := range newGame.playingField {
		newGame.playingField[i] = make([]int, size)
	}
	gs.list = append(gs.list, newGame)

	return gameID
}

func GetGameStatus(gameID int) (game, error) {

	if gameID > len(games) || gameID <= 0 {
		return game{}, errors.New("неверный идентификатор игры")
	}

	return games[gameID-1], nil

}

func AddPlayer(gameID int, playerID int, playerName string, computer bool) error {

	if gameID > len(games) || gameID <= 0 {
		return errors.New("неверный идентификатор игры")
	}
	g := games[gameID-1]
	if g.gameStatus != 0 {
		return errors.New("игра уже началась")
	}
	if playerID > len(g.players) {
		return errors.New("неверный идентификатор игрока")
	}
	g.players[playerID-1].name = playerName
	g.players[playerID-1].computer = computer
	g.players[playerID-1].init = true

	return nil

}

func StartGame(gameID int) error {

	fmt.Println(games)

	if gameID > len(games) || gameID <= 0 {
		return errors.New("неверный идентификатор игры")
	}
	g := games[gameID-1]
	if g.gameStatus != 0 {
		return errors.New("игра уже началась")
	}
	for i := range g.players {
		if !g.players[i].init {
			return errors.New("не все игроки добавлены")
		}
	}
	g.gameStatus = 1
	g.nextPlayer = 1

	return nil

}

func MakeMove(gameID int, playerID int, x int, y int) error {

	if gameID > len(games) || gameID <= 0 {
		return errors.New("неверный идентификатор игры")
	}
	g := games[gameID-1]
	if g.gameStatus == 1 {
		return errors.New("игра еще не началась")
	}
	if playerID > len(g.players) || playerID <= 0 {
		return errors.New("неверный идентификатор игрока")
	}
	if playerID != g.nextPlayer {
		return fmt.Errorf("ожидается ход от игрока %d", g.nextPlayer)
	}
	if x > g.size || x < 0 || y > g.size || y < 0 {
		return errors.New("ячейка находится за границами поля")
	}
	if g.playingField[x][y] != 0 {
		return errors.New("ячейка уже занята")
	}
	g.playingField[x][y] = playerID
	g.nextPlayer++
	if g.nextPlayer > len(g.players) {
		g.nextPlayer = 1
	}
	g.moves++
	if g.endGame(x, y, playerID) {
		g.gameStatus = 2
		return errors.New("вы выиграли")
	}
	if g.moves == g.size*g.size {
		g.gameStatus = 2
		return errors.New("ходов больше нет")
	}

	return nil
}

func (g game) endGame(x int, y int, team int) bool {

	minX := x - g.winSize + 1
	maxX := x + g.winSize
	minY := y - g.winSize + 1
	maxY := y + g.winSize

	qLength1 := 0
	qLength2 := 0
	j := minY
	for i := minX; i < maxX; i++ {
		if i >= 0 && i < g.size {
			if g.playingField[i][y] == team {
				qLength1++
			} else {
				qLength1 = 0
			}
		}
		if i >= 0 && j >= 0 && i < g.size && j < g.size {
			if g.playingField[i][j] == team {
				qLength2++
			} else {
				qLength2 = 0
			}
		}
		if qLength1 == g.winSize || qLength2 == g.winSize {
			return true
		}
		j++
	}

	qLength1 = 0
	qLength2 = 0
	j = maxX - 1
	for i := minY; i < maxY; i++ {
		if i >= 0 && i < g.size {
			if g.playingField[x][i] == team {
				qLength1++
			} else {
				qLength1 = 0
			}
		}
		if i >= 0 && j >= 0 && i < g.size && j < g.size {
			if g.playingField[j][i] == team {
				qLength2++
			} else {
				qLength2 = 0
			}
		}
		if qLength1 == g.winSize || qLength2 == g.winSize {
			return true
		}
		j--
	}

	return false
}

func (gs GamesStruct) PrintGame(gameID int) (interface{}, error) {
	if len(gs.list) == 0 {
		return nil, errors.New("неверный идентификатор игры")
	}
	g := gs.list[gameID-1]
	var positions = make([]position, g.size*g.size)
	var i int = 0
	for x := range g.playingField {
		for y := range g.playingField[x] {
			positions[i].x = x
			positions[i].y = y
			positions[i].plr = g.playingField[x][y]
			i++
		}
	}

	return positions, nil
}
