package systems

import (
  "engo.io/ecs"
  "engo.io/engo"
  "engo.io/engo/common"
)

type Stone struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	BoardComponent
	PlayerComponent
}
type PlayerComponent bool

type BoardComponent struct {
	X int
	Y int
}

func newStone(x int, y int, player bool) Stone {
	s := Stone{BasicEntity: ecs.NewBasic()}
	s.PlayerComponent = PlayerComponent(player)
	s.BoardComponent = BoardComponent{X: x, Y: y}

	ts := "textures/"
	if player {
		ts += "true.png"
	} else {
		ts += "false.png"
	}
	texture, err := common.LoadedSprite(ts)
	if err != nil {
		panic("Unable to load texture" + err.Error())
	}
	s.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{X: 1, Y: 1},
	}
	s.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{float32(x * 50), float32(y * 50)},
		Width:    10,
		Height:   10,
	}
	return s
}

func (s *Stone) Update() {
  x := engo.CanvasWidth()
  y := engo.CanvasHeight()
  w := x/6.
  h := y/6.
  s.SpaceComponent.Position.X = float32(200 + w * float32(s.BoardComponent.X) - x/2)
  s.SpaceComponent.Position.Y = float32(200 + h * float32(s.BoardComponent.Y) - y/2)
  s.SpaceComponent.Width = w
  s.SpaceComponent.Height = h
  s.RenderComponent.Scale = engo.Point{X: w/7 ,Y: h/7}
}
