package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	EMPTY    = ' '
	PLAYER_X = 'K'
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
func displayBoard() string {
	var sb strings.Builder
	sb.WriteString("  0   1   2\n")
	for i := 0; i < 3; i++ {
		sb.WriteString(fmt.Sprintf("%d", i))
		for j := 0; j < 3; j++ {
			sb.WriteString(fmt.Sprintf(" %c ", board[i][j]))
			if j < 2 {
				sb.WriteString("|")
			}
		}
		sb.WriteString("\n")
		if i < 2 {
			sb.WriteString(" ---|---|---\n")
		}
	}
	return sb.String()
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

// Verifica se há um vencedor
func checkWinner() rune {
	// Verifica linhas, colunas e diagonais
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != EMPTY {
			return board[i][0]
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != EMPTY {
			return board[0][i]
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != EMPTY {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != EMPTY {
		return board[0][2]
	}
	return EMPTY
}

// Verifica se há empate
func isDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == EMPTY {
				return false
			}
		}
	}
	return true
}

// Handler para exibir o tabuleiro
func boardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, displayBoard())
}

// Handler para realizar jogadas
func moveHandler(w http.ResponseWriter, r *http.Request) {
	rowStr := r.URL.Query().Get("row")
	colStr := r.URL.Query().Get("col")

	row, err1 := strconv.Atoi(rowStr)
	col, err2 := strconv.Atoi(colStr)

	if err1 != nil || err2 != nil || !isValidMove(row, col) {
		http.Error(w, "Movimento inválido", http.StatusBadRequest)
		return
	}

	makeMove(row, col)

	// Verificar se há vencedor ou empate
	if winner := checkWinner(); winner != EMPTY {
		fmt.Fprintf(w, "Jogador %c venceu!\n", winner)
	} else if isDraw() {
		fmt.Fprint(w, "Empate!\n")
	} else {
		switchPlayer()
		fmt.Fprintf(w, "Jogador %c fez uma jogada.\n", currentPlayer)
	}
}

func main() {
	initializeBoard()

	// Configurar handlers
	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/move", moveHandler)

	// Escutar na porta 8081
	fmt.Println("Servidor rodando na porta 8081...")
	http.ListenAndServe(":8081", nil)
}
