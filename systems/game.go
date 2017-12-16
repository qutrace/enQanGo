package systems

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/qutrace/qanGo/game"
)

type GameSystem struct {
	world *ecs.World

	board *game.Board
	move *game.Move
	running bool
}

func printBoard(g game.Board) {
	engo.Mailbox.Dispatch(PrintBoardMessage(g))
}

func (gs *GameSystem) New(world *ecs.World) {
	gs.world = world
	setstate := func(m engo.Message) {
		state, ok := m.(SetStateMessage)
		if !ok {return}
		board := state.GetBoard()
		gs.board = &board
		gs.running = true
		printBoard(board)
	}
	getmove := func(m engo.Message) {
		move, ok := m.(MoveMessage)
		if !ok {return}
		mo := game.Move(move)
		if gs.running {
			gs.move = &mo
		}
	}
	engo.Mailbox.Listen("MoveMessage", getmove)
	engo.Mailbox.Listen("SetStateMessage", setstate)

	gs.board = &game.Board{}
	gs.running = true
}

func (gs *GameSystem) Update(dt float32) {
	if engo.Input.Button("Restart").JustPressed() {
		engo.Mailbox.Dispatch(SetStateMessage{game.Board{}})
	}
	if gs.board == nil {return}
	if gs.move == nil {return}
	if !gs.running {return}

	err := gs.board.Apply(*gs.move)
	if err != nil {
		fmt.Println(err)
	} else {
		printBoard(*gs.board)
	}

	gs.move = nil

	state := gs.board.GetState()
	if state != game.Unknown {
		fmt.Println(state)
		gs.running = false
	}
}

func (gs *GameSystem) Remove(basic ecs.BasicEntity) {
}


