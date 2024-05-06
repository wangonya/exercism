package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupied := 0
	for _, x := range cb[file] {
		if x {
			occupied++
		}
	}
	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupied := 0

	if rank > 0 && rank <= 8 {
		for _, x := range cb {
			if x[rank-1] {
				occupied++
			}
		}
	}

	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	lcb := len(cb)
	lx := 1
	for _, x := range cb {
		lx = len(x)
		break
	}
	return lcb * lx
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupied := 0
	for _, x := range cb {
		for _, y := range x {
			if y {
				occupied++
			}
		}
	}
	return occupied
}
