package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"

)

type CursorSystem struct {
	mouseTracker *MouseTracker
}

func (c *CursorSystem) Update(dt float32) {
	if c.mouseTracker == nil {
		return
	}
	switch engo.Input.Mouse.Action {
	case engo.Press:
		switch engo.Input.Mouse.Button {
		case engo.MouseButtonLeft:
			c.DispatchMessage()
		}

	}
}

func (c *CursorSystem) DispatchMessage() {
	m := CursorMessage{
		X: c.mouseTracker.MouseX,
		Y: c.mouseTracker.MouseY,
	}
	engo.Mailbox.Dispatch(m)
	//fmt.Println(m)
	x:= engo.CanvasWidth()
	y:= engo.CanvasHeight()
	mx := c.mouseTracker.MouseX
	my := c.mouseTracker.MouseY
	move := MoveMessage{
		X: int(int(mx + x/2 - 200) / int(x/6)),
		Y: int(int(my + y/2 - 200) / int(y/6)),
	}
	//fmt.Println(move)
	engo.Mailbox.Dispatch(move)
}

func (c *CursorSystem) New(w *ecs.World) {
	m := MouseTracker{ecs.NewBasic(), common.MouseComponent{Track: true}}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&m.BasicEntity, &m.MouseComponent, nil, nil)
		}
	}
	c.mouseTracker = &m
}

func (c *CursorSystem) Remove(basic ecs.BasicEntity) {
	if c.mouseTracker == nil {
		return
	}
	if basic.ID() == c.mouseTracker.BasicEntity.ID() {
		c.mouseTracker = nil
	}
}

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type CursorMessage engo.Point

func (c CursorMessage) Type() string {
	return "CursorMessage"
}
