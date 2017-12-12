package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/qutrace/qanGo/game"
)

type DrawingSystem struct {
	world    *ecs.World
	entities []*Stone
}

func (ds *DrawingSystem) New(w *ecs.World) {
	ds.world = w
	print := func(m engo.Message) {
		board, ok := m.(PrintBoardMessage)
		if !ok {
			return
		}
		ds.DisplayBoard(game.Board(board))
	}
	engo.Mailbox.Listen("PrintBoardMessage", print)
}
func (ds *DrawingSystem) Update(dt float32) {
	for _,s := range ds.entities {
		s.Update()
	}
}
func (ds *DrawingSystem) Remove(basic ecs.BasicEntity) {
	del := -1
	for index, e := range ds.entities {
		if e.BasicEntity.ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		ds.entities = append(ds.entities[:del], ds.entities[del+1:]...)
	}
}

func (ds *DrawingSystem) AddStone(s Stone) {
	ds.entities = append(ds.entities, &s)
	for _,system := range ds.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
				sys.Add(&s.BasicEntity, &s.RenderComponent, &s.SpaceComponent)
		}
	}
	s.Update()
}


func (ds *DrawingSystem) DisplayBoard(b game.Board) {
	es := ds.entities
	ds.entities = make([]*Stone, 0)
	for _, e := range es {
		ds.world.RemoveEntity(e.BasicEntity)
	}

	for y, row := range b {
		for x, cell := range row {
			if cell != nil { ds.AddStone(newStone(x,y, *cell)) }
		}
	}
}
