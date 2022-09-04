package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/janmerhar/tarock/internal/gui"
)

var (
	background *ebiten.Image
	card_back  *ebiten.Image
	// horizontal_rectangle *image.RGBA
)

func init() {
	background, _, _ = ebitenutil.NewImageFromFile("./img/table-top.png")
	card_back, _, _ = ebitenutil.NewImageFromFile("./img/tarock22.png")
	// horizontal_rectangle = ebiten.NewImageFromImage(image.Rect(100, 100, 100, 100))
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(background, nil)

	// card_back_op := &ebiten.DrawImageOptions{}
	// card_back_op.GeoM.Translate(float64(0), float64(0))
	// card_back_op.GeoM.Scale(.25, .25)
	// screen.DrawImage(card_back, card_back_op)

	// card_back_op.GeoM.Translate(float64(200), float64(0))
	// screen.DrawImage(card_back, card_back_op)

	// gui.DrawEntireDeck(screen, card_back)
	gui.DrawEntireDeckTop(screen, card_back)
	gui.DrawEntireDeckBottom(screen, card_back)
	// card_back.Clear()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
