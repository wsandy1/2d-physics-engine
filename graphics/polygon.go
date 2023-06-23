package graphics

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/wsandy1/2d-physics-engine/vector2d"
)

var (
	// blank image used to colour triangle vertices
	whiteImage = ebiten.NewImage(3, 3)

	// whiteSubImage is an internal sub image of whiteImage.
	// Use whiteSubImage at DrawTriangles instead of whiteImage in order to avoid bleeding edges.
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
}

// struct representing a polygon
type Polygon struct {
	Color    color.RGBA
	Vertices []vector2d.Vec2
}

// draws a polygon to the supplied image (usually the screen)
func (p Polygon) Draw(screen *ebiten.Image, position vector2d.Vec2, aa bool) {
	var path vector.Path
	// start path at first vertex
	path.MoveTo(p.Vertices[0].X(), p.Vertices[0].Y())
	for i := 1; i < len(p.Vertices); i++ {
		// draw a line to each subsequent vertex
		path.LineTo(p.Vertices[i].X(), p.Vertices[i].Y())
	}
	path.Close()

	// populate list of vertices and indices
	var vs []ebiten.Vertex
	var is []uint16
	vs, is = path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		// apply provided position effectively shifting by X and Y
		r, g, b, a := p.Color.RGBA()
		vs[i].DstX = (vs[i].DstX + float32(position.X()))
		vs[i].DstY = (vs[i].DstY + float32(position.Y()))
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = float32(r) / float32(0xff)
		vs[i].ColorG = float32(g) / float32(0xff)
		vs[i].ColorB = float32(b) / float32(0xff)
		vs[i].ColorA = float32(a) / float32(0xff)
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.AntiAlias = aa
	op.FillRule = ebiten.EvenOdd
	screen.DrawTriangles(vs, is, whiteSubImage, op)
}

// rotate polygon around a given point
func (p *Polygon) Rotate(theta float64, centre vector2d.Vec2) {
	var shift []vector2d.Vec2
	for i := range p.Vertices {
		shift = append(shift, vector2d.Sub(p.Vertices[i], centre))
	}

	var rot []vector2d.Vec2
	for i := range shift {
		rot = append(rot, shift[i].Rotate(theta))
	}

	var back []vector2d.Vec2
	for i := range rot {
		back = append(back, vector2d.Add(rot[i], centre))
	}

	p.Vertices = back
}
