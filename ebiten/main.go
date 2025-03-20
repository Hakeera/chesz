package main

import (
        "image/color"
        "log"

        "github.com/hajimehoshi/ebiten/v2"
        "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
        screenWidth  = 640
        screenHeight = 640
        tileSize     = screenWidth / 8
)

type Game struct{
	AmazImg *ebiten.Image
}

func (g *Game) Update() error {
        return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
        for row := 0; row < 8; row++ {
                for col := 0; col < 8; col++ {
                        x := col * tileSize
                        y := row * tileSize
                        rect := ebiten.NewImage(tileSize, tileSize)
                        if (row+col)%2 == 0 {
                                rect.Fill(color.RGBA{240, 217, 181, 255}) // Cor clara
                        } else {
                                rect.Fill(color.RGBA{181, 136, 99, 255}) // Cor escura
                        }
                        op := &ebiten.DrawImageOptions{}
                        op.GeoM.Translate(float64(x), float64(y))
                        screen.DrawImage(rect, op)
                }
        }
        ebitenutil.DebugPrint(screen, "cheizin")
	screen.DrawImage(
		g.AmazImg,
		&ebiten.DrawImageOptions{},
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
        return screenWidth, screenHeight
}

func main() {
        ebiten.SetWindowSize(screenWidth, screenHeight)
        ebiten.SetWindowTitle("CHEISZ")
		amazImg, _, err := ebitenutil.NewImageFromFile("media/black-amazon.png")
	if err != nil {
	// Handel error
	log.Fatal(err)
		}
		
		if err := ebiten.RunGame(&Game{AmazImg: amazImg}); err != nil {
                log.Fatal(err)
        }
}
