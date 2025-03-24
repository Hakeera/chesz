package pieces

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Pieces armazenará todas as imagens das peças
type Pieces struct {
	WhiteQueen *ebiten.Image
	BlackQueen *ebiten.Image
}

// LoadPieces carrega as imagens das peças
func LoadPieces() *Pieces {
	whiteQueen, _, err := ebitenutil.NewImageFromFile("media/white-queen.png")
	if err != nil {
		log.Fatal(err)
	}
	blackQueen, _, err := ebitenutil.NewImageFromFile("media/black-queen.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Pieces{
		WhiteQueen: whiteQueen,
		BlackQueen: blackQueen,
	}
}

