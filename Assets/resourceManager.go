package Assets

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"log"
)

const (
	Desert = iota
	Eagle
	Raptor
)

const (
	fontSize = 24
)

var instance *resourceManager

func ResourceManager() *resourceManager {
	if instance == nil {
		instance = newResourceManager()
	}
	
	return instance
}

type resourceManager struct {

	images []*ebiten.Image
	font font.Face
}

func newResourceManager() *resourceManager {
	r := &resourceManager{}
	r.LoadAssets()
	return r
}

func (r*resourceManager) LoadAssets() {
	var images []*ebiten.Image = make([]*ebiten.Image,Raptor+1)
	
	images[Desert] = ebiten.NewImageFromImage(loadImage(DesertTexture))
	images[Eagle] = ebiten.NewImageFromImage(loadImage(EagleTexture))
	images[Raptor] = ebiten.NewImageFromImage(loadImage(RaptorTexture))

	r.images = append(r.images, images...)
	
	
	r.font = loadFont(SansationFont)
}

func (r*resourceManager) GetTextures() []*ebiten.Image {
	return r.images
}

func (r*resourceManager) GetTexture(texture uint) *ebiten.Image {
	return r.images[texture]
}

func (r*resourceManager) GetFont() font.Face {
	return r.font
}

func loadImage(texture []byte) *ebiten.Image {
	image, _, err := image.Decode(bytes.NewReader(texture))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(image)
}

func loadFont(f []byte) font.Face {
	tt, err := opentype.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	gamefont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull ,
	})
	if err != nil {
		log.Fatal(err)
	}
	return gamefont
}