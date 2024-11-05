package main

func isValidSudoku(board [][]byte) bool {
	rows := [9]map[byte]struct{}{}
	cols := [9]map[byte]struct{}{}
	boxes := [9]map[byte]struct{}{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cellVal := board[i][j]
			if cellVal == '.' {
				continue
			}
			if rows[i] == nil {
				rows[i] = map[byte]struct{}{}
			}
			if _, ok := rows[i][cellVal]; ok {
				return false
			}
			if cols[j] == nil {
				cols[j] = map[byte]struct{}{}
			}
			if _, ok := cols[j][cellVal]; ok {
				return false
			}
			if boxes[i/3*3+j/3] == nil {
				boxes[i/3*3+j/3] = map[byte]struct{}{}
			}
			if _, ok := boxes[i/3*3+j/3][cellVal]; ok {
				return false
			}
			rows[i][cellVal] = struct{}{}
			cols[j][cellVal] = struct{}{}
			boxes[i/3*3+j/3][cellVal] = struct{}{}
		}
	}
	return true
}
