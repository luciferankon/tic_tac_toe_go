package player

type Player struct {
	Name   string
	Symbol string
	Moves  []int
}

func (player Player) GetName() string {
	return player.Name
}

func (player *Player) MakeMove(move int){
	player.Moves = append(player.Moves, move)
}

func (player Player) HasMoves(moves []int) bool {
	flag := true
	for i := 0; i < len(moves); i++ {
		for j := 0; j < len(player.Moves); j++ {
			flag = player.Moves[j] == moves[i]
			if flag {
				break
			}
		}
		if flag == false {
			break
		}
	}
	return flag
}
