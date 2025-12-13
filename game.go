package hexReversi

import "fmt"

type Piece bool

const (
	PieceBlack Piece = true
	PieceWhite Piece = false
)

func (p Piece) Opponent() Piece { return !p }

func (p Piece) String() string {
	switch p {
	case PieceBlack:
		return "PieceBlack"
	case PieceWhite:
		return "PieceWhite"
	default:
		panic("invalid piece")
	}
}

type Cell uint8

const (
	CellEmpty Cell = iota
	CellBlack
	CellWhite
)

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

const (
	BoardRadius = 4
	BoardWidth  = BoardRadius*2 + 1
)

type Board [][]Cell

func NewBoard() Board {
	board := make([][]Cell, BoardWidth)
	for i := range BoardWidth {
		board[i] = make([]Cell, BoardWidth)
	}
	board[3][3] = CellWhite
	board[3][4] = CellBlack
	board[4][3] = CellBlack
	board[4][5] = CellWhite
	board[5][3] = CellWhite
	board[5][4] = CellBlack
	return board
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (b Board) isInBoard(row, col int) bool {
	maxColumn := BoardRadius*2 - abs(BoardRadius-row)
	return (0 <= col && col <= maxColumn) && (0 <= row && row <= BoardRadius*2)
}

type Position struct {
	Row, Column int
}

func (b Board) collectFlippableInDirection(row, col, dy, dx int, piece Piece) (flips []Position) {
	row += dy
	col += dx
	for b.isInBoard(row, col) {
		switch b[row][col] {
		case cellFromPiece(piece.Opponent()):
			flips = append(flips, Position{row, col})
		case cellFromPiece(piece):
			return flips
		case CellEmpty:
			return []Position{}
		}
		row += dy
		col += dx
	}
	return []Position{}
}

func (b Board) collectFlippable(row, col int, piece Piece) (flips []Position) {
	dirs := []struct{ x, y int }{
		{1, 0}, {1, -1}, {0, -1},
		{-1, 0}, {-1, 1}, {0, 1},
	}
	for _, dir := range dirs {
		dx, dy := dir.x, dir.y
		flipsInDir := b.collectFlippableInDirection(row, col, dy, dx, piece)
		flips = append(flips, flipsInDir...)
	}
	return flips
}

func (b Board) isPlaced(row, column int) bool {
	return b[row][column] != CellEmpty
}

func (b Board) PutPiece(row, col int, piece Piece) error {
	if !b.isInBoard(row, col) {
		return fmt.Errorf("position(%v,%v) is out of board", row, col)
	}
	if b.isPlaced(row, col) {
		return fmt.Errorf("board[%d][%d] is not empty, got %v", row, col, b[row][col])
	}
	flips := b.collectFlippable(row, col, piece)
	if len(flips) == 0 {
		return fmt.Errorf("flippable piece not exists")
	}
	for _, flip := range flips {
		b[flip.Row][flip.Column] = cellFromPiece(piece)
	}
	b[row][col] = cellFromPiece(piece)
	return nil
}
