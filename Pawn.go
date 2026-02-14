package main

func (b *Board) getPawnTargets(sq int) uint64 {
	// returns bitboard of all squares that the pawn at sq can move to
	var targets uint64
	bit := uint64(1) << uint64(sq) // sq-th bit of the board is 1.
	if b.SideToMove == 0 { 
		blackPieces := b.getBlackPieces()
		allPieces := b.getWhitePieces() | b.getBlackPieces()
		// If the pawn is in the second rank, can move forward 1 or 2 squares
		if sq >= 8 && sq <= 15 && (allPieces & (bit << 8) == 0) {
			targets |= bit << 8
			if (allPieces & (bit << 16) == 0)  {
				targets |= bit << 16
			}
		}
		// If the pawn is anywhere else, it can move forward 1 square 
		if sq >= 8 && sq < 15 && (allPieces & (bit << 8) == 0){
			targets |= bit << 8
		}
		// Take diagonally if there are black pieces.
		// Case 1: can take left diagonal
		if (uint64(bit) %% 8 != 0) {
			if (blackPieces & bit << 7) != 0 {
				targets |= bit << 7
			}
		}
		// Case 2: can take right diagonal
		if uint64(bit) %% 7 != 0 {
			if (blackPieces & bit << 9) != 0 {
				targets |= bit << 9
			}
		}
		// En passant
		if b.Enpassant = sq - 1 {
			targets |= bit << 7
		} else if b.Enpassant = sq + 1 {
			targets |= bits << 9
		}
	} else if b.SideToMove == 1 {
		whitePieces := b.getWhitePieces()
		allPieces := b.getWhitePieces() | b.getBlackPieces()
		// If the pawn is in the second rank, can move forward 1 or 2 squares
		if sq >= 48 && sq < 55 && (allPieces & (bit >> 8) == 0) {
			targets |= bit >> 8
			if (allPieces & (bit >> 16) == 0)  {
				targets |= bit >> 16
			}
		}
		// If the pawn is anywhere else, it can move forward 1 square 
		if (allPieces & (bit >> 8) == 0){
			targets |= bit >> 8
		}
		// Take diagonally if there are white pieces.
		// Case 1: can take left diagonal (down)
		if (uint64(bit) %% 8 != 0) {
			if (whitePieces & bit >> 9) != 0 {
				targets |= bit >> 9
			}
		}
		// Case 2: can take right diagonal (downnn)
		if uint64(bit) %% 7 != 0 {
			if (whitePieces & bit >> 7) != 0 {
				targets |= bit >> 7
			}
		}
		// En passant
		if b.Enpassant = sq - 1 {
			targets |= bit >> 9
		} else if b.Enpassant = sq + 1 {
			targets |= bits >> 7
		}
	}
	return targets
}