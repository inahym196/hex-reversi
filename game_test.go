package hexReversi_test

import (
	"testing"

	hexReversi "github.com/inahym196/hex-reversi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		err := board.PutPiece(4, 4, hexReversi.PieceBlack)
		require.NoError(t, err)

		assertBoardState(t, board, expected)
	})

	t.Run("ボード外判定", func(t *testing.T) {
		tests := []struct {
			row, col int
			wantErr  bool
		}{
			{0, 0, false}, {0, 4, false}, {4, 8, false},
			{0, 5, true}, {1, 6, true}, {2, 7, true},
			{3, 8, true}, {4, 9, true}, {10, 10, true},
		}
		for _, tt := range tests {
			board := hexReversi.NewBoard()

			if err := board.PutPiece(tt.row, tt.col, hexReversi.PieceBlack); (err != nil) != tt.wantErr {
				t.Errorf("expected wantErr %v, got %v", tt.wantErr, err)
			}
		}
	})

	t.Run("配置済み判定", func(t *testing.T) {
		board := hexReversi.NewBoard()
		err := board.PutPiece(3, 3, hexReversi.PieceBlack)
		assert.Error(t, err)
	})
}
