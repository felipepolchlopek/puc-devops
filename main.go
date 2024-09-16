package main

const (
	EMPTY    = ' '
	PLAYER_X = 'X'
	PLAYER_O = 'O'
)

var board [3][3]rune
var currentPlayer rune

// Inicializa o tabuleiro vazio
func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = EMPTY
		}
	}
	currentPlayer = PLAYER_X
}
