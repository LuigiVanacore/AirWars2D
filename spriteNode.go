package AirWars2D

import "github.com/hajimehoshi/ebiten/v2"

type SpriteNode struct {
	SceneNode
	sprite *ebiten.Image
}

func NewSpriteNode(sprite *ebiten.Image) *SpriteNode{
	return &SpriteNode{sprite: sprite}
}

func (s* SpriteNode) renderCurrent(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	target.DrawImage(s.sprite, op)
}