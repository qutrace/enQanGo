package systems

import (
	"github.com/qutrace/qanGo/game"
)

type PrintBoardMessage game.Board

func (p PrintBoardMessage) Type() string {
	return "PrintBoardMessage"
}

type MoveMessage game.Move

func (m MoveMessage) Type() string {
	return "MoveMessage"
}
