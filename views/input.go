// Package views
package views

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Mapeia colunas (letras) para índices
var colMap = map[rune]int{'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7}

// GetMove solicita a jogada do usuário e converte para índices do tabuleiro
func GetMove() (int, int, int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite seu movimento: ") 
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Removendo espaços extras e quebras de linha
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return 0, 0, 0, 0, fmt.Errorf("formato inválido. Use algo como 'e2 e4'")
	}

	fromRow, fromCol, err1 := parsePosition(parts[0])
	toRow, toCol, err2 := parsePosition(parts[1])
	if err1 != nil || err2 != nil {
		return 0, 0, 0, 0, fmt.Errorf("posição inválida")
	}

	return fromRow, fromCol, toRow, toCol, nil
}

// Converte "e2" para índices (6,4) do tabuleiro
func parsePosition(pos string) (int, int, error) {
	if len(pos) != 2 {
		return 0, 0, fmt.Errorf("posição inválida")
	}

	col, ok := colMap[rune(pos[0])]
	row, err := strconv.Atoi(string(pos[1]))

	if !ok || err != nil || row < 1 || row > 8 {
		return 0, 0, fmt.Errorf("posição inválida")
	}

	// Converter para índice da matriz (linha invertida porque matriz começa do topo)
	return 8 - row, col, nil
}

