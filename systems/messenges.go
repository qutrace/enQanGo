package systems

import (
	"github.com/qutrace/qanGo/game"
)

type PrintBoardMessage game.Board

func (p PrintBoardMessage) Type() string {
	return "PrintBoardMessage"
}

type SetStateMessage struct {
	game.Board
}
func (s SetStateMessage) Type() string {
	return "SetStateMessage"
}
func (s SetStateMessage) GetBoard() game.Board {
	return s.Board
}


type MoveMessage game.Move

func (m MoveMessage) Type() string {
	return "MoveMessage"
}
