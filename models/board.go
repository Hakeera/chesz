package models

import "fmt"

// Tipo para representar o tabuleiro de xadrez
type Board [8][8]*Piece

// Função para inicializar o tabuleiro com as peças na posição inicial
func NewBoard() Board {
    var board Board

    // Definição das peças iniciais
    setup := []string{"R", "N", "B", "Q", "K", "B", "N", "R"}
    colors := []string{"White", "Black"}

    // Configurar torres, cavalos, bispos, rainha e rei
    for i, piece := range setup {
        board[0][i] = &Piece{Type: piece, Color: colors[0]}
        board[7][i] = &Piece{Type: piece, Color: colors[1]}
    }

    // Configurar peões
    for i := range 8 { 
        board[1][i] = &Piece{Type: "P", Color: colors[0]}
        board[6][i] = &Piece{Type: "P", Color: colors[1]}
    }

    return board
}

func (b *Board) MovePiece(fromRow, fromCol, toRow, toCol int) {
    piece := b[fromRow][fromCol] // Obter a peça da posição de origem
    if piece == nil {
        fmt.Println("Não há peça na posição de origem.")
        return
    }

    // Validação de Movimento 
    if !b.IsValidMove(piece, fromRow, fromCol, toRow, toCol) {
        fmt.Println("Movimento inválido!")
        return
    }

    // Realizar a movimentação
    b[toRow][toCol] = piece
    b[fromRow][fromCol] = nil
    fmt.Printf("Peça %s movida de (%d, %d) para (%d, %d)\n", piece.Type, fromRow, fromCol, toRow, toCol)
}
