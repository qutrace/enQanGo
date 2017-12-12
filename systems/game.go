package systems

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"github.com/qutrace/qanGo/ai"
	"github.com/qutrace/qanGo/game"
)

type GameSystem struct {
	world *ecs.World

	scan    bool
	running bool
	cur     game.Board

	moves  chan game.Move
	board  chan game.Board
	abort  chan bool
	report chan error
	done   chan game.Board
}

func (gs *GameSystem) New(world *ecs.World) {
	gs.world = world
	gs.board = make(chan game.Board)
	gs.moves = make(chan game.Move)
	gs.abort = make(chan bool)
	gs.report = make(chan error)
	gs.done = make(chan game.Board)
	gs.running = true

	movefunc := func(m engo.Message) {
		move, ok := m.(MoveMessage)
		if !ok { return }
		if gs.scan {
			gs.moves <- game.Move(move)
			gs.scan = false
		}
	}
	engo.Mailbox.Listen("MoveMessage", movefunc)

	go game.PlayChan(&game.Board{}, gs.board, gs.moves, gs.abort, gs.report, gs.done)
}

func print(g game.Board) {
	engo.Mailbox.Dispatch(PrintBoardMessage(g))
}
func (gs *GameSystem) Update(dt float32) {
	if engo.Input.Button("Restart").JustPressed() {
		gs.restart()
	}

	if !gs.running {
		return
	}

	select {
		case err := <-gs.report:
			fmt.Println(err)
		case gs.cur = <-gs.board:
			print(gs.cur)
			if !gs.cur.GetPlayer() {
				gs.moves <- ai.GetMove(gs.cur)
			} else {
				gs.scan = true
			}
		case gs.cur = <-gs.done:
			print(gs.cur)
			fmt.Println(gs.cur.GetState())
			gs.running = false
		default:
	}

}

func (gs *GameSystem) restart() {
	if gs.running {
			gs.abort <- true
			<-gs.done
		}
		gs.scan = false

		go game.PlayChan(&game.Board{}, gs.board, gs.moves, gs.abort, gs.report, gs.done)
		gs.running = true

}

func (gs *GameSystem) Remove(basic ecs.BasicEntity) {}
