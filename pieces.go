//go:generate fyne bundle -o bundled.go pieces
package main

import (
	"fyne.io/fyne"
	"github.com/notnil/chess"
)

func resourceForPiece(piece chess.Piece) fyne.Resource {
	switch piece.Color() {
	case chess.Black:
		{
			switch piece.Type() {
			case chess.Pawn:
				return resourceBlackPawnSvg
			case chess.King:
				return resourceBlackKingSvg
			case chess.Queen:
				return resourceBlackQueenSvg
			case chess.Knight:
				return resourceBlackKnightSvg
			case chess.Bishop:
				return resourceBlackBishopSvg
			case chess.Rook:
				return resourceBlackRookSvg
			}
		}
	case chess.White:
		{
			switch piece.Type() {
			case chess.Pawn:
				return resourceWhitePawnSvg
			case chess.King:
				return resourceWhiteKingSvg
			case chess.Queen:
				return resourceWhiteQueenSvg
			case chess.Knight:
				return resourceWhiteKnightSvg
			case chess.Bishop:
				return resourceWhiteBishopSvg
			case chess.Rook:
				return resourceWhiteRookSvg
			}
		}
	}
	return nil
}
