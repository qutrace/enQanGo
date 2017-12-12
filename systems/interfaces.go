package systems

import (
	"engo.io/ecs"
)

type System interface {
	Update(dt float32)
	Remove(ecs.BasicEntity)
}
