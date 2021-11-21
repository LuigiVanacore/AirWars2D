package AirWars2D

import (
	"AirWars2D/Assets"
	"AirWars2D/Math"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Background = iota
	Air
	LayerCount
)


type World struct {
	window *ebiten.Image
	worldView View
	textures []*ebiten.Image

	sceneGraph SceneNode
	sceneLayers []*SceneNode

	worldBounds Math.Rect
	spawnPosition Math.Vector2D
	scrollSpeed float64
	player *Aircraft
}


func NewWorld(screenwidth, screenheight float64) *World {
	screenBound := screen.Bounds()
	worldView := NewView(float64(screenBound.Min.X), float64(screenBound.Min.Y), float64(screenBound.Max.X), float64(screenBound.Max.Y))
	worldBounds := Math.NewRect(0,0, float64(screenBound.Max.X), 2000)
	spawnPosition := Math.Vector2D{}
	spawnPosition.X = float64(screenBound.Max.X) / 2
	spawnPosition.Y = worldBounds.Height - float64(screenBound.Max.Y) / 2
	worldView.SetCenter(spawnPosition.X, spawnPosition.Y)
	world := World{window: screen,worldView: *worldView, spawnPosition: spawnPosition,scrollSpeed: -50}

	world.LoadTextures()
	world.BuildScene()

	return &world
}

func (w* World) LoadTextures() {
	w.textures = append(w.textures, Assets.ResourceManager().GetTextures()...)
}

func (w* World) BuildScene() {
	for i :=0; i < LayerCount; i++ {
		layer := NewSceneNode()
		w.sceneLayers[i] = layer
		w.sceneGraph.AttachChild(layer)
	}

	texture := Assets.ResourceManager().GetTexture(Assets.Desert)

	backgroundNode := NewSpriteNode(texture)

	w.sceneLayers[Background].AttachChild(backgroundNode)

}