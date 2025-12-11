package hexReversi_test

import (
	"testing"

	hexReversi "github.com/inahym196/hex-reversi"
)

const (
	E = hexReversi.CellEmpty
	B = hexReversi.CellBlack
	W = hexReversi.CellWhite
)

var initBoard = [][]hexReversi.Cell{
	//0 1  2  3  4  5  6  7  8
	{E, E, E, E, E, E, E, E, E}, // 0
	{E, E, E, E, E, E, E, E, E}, // 1
	{E, E, E, E, E, E, E, E, E}, // 2
	{E, E, E, W, B, E, E, E, E}, // 3
	{E, E, E, B, E, W, E, E, E}, // 4
	{E, E, E, W, B, E, E, E, E}, // 5
	{E, E, E, E, E, E, E, E, E}, // 6
	{E, E, E, E, E, E, E, E, E}, // 7
	{E, E, E, E, E, E, E, E, E}, // 8
}

func assertBoardState(t *testing.T, board hexReversi.Board, expected [][]hexReversi.Cell) {
	t.Helper()

	for row := range expected {
		for col := range expected[row] {
			if board[row][col] != expected[row][col] {
				t.Errorf("board[%d][%d]: expected %v, got %v",
					row, col, expected[row][col], board[row][col])
			}
		}
	}
}

func TestNewBoard(t *testing.T) {
	expected := initBoard
	board := hexReversi.NewBoard()
	assertBoardState(t, board, expected)
}

func TestBoard_PutPiece(t *testing.T) {
	t.Run("初期ボードへ1つ置く", func(t *testing.T) {
		board := hexReversi.NewBoard()
		expected := initBoard
		expected[4][4] = hexReversi.CellBlack

		if err := board.PutPiece(4, 4, hexReversi.PieceBlack); err != nil {
			t.Fatalf("expected nil, got %v", err)
		}

		assertBoardState(t, board, expected)
	})
}
