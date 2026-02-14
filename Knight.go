package main

func (b *Board) getKnightTargets(sq int) uint64 {
	// returns bitboard of all squares that the knight at sq can move to
	var targets uint64
	bit := uint64(1) << uint64(sq) // sq-th bit of the board is 1.
	if b.SideToMove == 0 {
		// We don't need to check for going over or under the bitboard representation (>63 or <0)
		// The CPU will drop that bit into void
		whitePieces := b.getWhitePieces()
		// up 2 and right 1
		targets |= (bit << 17) & ^(FileA) & ^whitePieces
		// up 2 and left 1
		targets |= (bit << 15) & ^(FileH) & ^whitePieces
		// up 1 and right 2
		targets |= (bit << 10) & ^(FileAB) & ^whitePieces
		// up 1 and left 2
		targets |= (bit << 6) & ^(FileGH) & ^whitePieces
		// down 2 and right 1
		targets |= (bit >> 15) & ^(FileA) & ^whitePieces
		// down 2 and left 1
		targets |= (bit >> 17) & ^(FileH) & ^whitePieces
		// down 1 and right 2
		targets |= (bit >> 6) & ^(FileAB) & ^whitePieces
		// down 1 and left 2
		targets |= (bit >> 10) & ^(FileGH) & ^whitePieces
	} else if b.SideToMove == 1 {
		blackPieces := b.getBlackPieces()
		// up 2 and right 1
		targets |= (bit << 17) & ^(FileA) & ^blackPieces
		// up 2 and left 1
		targets |= (bit << 15) & ^(FileH) & ^blackPieces
		// up 1 and right 2
		targets |= (bit << 10) & ^(FileAB) & ^blackPieces
		// up 1 and left 2
		targets |= (bit << 6) & ^(FileGH) & ^blackPieces
		// down 2 and right 1
		targets |= (bit >> 15) & ^(FileA) & ^blackPieces
		// down 2 and left 1
		targets |= (bit >> 17) & ^(FileH) & ^blackPieces
		// down 1 and right 2
		targets |= (bit >> 6) & ^(FileAB) & ^blackPieces
		// down 1 and left 2
		targets |= (bit >> 10) & ^(FileGH) & ^blackPieces
	}
	return targets
}
