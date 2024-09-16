package main

import "fmt"

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

// Exibe o tabuleiro
func displayBoard() {
	fmt.Println("  0   1   2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d", i)
		for j := 0; j < 3; j++ {
			fmt.Printf(" %c ", board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println(" ---|---|---")
		}
	}
}

// Verifica se a posição é válida e está disponível
func isValidMove(row, col int) bool {
	return row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == EMPTY
}

// Alterna entre os jogadores
func switchPlayer() {
	if currentPlayer == PLAYER_X {
		currentPlayer = PLAYER_O
	} else {
		currentPlayer = PLAYER_X
	}
}

// Faz a jogada no tabuleiro
func makeMove(row, col int) {
	board[row][col] = currentPlayer
}
