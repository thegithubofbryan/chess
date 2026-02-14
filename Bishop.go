package main

func (b *Board) getBishopTargets(sq int) uint64 {
	// returns bitboard of all squares that the bishop at sq can move to
	var targets uint64
	bit := uint64(1) << uint64(sq) // sq-th bit of the board is 1.
	neStop := false
	nwStop := false
	seStop := false
	swStop := false

	nePos := bit
	nwPos := bit
	sePos := bit
	swPos := bit

	var friendlyPieces, enemyPieces uint64
	if b.SideToMove == 0 {
		friendlyPieces = b.getWhitePieces()
		enemyPieces = b.getBlackPieces()
	} else {
		friendlyPieces = b.getBlackPieces()
		enemyPieces = b.getWhitePieces()
	}

	for i := 1; i < 8; i++ {
		if neStop && nwStop && seStop && swStop {
			break
		}
		// North east movement
		if nePos&FileH != 0 || (nePos<<9)&friendlyPieces != 0 {
			neStop = true
		}
		if !neStop {
			nePos <<= 9
			targets |= nePos
			if nePos&enemyPieces != 0 {
				neStop = true
			}
		}
		// Northwest movement
		if nwPos&FileA != 0 || (nwPos<<7)&friendlyPieces != 0 {
			nwStop = true
		}
		if !nwStop {
			nwPos <<= 7
			targets |= nwPos
			if nwPos&enemyPieces != 0 {
				nwStop = true
			}
		}
		// Southeast movement
		if sePos&FileH != 0 || (sePos>>7)&friendlyPieces != 0 {
			seStop = true
		}
		if !seStop {
			sePos >>= 7
			targets |= sePos
			if sePos&enemyPieces != 0 {
				seStop = true
			}
		}
		// Southwest movement
		if swPos&FileA != 0 || (swPos>>9)&friendlyPieces != 0 {
			swStop = true
		}
		if !swStop {
			swPos >>= 9
			targets |= swPos
			if swPos&enemyPieces != 0 {
				swStop = true
			}
		}
	}
	}
	return targets
}
