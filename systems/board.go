package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type BoardSystem struct {
	*BoardEntity
}
type BoardEntity struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

func (bs *BoardSystem) New(w *ecs.World) {

	texture, err := common.LoadedSprite("textures/board.png")
	if err != nil {
		panic("Unable to load texture" + err.Error())
	}

	r := common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}
	b := BoardEntity{ecs.NewBasic(), common.SpaceComponent{}, r}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&b.BasicEntity, &b.RenderComponent, &b.SpaceComponent)
		}
	}
	bs.BoardEntity = &b
}

func (bs *BoardSystem) Update(dt float32) {
	b := bs.BoardEntity
	x := engo.CanvasWidth()
	y := engo.CanvasHeight()
	b.SpaceComponent.Position = engo.Point{X: -x/2 + 200, Y: -y/2 + 200}
	b.RenderComponent.Scale = engo.Point{X: x / 42, Y: y / 42}
}
func (bs *BoardSystem) Remove(basic ecs.BasicEntity) {
	if bs.BoardEntity == nil {
		return
	}
	if bs.BoardEntity.ID() == basic.ID() {
		bs.BoardEntity = nil
	}
}
