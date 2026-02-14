package main

func (b *Board) getRookTargets(sq int) uint64 {
	var targets uint64
	bit := uint64(1) << uint64(sq)
	nStop := false
	sStop := false
	eStop := false
	wStop := false

	nPos := bit
	sPos := bit
	ePos := bit
	wPos := bit

	var friendlyPieces, enemyPieces uint64
	if b.SideToMove == 0 {
		friendlyPieces = b.getWhitePieces()
		enemyPieces = b.getBlackPieces()
	} else {
		friendlyPieces = b.getBlackPieces()
		enemyPieces = b.getWhitePieces()
	}

	for i := 1; i < 8; i++ {
		if nStop && sStop && eStop && wStop {
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
	}
	return targets
}
