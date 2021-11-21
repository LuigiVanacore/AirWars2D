package AirWars2D

import (
	"AirWars2D/Math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Node interface {
 AttachChild(child Node)
 DetachChild(node Node) bool
GetTransform() ebiten.GeoM
SetTransform(transform ebiten.GeoM)
 Update()
Render(target *ebiten.Image, op *ebiten.DrawImageOptions)
GetWorldPosition() Math.Vector2D
GetWorldTransform() ebiten.GeoM
	AttachParent(node Node)
}

type SceneNode struct {
	children []Node
	parent Node
	transform ebiten.GeoM
}

func NewSceneNode() *SceneNode {
	return &SceneNode{}
}

func  (s* SceneNode) AttachChild(child Node) {
	child.AttachParent(s)

	s.children = append(s.children,child)
}

func (s *SceneNode) DetachChild(node Node) bool {
	for i, child := range s.children {
		if child == node {
			s.children[i] = s.children[len(s.children)-1]
			s.children  = s.children[:len(s.children)-1]
			return true
		}
	}
	return false
}

func (s* SceneNode) AttachParent(node Node) {
	s.parent = node
}

func (s *SceneNode) GetTransform() ebiten.GeoM {
	return s.transform
}

func (s* SceneNode) SetTransform(transform ebiten.GeoM) {
	s.transform = transform
}

func  (s *SceneNode) Update() {
	s.updateCurrent()
	s.updateChildren()
}

func (s* SceneNode) updateCurrent() {

}

func (s* SceneNode) updateChildren() {
	for _, child := range s.children {
		child.Update()
	}
}

func (s* SceneNode) Render(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM = s.GetTransform()

	s.renderCurrent(target,op)
	s.renderChildren(target, op)
}

func (s* SceneNode) renderCurrent( *ebiten.Image,  *ebiten.DrawImageOptions) {

}

func (s* SceneNode) renderChildren(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	for _, child :=  range s.children {
		child.Render(target, op)
	}
}

func (s* SceneNode) GetWorldPosition() Math.Vector2D {
	transform := s.GetWorldTransform()
	x,y := transform.Apply(0, 0)
	return Math.Vector2D{X: x, Y: y }
}


func (s* SceneNode) GetWorldTransform() ebiten.GeoM {
	transform := ebiten.GeoM{}

	for node := Node(s); node != nil; node = s.parent {
		getTransform := node.GetTransform()
		transform.Concat(getTransform)
	}

	return transform
}