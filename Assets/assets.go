package Assets


import (
	_ "embed"
	_ "image/png"
)


//go:embed Desert.png
var DesertTexture []byte

//go:embed Eagle.png
var EagleTexture []byte

//go:embed Raptor.png
var RaptorTexture []byte

//go:embed Sansation.ttf
var SansationFont []byte