// Board and Moves
package (
	"math/bits"
)

const (
	FileA uint64 = 0x0101010101010101
	FileB uint64 = 0x0202020202020202
	FileG uint64 = 0x4040404040404040
	FileH uint64 = 0x8080808080808080
	Rank1 uint64 = 0x00000000000000FF
	Rank8 uint64 = 0xFF00000000000000
	
	// Combined masks for moves that jump two files
	FileAB uint64 = FileA | FileB
	FileGH uint64 = FileG | FileH
)

type Board struct {
	Pieces [12]uint64 // Represents 6 black and 6 white pieces
	// White, 1: pawn, 2: knight, 3: bishop, 4: rook, 5: queen, 6: king
	// Black, 7: pawn, 8: knight, 9: bishop, 10: rook, 11: queen, 12: king
	// Each entry of the array stores 64 bits so each bit represents
	// a square a piece is on. If a white pawn is on square A1, bit 1 is 1.
	SideToMove int // 0 white, 1 black
	CastlingRights uint8 // 1: White K, 2: White Q, 4: Black K, 8: Black Q
	Enpassant int // 0-63, -1 if no en passant. 
	HalfMove int // 50 move rule counter (pawn/piece capture)
	FullMove int // Number of turns 
}

func getIndex(rank, file int8) int8{
	return rank*8 + file
}

func getCoords(index int8) (rank, file int8){
	return index/8, index%8
}

type Move struct  {
	From, To, uint8
	Promotion int
	Piece int8 // Piece moving
	Captured int8 // Piece captured
	Flags int8 // 1: En passant, 2: Castling
}

func (b *Board) getLeastSigBit(bitboard uint64) int {
	// Finds the rightmost non zero entry of the bitboard 
	return bits.TrailingZeros64(bitboard)
}

func (b *Board) getPieceAt(sq int) int8 {
	// returns 1 if white pawn, etc as mentioned in Board 
	for i:=0; i < 12; i++ {
		if b.Pieces[i] & (1 << sq) != 0 {
			return int8(i+1)
		}
	}
	return 0
}

func (b *Board) getBlackPieces() uint64 {
	return b.Pieces[6] | b.Pieces[7] | b.Pieces[8] | b.Pieces[9] | b.Pieces[10] | b.Pieces[11]
}

func (b *Board) getWhitePieces() uint64 {
	return b.Pieces[0] | b.Pieces[1] | b.Pieces[2] | b.Pieces[3] | b.Pieces[4] | b.Pieces[5]
}

func (b *Board) isOccupied(sq, colour int) uint64 {
	// returns the white/black piece 1-12 if occupied. 
	// colour = 0 if white, 1 if black, 2 if doesn't matter (pawn move)
	if colour == 0 {
		for i:=0; i < 6; i++ {
			if b.Pieces[i] & (1 << sq) != 0 {
				return int8(i+1)
			}
		}
	} else if colour == 1 {
		for i:= 6; i < 12; i++ {
			if b.pieces[i] & (1 <, sq) != 0 {
				return int8(i+1)
			}
		}
	} else if colour == 2 {
		for i:=0; i < 12; i++ {
			if b.Piees[i] & (1 << sq) != 0 {
				return int8(i+1)
			}
		}
	}
	else {
		return int8(0)
	}
	
}

func (b *Board)generatePossibleMoves(sq int, board Board) []Move {
	moves = make([]Move, 0, 35)
	piece := getPieceAt(sq)
	if piece == 0 {
		return moves
	}

	switch piece {
		case 1: // white pawn. 
			moves = append(moves, getPawnTargets(sq))
		case 2: // white knight
			moves = append(moves, getKnightTargets(sq))
		case 3: // white bishop
			moves = append(moves, getBishopTargets(sq))
		case 4:// white rook
			moves = append(moves, getRookTargets(sq))
		case 5: // white queen
			moves = append(moves, getQueenTargets(sq))
		case 6: // white king
			moves = append(moves, getKingTargets(sq))
	}
	
	return moves
}



func (b *Board) generateLegalMoves() []Move {
	moves = make([]Move, 0, 35)
	// bitboard of all pieces of current side to play
	var currentSidePieces uint64
	if b.SideToMove == 0 {
		currentSidePieces = b.getWhitePieces()
	} else {
		currentSidePieces = b.getBlackPieces()
	}

	for currentSidePieces != 0 {
		sqToMove := b.getLeastSigBit(currentSidePieces)
		possibleMoves := b.generatePossibleMoves(sqToMove)
		

		currentSidePieces &= currentSidePieces - 1
	}
}
 
func isLegalMove(move Move, board Board) bool {
	if move.From == move.To {
		return false
	}
	if move.From < 0 || move.From > 63 || move.To < 0 || move.To > 63 {
		return false
	}

}


