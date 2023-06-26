package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wsandy1/2d-physics-engine/graphics"
	"github.com/wsandy1/2d-physics-engine/vector2"
)

type Game struct{}

var (
	inprog_vertices []vector2.Vec2
	first_saved     vector2.Vec2
	engine          PhysicsEngine
	message         string
)

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		inprog_vertices = append(inprog_vertices, vector2.Vec2{X: float32(x), Y: float32(y)})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		for i := 1; i < len(inprog_vertices); i++ {
			inprog_vertices[i] = vector2.Sub(inprog_vertices[i], inprog_vertices[0])
		}
		first_saved = inprog_vertices[0]

		inprog_vertices[0] = vector2.Vec2{X: 0, Y: 0}

		new_object := RigidBody{
			mass:             100,
			current_position: first_saved,
			last_position:    first_saved,
			acceleration:     vector2.Vec2{X: 0, Y: 0},
			poly: graphics.Polygon{
				Color:    color.RGBA{255, 0, 0, 255},
				Vertices: inprog_vertices,
			},
			point_forces: []PointForce{},
		}

		engine.RigidBodies = append(engine.RigidBodies, new_object)
		inprog_vertices = []vector2.Vec2{}
	}

	engine.Update(float32(ebiten.ActualTPS()))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range engine.RigidBodies {
		engine.RigidBodies[i].poly.Draw(screen, engine.RigidBodies[i].current_position, true)
	}

	ebitenutil.DebugPrint(screen, message)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	CalculateCentreOfMass([]vector2.Vec2{{X: 0, Y: 0}, {X: 10, Y: 0}, {X: 10, Y: 10}, {X: 10, Y: 0}})
	engine = PhysicsEngine{
		Gravity:     vector2.Vec2{X: 0, Y: 9.81},
		RigidBodies: []RigidBody{},
		substeps:    16,
	}
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("2D Physics Engine")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
