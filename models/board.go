package models

import (
	"chesz/views"
)

// Tipo para representar o tabuleiro de xadrez
type Board [8][8]*Piece

// Função para inicializar o tabuleiro com as peças na posição inicial
func NewBoard() Board {
    var board Board

    // Definição das peças iniciais
    setup := []string{"R", "N", "B", "Q", "K", "B", "N", "R"}
    colors := []string{"Black", "White"}

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

func (b *Board) MovePiece(fromRow, fromCol, toRow, toCol int, turn string) bool {
    piece := b[fromRow][fromCol]

    if piece == nil {
        views.PrintMessage("Não há peça na posição de origem!")
        return false // Não há peça na posição de origem
    }

    // Verifica se a peça pertence ao jogador do turno atual
    if (turn == "White" && piece.Color != "White") || (turn == "Black" && piece.Color != "Black") {
        views.PrintMessage("Você só pode mover suas próprias peças!")
        return false
    }

    if !b.IsValidMove(piece, fromRow, fromCol, toRow, toCol) {
        return false // Movimento inválido
    }

    // Simular o movimento
    temp := b[toRow][toCol] // Salvar peça de destino, caso exista
    b[toRow][toCol] = piece
    b[fromRow][fromCol] = nil

    // Verificar se o próprio Rei ficou em xeque
    if b.isKingInCheck(piece.Color) {
        // Reverter o movimento, pois é ilegal
        b[fromRow][fromCol] = piece
        b[toRow][toCol] = temp
        views.PrintMessage("Movimento inválido: Rei em Xeque!")
        return false
    }

    return true // Movimento válido
}

func (b *Board) isSquareAttacked(row, col int, attackerColor string) bool {
    for r := range 8 { 
        for c := range 8 { 
            piece := b[r][c]
            if piece != nil && piece.Color == attackerColor {
                if b.IsValidMove(piece, r, c, row, col) {
                    return true // Se alguma peça adversária puder se mover para essa casa, ela está atacada
                }
            }
        }
    }
    return false
}

func (b *Board) isKingInCheck(kingColor string) bool {
    // Encontrar a posição do Rei
    var kingRow, kingCol int
    for r := range 8 { 
        for c := range 8 { 
            piece := b[r][c]
            if piece != nil && piece.Type == "K" && piece.Color == kingColor {
                kingRow, kingCol = r, c
                break
            }
        }
    }

    // Verificar se o Rei está sendo atacado
    opponentColor := "White"
    if kingColor == "White" {
        opponentColor = "Black"
    }

    return b.isSquareAttacked(kingRow, kingCol, opponentColor)
}

func (b *Board) IsCheckmate(color string) bool {
    if !b.isKingInCheck(color) {
        return false // Se o rei não está em xeque, não é xeque-mate
    }

    for fromRow, row := range b {
        for fromCol, piece := range row {
            if piece != nil && piece.Color == color {
                for toRow := 0; toRow < 8; toRow++ {
                    for toCol := 0; toCol < 8; toCol++ {
                        if b.IsValidMove(piece, fromRow, fromCol, toRow, toCol) {
                            // Simula o movimento
                            temp := b[toRow][toCol]
                            b[toRow][toCol] = piece
                            b[fromRow][fromCol] = nil

                            // Verifica se ainda está em xeque
                            stillInCheck := b.isKingInCheck(color)

                            // Reverte o movimento
                            b[fromRow][fromCol] = piece
                            b[toRow][toCol] = temp

                            if !stillInCheck {
                                return false // Ainda há movimentos legais
                            }
                        }
                    }
                }
            }
        }
    }

    return true // Não há movimentos válidos, então é xeque-mate
}

