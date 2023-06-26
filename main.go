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

type PhysicsEngine struct {
	objects  []PhysicsObject
	gravity  vector2.Vec2
	substeps uint16
}

func (e *PhysicsEngine) Update(tps float32) {
	var dt float32 = 1 / tps
	// sub_dt := dt / float32(e.substeps)
	// for i := 0; i < int(e.substeps); i++ {
	e.ApplyGravity()
	e.N2Solve()
	e.VerletSolve(dt)
	// }

}

func (e *PhysicsEngine) ApplyGravity() {
	for i := range e.objects {
		e.objects[i].Accelerate(e.gravity)
	}
}

func (e *PhysicsEngine) VerletSolve(dt float32) {
	for i := range e.objects {
		vel := vector2.Sub(e.objects[i].current_position, e.objects[i].last_position)
		e.objects[i].last_position = e.objects[i].current_position
		e.objects[i].current_position = vector2.Add(vector2.Add(e.objects[i].current_position, vel), vector2.ConstMul(vector2.ConstMul(e.objects[i].acceleration, dt), dt))
		e.objects[i].acceleration = vector2.Vec2{X: 0, Y: 0}
	}
}

func (e *PhysicsEngine) N2Solve() {
	for i := range e.objects {
		var resultant vector2.Vec2
		for j := range e.objects[i].pointforces {
			resultant = vector2.Add(resultant, e.objects[i].pointforces[j].force)
		}
		e.objects[i].Accelerate(vector2.ConstDiv(resultant, float32(e.objects[i].mass)))
	}
}

type PhysicsObject struct {
	mass             uint16
	current_position vector2.Vec2
	last_position    vector2.Vec2
	acceleration     vector2.Vec2
	poly             graphics.Polygon
	pointforces      []PointForce
}

func (o *PhysicsObject) Accelerate(v vector2.Vec2) {
	o.acceleration = vector2.Add(o.acceleration, v)
}

type PointForce struct {
	origin vector2.Vec2
	force  vector2.Vec2
}

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

		new_object := PhysicsObject{
			mass:             100,
			current_position: first_saved,
			last_position:    first_saved,
			acceleration:     vector2.Vec2{X: 0, Y: 0},
			poly: graphics.Polygon{
				Color:    color.RGBA{255, 0, 0, 255},
				Vertices: inprog_vertices,
			},
			pointforces: []PointForce{{
				origin: vector2.Vec2{X: 0, Y: 0},
				force:  vector2.Vec2{X: 0, Y: 0},
			}},
		}

		engine.objects = append(engine.objects, new_object)
		inprog_vertices = []vector2.Vec2{}
	}

	engine.Update(float32(ebiten.ActualTPS()))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range engine.objects {
		engine.objects[i].poly.Draw(screen, engine.objects[i].current_position, true)
	}

	ebitenutil.DebugPrint(screen, message)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	engine = PhysicsEngine{
		gravity:  vector2.Vec2{X: 0, Y: 9.81},
		objects:  []PhysicsObject{},
		substeps: 16,
	}
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("2D Physics Engine")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
