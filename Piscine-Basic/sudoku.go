package main

func solveSudoku(board [][]int) bool {
	for i := 0; i < 9; i++ { //this code looks at every spot on the sudoku board by going row by row, column by column, and checking every number on the board
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 { // at any point on the board, if we find zero, it means this spot is empty and needs a number
				for k := 1; k <= 9; k++ { // this line of code puts numbers in empty spot begining with 1
					if isValid(board, i, j, k) { // isValid function checks if the number is valid for that empty space
						board[i][j] = k         // then places the number temporarily
						if solveSudoku(board) { // we call the solveSudoku function to repeatedly solve the rest of the puzzle
							return true // if it successfully solve the whole puzzle it return true
						}
						board[i][j] = 0 // this means no number fits this cell
					}
				}
				return false // if no number fits return false and backtracks to the former cell
			}
		}
	}
	return true //return true if the puzzle is solved successfully
}

func isValid(board [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num ||
			board[row][i] == num ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}
