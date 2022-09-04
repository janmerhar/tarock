package gui

import "github.com/hajimehoshi/ebiten/v2"

func DrawEntireDeckTop(screen *ebiten.Image, card_back *ebiten.Image) {
	card_back_op := &ebiten.DrawImageOptions{}
	card_back_op.GeoM.Translate(float64(1280-15*75), float64(0))
	card_back_op.GeoM.Scale(.25, .25)

	for i := 0; i < 16; i++ {
		screen.DrawImage(card_back, card_back_op)
		card_back_op.GeoM.Translate(float64(75), float64(0))
	}
}

func DrawEntireDeckBottom(screen *ebiten.Image, card_back *ebiten.Image) {
	card_back_op := &ebiten.DrawImageOptions{}
	card_back_op.GeoM.Translate(float64(1280-15*75), float64(2330))
	card_back_op.GeoM.Scale(.25, .25)

	for i := 0; i < 16; i++ {
		screen.DrawImage(card_back, card_back_op)
		card_back_op.GeoM.Translate(float64(75), float64(0))
	}
}
