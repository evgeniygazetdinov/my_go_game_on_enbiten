package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth        = 320
	screenHeight       = 240
	gridSize           = 10
	xGridCountInScreen = screenWidth / gridSize
	yGridCountInScreen = screenHeight / gridSize
	frameOX            = 0
	frameOY            = 32
	frameWidth         = 32
	frameHeight        = 32
	frameCount         = 8
)

var (
	runnerImage *ebiten.Image
)

type Position struct {
	X int
	Y int
}

type Game struct {
	count          int
	personPosition []Position
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		fmt.Println("left")
		g.count += 7
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && inpututil.KeyPressDuration(ebiten.KeyArrowRight) != 0 {
		fmt.Println()
		g.count += 7
		fmt.Println("right")
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		fmt.Println("down")
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		fmt.Println("up")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	vector.DrawFilledRect(screen, float32(10), float32(15), gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
