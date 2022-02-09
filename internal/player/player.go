package player

type player struct {
	team int
}

func NewPlayer(team int) player {
	return player{team}
}

func (p *player) GetPayerTeam() int {
	return p.team
}
