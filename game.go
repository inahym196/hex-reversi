package hexReversi

import "fmt"

type Piece byte

const (
	PieceBlack Piece = iota
	PieceWhite
)

type Cell byte

const (
	CellEmpty Cell = iota
	CellBlack
	CellWhite
)

func cellFromPiece(p Piece) Cell {
	switch p {
	case PieceWhite:
		return CellWhite
	case PieceBlack:
		return CellBlack
	default:
		panic("invalid piece")
	}
}

func (c Cell) String() string {
	switch c {
	case CellEmpty:
		return "CellEmpty"
	case CellBlack:
		return "CellBlack"
	case CellWhite:
		return "CellWhite"
	default:
		panic("invalid cell")
	}
}

const (
	BoardRadius   = 4
	BoardDiameter = BoardRadius*2 + 1
)

type Board [][]Cell

func NewBoard() Board {
	board := make([][]Cell, BoardDiameter)
	for i := range BoardDiameter {
		board[i] = make([]Cell, BoardDiameter)
	}
	board[3][3] = CellWhite
	board[3][4] = CellBlack
	board[4][3] = CellBlack
	board[4][5] = CellWhite
	board[5][3] = CellWhite
	board[5][4] = CellBlack
	return board
}

func (b Board) isPlaced(row, column int) bool {
	return b[row][column] != CellEmpty
}

func (b Board) PutPiece(row, col int, piece Piece) error {
	if b.isPlaced(row, col) {
		return fmt.Errorf("board[%d][%d] is not empty, got %v", row, col, b[row][col])
	}
	b[row][col] = cellFromPiece(piece)
	return nil
}
