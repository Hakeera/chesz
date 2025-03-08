package views

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetMove solicita ao jogador um movimento e converte a entrada para coordenadas
func GetMove() (int, int, int, int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Digite seu movimento (ex: e2 e4): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			PrintMessage("Entrada inválida. Formato correto: e2 e4")
			continue
		}

		fromRow, fromCol, valid1 := parsePosition(parts[0])
		toRow, toCol, valid2 := parsePosition(parts[1])
		if valid1 && valid2 {
			return fromRow, fromCol, toRow, toCol
		}

		PrintMessage("Coordenadas inválidas! Tente novamente.")
	}
}

// parsePosition converte "e2" para coordenadas da matriz
func parsePosition(pos string) (int, int, bool) {
	if len(pos) != 2 {
		return 0, 0, false
	}
	col := int(pos[0] - 'a')
	row, err := strconv.Atoi(string(pos[1]))
	if err != nil || row < 1 || row > 8 || col < 0 || col > 7 {
		return 0, 0, false
	}
	return 8 - row, col, true // Convertendo para índice da matriz
}

