package main

func (b *Board) getQueenTargets(sq int) uint64 {
	var targets uint64
	bit := uint64(1) << uint64(sq)
	nStop := false
	sStop := false
	eStop := false
	wStop := false
	neStop := false
	nwStop := false
	seStop := false
	swStop := false

	nPos := bit
	sPos := bit
	ePos := bit
	wPos := bit
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
		if nStop && sStop && eStop && wStop && neStop && nwStop && seStop && swStop {
			break
		}
		// North movement
		if nPos&Rank8 != 0 || (nPos<<8)&friendlyPieces != 0 {
			nStop = true
		}
		if !nStop {
			nPos <<= 8
			targets |= nPos
			if nPos&enemyPieces != 0 {
				nStop = true
			}
		}
		// South movement
		if sPos&Rank1 != 0 || (sPos>>8)&friendlyPieces != 0 {
			sStop = true
		}
		if !sStop {
			sPos >>= 8
			targets |= sPos
			if sPos&enemyPieces != 0 {
				sStop = true
			}
		}
		// East movement
		if ePos&FileH != 0 || (ePos<<1)&friendlyPieces != 0 {
			eStop = true
		}
		if !eStop {
			ePos <<= 1
			targets |= ePos
			if ePos&enemyPieces != 0 {
				eStop = true
			}
		}
		// West movement
		if wPos&FileA != 0 || (wPos>>1)&friendlyPieces != 0 {
			wStop = true
		}
		if !wStop {
			wPos >>= 1
			targets |= wPos
			if wPos&enemyPieces != 0 {
				wStop = true
			}
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
	return targets
}
