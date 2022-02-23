package game

import (
	"errors"
	"fmt"
)

const winSize int = 3

type game struct {
	size         int
	moves        int
	lastMove     int
	playingField [][]int
	playersMap   map[int]bool
}

func NewGame(size int, team1 int, team2 int) game {

	newGame := game{size, 0, 0, make([][]int, size), make(map[int]bool)}
	for i := range newGame.playingField {
		newGame.playingField[i] = make([]int, size)
	}
	newGame.playersMap[team1] = false
	newGame.playersMap[team2] = true

	return newGame
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
func (g *game) NewMove(playerTeam int, x int, y int) (bool, bool, error) {
	if x > g.size || x < 0 || y > g.size || y < 0 {
		return false, false, errors.New("ячейка находится за границами поля")
	}
	if g.playingField[x][y] != 0 {
		return false, false, errors.New("ячейка уже занята")
	}
	if g.lastMove == playerTeam {
		return false, false, errors.New("этот игрок только что ходил")
	}
	playerComputer, ok := g.playersMap[playerTeam]
	if !ok {
		return false, false, errors.New("такого игрока нет")
	}
	if playerComputer {
		x, y = g.genereteMove(playerTeam)
	}
	g.playingField[x][y] = playerTeam
	g.lastMove = playerTeam
	g.moves++
	if g.endGame(x, y, playerTeam) {
		return true, true, errors.New("вы выиграли")
	}
	if g.moves == g.size*g.size {
		return true, true, errors.New("ходов больше нет")
	}

	return false, true, nil
}
func (g game) endGame(x int, y int, team int) bool {

	minX := x - winSize + 1
	maxX := x + winSize
	minY := y - winSize + 1
	maxY := y + winSize

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
		if qLength1 == winSize || qLength2 == winSize {
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
		if qLength1 == winSize || qLength2 == winSize {
			return true
		}
		j--
	}

	return false
}

func (g *game) PrintGame() {
	for i := range g.playingField {
		fmt.Println(g.playingField[i])
	}
}
