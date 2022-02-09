package game

import "errors"

type game struct {
	size         int
	moves        int
	playingField [][]int
}

func NewGame(size int) game {
	newGame := game{size, 0, make([][]int, size)}
	for i := range newGame.playingField {
		newGame.playingField[i] = make([]int, size)
	}
	return newGame
}

func (g *game) NewMove(playerTeam int, x int, y int) (bool, error) {
	if x > g.size || x < 0 || y > g.size || y < 0 {
		return false, errors.New("ячейка находится за границами поля")
	}
	if g.playingField[x][y] == 0 {
		g.playingField[x][y] = playerTeam
	} else {
		return false, errors.New("ячейка уже занята")
	}
	g.moves++

	return (g.moves == g.size*g.size), nil
}

func (g *game) GenereteMove(playerTeam int) bool {
	for x := range g.playingField {
		for y := range g.playingField[x] {
			if g.playingField[x][y] == 0 {
				g.playingField[x][y] = playerTeam
				g.moves++
				return (g.moves == g.size*g.size)
			}
		}
	}
	return false
}

func (g *game) GetPlayingField() [][]int {
	return g.playingField
}
