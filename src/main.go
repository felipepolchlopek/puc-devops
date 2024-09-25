package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	EMPTY    = ' '
	PLAYER_X = 'X'
	PLAYER_O = 'O'
)

var Board [3][3]rune
var CurrentPlayer rune

// Inicializa o tabuleiro vazio
func InitializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			Board[i][j] = EMPTY
		}
	}
	CurrentPlayer = PLAYER_X
}

// Exibe o tabuleiro
func DisplayBoard() string {
	var sb strings.Builder
	sb.WriteString("  0   1   2\n")
	for i := 0; i < 3; i++ {
		sb.WriteString(fmt.Sprintf("%d", i))
		for j := 0; j < 3; j++ {
			sb.WriteString(fmt.Sprintf(" %c ", Board[i][j]))
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
func IsValidMove(row, col int) bool {
	return row >= 0 && row < 3 && col >= 0 && col < 3 && Board[row][col] == EMPTY
}

// Alterna entre os jogadores
func SwitchPlayer() {
	if CurrentPlayer == PLAYER_X {
		CurrentPlayer = PLAYER_O
	} else {
		CurrentPlayer = PLAYER_X
	}
}

// Faz a jogada no tabuleiro
func MakeMove(row, col int) {
	Board[row][col] = CurrentPlayer
}

// Verifica se há um vencedor
func CheckWinner() rune {
	for i := 0; i < 3; i++ {
		if Board[i][0] == Board[i][1] && Board[i][1] == Board[i][2] && Board[i][0] != EMPTY {
			return Board[i][0]
		}
		if Board[0][i] == Board[1][i] && Board[1][i] == Board[2][i] && Board[0][i] != EMPTY {
			return Board[0][i]
		}
	}
	if Board[0][0] == Board[1][1] && Board[1][1] == Board[2][2] && Board[0][0] != EMPTY {
		return Board[0][0]
	}
	if Board[0][2] == Board[1][1] && Board[1][1] == Board[2][0] && Board[0][2] != EMPTY {
		return Board[0][2]
	}
	return EMPTY
}

// Verifica se há empate
func IsDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Board[i][j] == EMPTY {
				return false
			}
		}
	}
	return true
}

// Handler para exibir o tabuleiro
func BoardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, DisplayBoard())
}

// Handler para realizar jogadas
func MoveHandler(w http.ResponseWriter, r *http.Request) {
	rowStr := r.URL.Query().Get("row")
	colStr := r.URL.Query().Get("col")

	row, err1 := strconv.Atoi(rowStr)
	col, err2 := strconv.Atoi(colStr)

	if err1 != nil || err2 != nil || !IsValidMove(row, col) {
		http.Error(w, "Movimento inválido", http.StatusBadRequest)
		return
	}

	MakeMove(row, col)

	if winner := CheckWinner(); winner != EMPTY {
		fmt.Fprintf(w, "Jogador %c venceu!\n", winner)
	} else if IsDraw() {
		fmt.Fprint(w, "Empate!\n")
	} else {
		SwitchPlayer()
		fmt.Fprintf(w, "Jogador %c fez uma jogada.\n", CurrentPlayer)
	}
}

func main() {
	InitializeBoard()

	http.HandleFunc("/board", BoardHandler)
	http.HandleFunc("/move", MoveHandler)

	fmt.Println("Servidor rodando na porta 8081...")
	http.ListenAndServe(":8081", nil)
}
