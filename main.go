package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
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
	w                  = screenWidth
	h                  = screenHeight
)

var (
	runnerImage *ebiten.Image
	err         error
	background  *ebiten.Image
	spaceShip   *ebiten.Image
	playerOne   player
	op          *ebiten.DrawImageOptions
)

type Position struct {
	X int
	Y int
}

// Create the player class
type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
	count      int
}

type Game struct {
	// count          int
	personPosition []Position
}

func (g *Game) Update() error {
	// s := runnerImage.Bounds().Size()
	// op := &ebiten.DrawImageOptions{}

	// if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {

	//
	// 	op.GeoM.Translate(-float64(s.X)+100, -float64(s.Y)+10)
	// 	fmt.Println("left")

	// } else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && inpututil.KeyPressDuration(ebiten.KeyArrowRight) != 0 {
	// 	fmt.Println()
	// 	g.count += 7
	// 	fmt.Println("right")

	// } else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
	// 	fmt.Println("down")
	// } else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
	// 	fmt.Println("up")
	// }
	movePlayer()
	return nil
}

func init() {
	// background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	// spaceShip, _, err = ebitenutil.NewImageFromFile("assets/1.png", ebiten.FilterDefault)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	playerOne = player{runnerImage, screenWidth / 6.0, screenHeight / 6.0, 4, 0}
}

func movePlayer() {

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerOne.yPos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerOne.yPos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// playerOne.xPos -= playerOne.speed
		fmt.Println("left")
		playerOne.count += 3
		// playerOne.yPos += 100
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerOne.xPos += 1.0
		fmt.Println("right")
		fmt.Println(playerOne.xPos)
		playerOne.count += 3
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/3, screenHeight/3)
	i := (playerOne.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(playerOne.image.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	vector.DrawFilledRect(screen, float32(10), float32(15), gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff}, false)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"log"

// 	"github.com/hajimehoshi/ebiten"
// 	"github.com/hajimehoshi/ebiten/ebitenutil"
// )

// // Our game constants
// const (
// 	screenWidth, screenHeight = 640, 480
// )

// // Create our empty vars
// var (
// 	err        error
// 	background *ebiten.Image
// 	spaceShip  *ebiten.Image
// 	playerOne  player
// )

// // Create the player class
// type player struct {
// 	image      *ebiten.Image
// 	xPos, yPos float64
// 	speed      float64
// }

// // Run this code once at startup
// func init() {
// 	// background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	spaceShip, _, err = ebitenutil.NewImageFromFile("1.png", ebiten.FilterDefault)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	playerOne = player{spaceShip, screenWidth / 5.0, screenHeight / 5.0, 4}
// }

// func movePlayer() {
// 	if ebiten.IsKeyPressed(ebiten.KeyUp) {
// 		playerOne.yPos -= playerOne.speed
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyDown) {
// 		playerOne.yPos += playerOne.speed
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
// 		playerOne.xPos -= playerOne.speed
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyRight) {
// 		playerOne.xPos += playerOne.speed
// 	}
// }

// func update(screen *ebiten.Image) error {
// 	movePlayer()
// 	if ebiten.IsDrawingSkipped() {
// 		return nil
// 	}
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(0, 0)
// 	// screen.DrawImage(background, op)

// 	playerOp := &ebiten.DrawImageOptions{}
// 	playerOp.GeoM.Translate(playerOne.xPos, playerOne.yPos)
// 	screen.DrawImage(playerOne.image, playerOp)

// 	return nil
// }

// func main() {
// 	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Hello, World!"); err != nil {
// 		log.Fatal(err)
// 	}
// }
