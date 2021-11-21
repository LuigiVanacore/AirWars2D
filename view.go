package AirWars2D

import (
	"AirWars2D/Math"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type View struct {
	center Math.Vector2D
	size Math.Vector2D
	rotation float64
	viewport Math.Rect
	transform ebiten.GeoM
	inverseTransform ebiten.GeoM
	transformUpdated bool
	invTransformUpdated bool
}

func NewDefaultView() *View{
	v := View{
		viewport: Math.Rect{
			Position: Math.Vector2D{0,0},
			Width:    1,
			Height:   1,
		},
	}
	v.Reset(0,0, 1000,  1000)
	return &v
}

func NewView(x,y,width,height float64) *View {
	v := NewDefaultView()
	v.Reset(x, y, width, height)
	return v
}

func (v* View) SetCenter(x, y float64) {
	v.center.X = x
	v.center.Y = y
}

func (v* View) GetCenter() Math.Vector2D {
	return v.center
}

func (v* View) SetSize(x, y float64) {
	v.size.X = x
	v.size.Y = y

	v.transformUpdated = false
	v.invTransformUpdated = false
}

func (v* View) GetSize() Math.Vector2D {
	return v.size
}

func (v* View) SetRotation(angle float64) {
	v.rotation = angle

	v.transformUpdated = false
	v.invTransformUpdated = false
}

func (v* View) GetRotation() float64 {
	return v.rotation
}

func (v* View) SetViewport(viewport Math.Rect)  {
	v.viewport = viewport
}

func (v* View) GetViewport() Math.Rect {
	return  v.viewport
}

func (v* View) Reset(x, y, width, height float64)  {
	v.center.X = x + width / 2
	v.center.Y = y + height / 2
	v.size.X = width
	v.size.Y = height
	v.rotation = 0
	v.transformUpdated = false
	v.invTransformUpdated = false
}

func (v* View) Move(x, y float64) {
	v.SetCenter(v.center.X + x, v.center.Y + y)
}

func (v* View) Rotation(angle float64) {
	v.SetRotation(v.rotation + angle)
}

func (v* View) Zoom(factor float64) {
	v.SetSize(v.size.X * factor, v.size.Y * factor)
}

func (v* View) GetTransform() ebiten.GeoM {
	if !v.transformUpdated {

		angle := v.rotation *  Math.FLOAT_PI / 180
		cosine := math.Cos(angle)
		sine := math.Sin(angle)
		tx := -v.center.X * cosine - v.center.Y * sine + v.center.X
		ty := v.center.X * sine - v.center.Y * cosine + v.center.Y

		a := 2 / v.size.X
		b := 2  / v.size.Y
		c := -a * v.center.X
		d := -b * v.center.Y

		v.transform = ebiten.GeoM{}
		v.transform.SetElement(0,0, a * cosine)
		v.transform.SetElement(0,1, a * sine)
		v.transform.SetElement(0,2, a * tx + c)
		v.transform.SetElement(1,0, -b * sine)
		v.transform.SetElement(1,1, b * cosine)
		v.transform.SetElement(1,2, b * ty + d)
		v.transform.SetElement(2,0, 0)
		v.transform.SetElement(2,1, 0)
		v.transform.SetElement(2,2, 1)
		v.transformUpdated = true
	}
	return v.transform
}

func (v* View) GetInverseTransform() ebiten.GeoM {
	if !v.invTransformUpdated {
		transform := v.GetTransform()
		transform.Invert()
		v.inverseTransform = transform
		v.invTransformUpdated = true
	}
	return v.inverseTransform

}