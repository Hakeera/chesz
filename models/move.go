package models

import "fmt"

// Alteração na função IsValidMove para aceitar a peça
func (b *Board) IsValidMove(piece *Piece, fromRow, fromCol, toRow, toCol int) bool {
    if piece == nil {
        fmt.Println("Não há peça na posição de origem.")
        return false
    }

    switch piece.Type {
    case "P":
        return isValidPawnMove(b, piece, fromRow, fromCol, toRow, toCol)
    case "B":
        return isValidBishopMove(b, piece, fromRow, fromCol, toRow, toCol)
    case "R":
        return isValidRookMove(b, piece, fromRow, fromCol, toRow, toCol)
    case "N":
        return isValidKnightMove(b, piece, fromRow, fromCol, toRow, toCol) // Movimentação do cavalo
    default:
        fmt.Println("Regra de movimentação ainda não implementada para esta peça.")
        return false
    }
}

// Lógica de movimentação do peão
func isValidPawnMove(b *Board, piece *Piece, fromRow, fromCol, toRow, toCol int) bool {
    direction := 1 // Direção para peças brancas (movem "para baixo" no índice da matriz)

    if piece.Color == "Black" {
        direction = -1 // Peças pretas movem "para cima"
    }

    // Movimento de uma casa para frente
    if toRow == fromRow+direction && fromCol == toCol && b[toRow][toCol] == nil {
        return true
    }

    // Primeiro movimento: duas casas para frente
    if (fromRow == 1 && piece.Color == "White" || fromRow == 6 && piece.Color == "Black") &&
        toRow == fromRow+(2*direction) && fromCol == toCol &&
        b[toRow][toCol] == nil && b[fromRow+direction][toCol] == nil {
        return true
    }

    // Captura na diagonal
    if toRow == fromRow+direction && (toCol == fromCol+1 || toCol == fromCol-1) &&
        b[toRow][toCol] != nil && b[toRow][toCol].Color != piece.Color {
        return true
    }

    return false
}

// Movimento do Bispo
func isValidBishopMove(b *Board, piece *Piece, fromRow, fromCol, toRow, toCol int) bool {
    if abs(toRow-fromRow) != abs(toCol-fromCol) {
        return false // O bispo só se move na diagonal
    }

    return b.isPathClear(piece, fromRow, fromCol, toRow, toCol) 
}

// Movimento do Cavalo
func isValidKnightMove(b *Board, piece *Piece,fromRow, fromCol, toRow, toCol int) bool {

    // Verificar se o movimento segue a regra do cavalo: "L"
    rowDiff := abs(toRow - fromRow)
    colDiff := abs(toCol - fromCol)

    if (rowDiff == 2 && colDiff == 1) || (rowDiff == 1 && colDiff == 2) {
        // Movimento válido para o cavalo
	        // Verificar captura: se a casa de destino tem uma peça da mesma cor, é inválido
        if b[toRow][toCol] != nil {
            if b[toRow][toCol].Color == piece.Color {
		println("PEÇA DE MESMA COR")
                return false // Não pode capturar peça da mesma cor
            }
        }
        return true
    }

    return false // Movimento inválido para o cavalo
}

// Movimento da Torre
func isValidRookMove(b *Board, piece *Piece, fromRow, fromCol, toRow, toCol int) bool {
    if fromRow != toRow && fromCol != toCol {
        return false // A torre só se move em linha reta
    }

    return b.isPathClear(piece, fromRow, fromCol, toRow, toCol)
}

// Função auxiliar para obter o valor absoluto
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func (b *Board) isPathClear(piece *Piece, fromRow, fromCol, toRow, toCol int) bool {
    rowDiff := toRow - fromRow
    colDiff := toCol - fromCol

    rowStep, colStep := 0, 0
    if rowDiff != 0 {
        rowStep = rowDiff / abs(rowDiff) // Define a direção da linha (-1, 0, 1)
    }
    if colDiff != 0 {
        colStep = colDiff / abs(colDiff) // Define a direção da coluna (-1, 0, 1)
    }

    r, c := fromRow+rowStep, fromCol+colStep
    for r != toRow || c != toCol {
        if b[r][c] != nil {
            return false // Tem peça bloqueando
        }
        r += rowStep
        c += colStep
    }

    // Verificar se a casa de destino está ocupada por uma peça da mesma cor
    if b[toRow][toCol] != nil && b[toRow][toCol].Color == piece.Color {
        return false // Não pode mover para uma casa com peça da mesma cor
    }

    return true
}
