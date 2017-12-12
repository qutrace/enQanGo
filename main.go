package main

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/qutrace/enqanGo/systems"
)

type myScene struct{}

func (*myScene) Type() string {
	return "myGame"
}

func (*myScene) Preload() {
	engo.Files.Load("textures/true.png")
	engo.Files.Load("textures/false.png")
	engo.Files.Load("textures/board.png")
}

func (*myScene) Setup(world *ecs.World) {
	engo.Input.RegisterButton("Restart", engo.Space)
	common.SetBackground(color.White)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	world.AddSystem(&systems.DrawingSystem{})
	world.AddSystem(&systems.BoardSystem{})
	world.AddSystem(&systems.CursorSystem{})
	world.AddSystem(&systems.GameSystem{})

}

func main() {
	opts := engo.RunOptions{
		Title:  "HelloWorld",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &myScene{})
	fmt.Println("test")
}
