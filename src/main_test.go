package main

import (
	"testing"
)

func TestInitializeBoard(t *testing.T) {
	InitializeBoard()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Board[i][j] != EMPTY {
				t.Errorf("Expected Board[%d][%d] to be EMPTY, got %c", i, j, Board[i][j])
			}
		}
	}
}

// Testa a alternância de jogadores
func TestSwitchPlayer(t *testing.T) {
	CurrentPlayer = PLAYER_X
	SwitchPlayer()
	if CurrentPlayer != PLAYER_O {
		t.Errorf("Expected CurrentPlayer to be %c, got %c", PLAYER_O, CurrentPlayer)
	}

	SwitchPlayer()
	if CurrentPlayer != PLAYER_X {
		t.Errorf("Expected CurrentPlayer to be %c, got %c", PLAYER_X, CurrentPlayer)
	}
}

// Testa se um movimento é válido
func TestIsValidMove(t *testing.T) {
	InitializeBoard()

	// Testa um movimento válido
	if !IsValidMove(1, 1) {
		t.Errorf("Expected move to be valid")
	}

	// Faz o movimento
	MakeMove(1, 1)

	// Testa um movimento inválido na mesma posição
	if IsValidMove(1, 1) {
		t.Errorf("Expected move to be invalid after the position is taken")
	}
}

// Testa a função de verificar o vencedor
func TestCheckWinner(t *testing.T) {
	InitializeBoard()

	// Simula uma vitória do PLAYER_X
	Board[0][0] = PLAYER_X
	Board[0][1] = PLAYER_X
	Board[0][2] = PLAYER_X

	winner := CheckWinner()
	if winner != PLAYER_X {
		t.Errorf("Expected winner to be %c, got %c", PLAYER_X, winner)
	}

	// Testa quando não há vencedor
	InitializeBoard()
	winner = CheckWinner()
	if winner != EMPTY {
		t.Errorf("Expected no winner, got %c", winner)
	}
}

// Testa se há empate
func TestIsDraw(t *testing.T) {
	InitializeBoard()

	// Preenche o tabuleiro sem vencedores
	Board = [3][3]rune{
		{PLAYER_X, PLAYER_O, PLAYER_X},
		{PLAYER_X, PLAYER_X, PLAYER_O},
		{PLAYER_O, PLAYER_X, PLAYER_O},
	}

	if !IsDraw() {
		t.Errorf("Expected the game to be a draw")
	}

	// Testa um jogo que não é empate
	InitializeBoard()
	if IsDraw() {
		t.Errorf("Expected the game to not be a draw")
	}
}
