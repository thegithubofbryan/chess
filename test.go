package main

import (
	"fmt"
)

const (
	FileA uint64 = 0x0101010101010101
	FileB uint64 = 0x0202020202020202
	FileG uint64 = 0x4040404040404040
	FileH uint64 = 0x8080808080808080

	// Combined masks for moves that jump two files
	FileAB uint64 = FileA | FileB
	FileGH uint64 = FileG | FileH
)

func main() {
	// b := uint64(1) << uint64(10)
	fmt.Println(uint64(1) << FileA)
	fmt.Println(uint64(1) << ^FileA)
	// fmt.Println(FileG)
	// fmt.Println(FileH)
	// fmt.Println(FileAB)
	// fmt.Println(FileGH)
}
