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

type Game struct {
	PieceImages map[string]*ebiten.Image       // Armazena as imagens das peças
	Pieces      map[[2]int]string              // Mapeia posições para peças
	Selected    *[2]int                        // Armazena a posição da peça selecionada
}

// Atualiza o jogo, capturando cliques do jogador
func (g *Game) Update() error {
	// Capturar clique do mouse
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		col := x / tileSize
		row := y / tileSize

		if g.Selected == nil {
			// Seleciona a peça se houver uma na posição clicada
			if _, exists := g.Pieces[[2]int{row, col}]; exists {
				g.Selected = &[2]int{row, col}
			}
		} else {
			// Se o jogador clicar na mesma peça, cancelar a seleção
			if g.Selected[0] == row && g.Selected[1] == col {
				g.Selected = nil
				return nil
			}

			// Verifica se é um movimento válido (por ora, qualquer posição vazia)
			if g.IsValidMove(g.Selected[0], g.Selected[1], row, col) {
				// Atualiza a posição da peça
				g.Pieces[[2]int{row, col}] = g.Pieces[*g.Selected]
				delete(g.Pieces, *g.Selected)
			}

			// Reseta a seleção
			g.Selected = nil
		}
	}
	return nil
}

// Verifica se o movimento da peça é válido
func (g *Game) IsValidMove(fromRow, fromCol, toRow, toCol int) bool {
	// Apenas para teste: permitir qualquer movimento dentro do tabuleiro
	return (toRow >= 0 && toRow < 8 && toCol >= 0 && toCol < 8)
}

// Desenha o tabuleiro e as peças
func (g *Game) Draw(screen *ebiten.Image) {
	// Desenhar tabuleiro
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			x := col * tileSize
			y := row * tileSize
			rect := ebiten.NewImage(tileSize, tileSize)

			// Destacar casa selecionada
			if g.Selected != nil && g.Selected[0] == row && g.Selected[1] == col {
				rect.Fill(color.RGBA{200, 200, 50, 255}) // Amarelo para seleção
			} else if (row+col)%2 == 0 {
				rect.Fill(color.RGBA{240, 217, 181, 255}) // Cor clara
			} else {
				rect.Fill(color.RGBA{181, 136, 99, 255}) // Cor escura
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(rect, op)
		}
	}

	// Desenhar peças no tabuleiro
	for pos, piece := range g.Pieces {
		img, exists := g.PieceImages[piece]
		if !exists {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileSize)/float64(img.Bounds().Dx()), float64(tileSize)/float64(img.Bounds().Dy())) // Ajustar tamanho
		op.GeoM.Translate(float64(pos[1]*tileSize), float64(pos[0]*tileSize))                                    // Ajustar posição
		screen.DrawImage(img, op)
	}
}

// Define o layout da janela
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Função principal
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("CHEISZ")

	// Carregar imagens das peças
	queenImg, _, err := ebitenutil.NewImageFromFile("media/black-amazon.png")
	if err != nil {
		log.Fatal(err)
	}

	// Criar instância do jogo com a dama preta no tabuleiro
	game := &Game{
		PieceImages: map[string]*ebiten.Image{
			"black-queen": queenImg,
		},
		Pieces: map[[2]int]string{
			{0, 3}: "black-queen", // Coloca a dama preta na posição inicial (d8)
		},
	}

	// Iniciar o jogo
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

