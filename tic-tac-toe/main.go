package main

import "fmt"


func main(){
	fmt.Println("tic tac toe!")
	fmt.Println("Two player mode: ")
	board := [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	printBoard(board)

	playerPiece := "X"
	moves := 0

	for moves < 9 {
		var playerInputRow int
		var playerInputCol int
		fmt.Printf("Player %s, enter coordinates for your move!\n", playerPiece)
		fmt.Println("Row: ")
		fmt.Scanf("%d", &playerInputRow)
		fmt.Println("Column: ")
		fmt.Scanf("%d", &playerInputCol)
		if !checkInput(playerInputRow, playerInputCol, board) {
			continue
		}

		board = makeMove(playerInputRow, playerInputCol, playerPiece, board)
		printBoard(board)

		if checkBoardForWin(playerPiece, board) && moves < 9 {
			fmt.Printf("Player %s wins!", playerPiece)
			break
		} else if checkBoardForWin(playerPiece, board) && moves == 9{
			fmt.Println("It's a draw!")
		}

		if playerPiece == "X" {
			playerPiece = "O"
		} else {
			playerPiece = "X"
		}

		moves += 1
	}


}

func printBoard(board [3][3]string) {
	for line := range board {
		fmt.Println(board[line])
	}
}

func checkInput(row int, column int, currentBoard [3][3]string) bool{
	if row > 2 || row < 0 || column > 2 || column < 0 {
		fmt.Println("Coordinates off the board! Try a number between 0 and 2: ")
		return false
	}
	if currentBoard[row][column] != "_" {
		fmt.Println("Space taken, choose another space.")
		return false
	}
	return true
}

func makeMove(row int, column int, piece string, currentBoard [3][3]string) [3][3]string {
	currentBoard[row][column] = piece
	return currentBoard
}

func checkBoardForWin(playerPiece string, currentBoard [3][3]string) (bool) {
	//row wins
	if currentBoard[0][0] == currentBoard[0][1] && currentBoard[0][1] == currentBoard[0][2] && currentBoard[0][0] != "_" {
		return true
	}

	//diagonal left wins
	if currentBoard[0][0] == currentBoard[1][1] && currentBoard[1][1] == currentBoard[2][2] && currentBoard[0][0] != "_"  {
		return true
	}

	//diagonal right wins
	if currentBoard[0][2] == currentBoard[1][1] && currentBoard[1][1] == currentBoard[2][0] && currentBoard[0][2] != "_"  {
		return true
	}
	
	//vertical wins, first 
	if currentBoard[0][0] == currentBoard[1][0] && currentBoard[1][0] == currentBoard[2][0] && currentBoard[0][0] != "_"  {
		return true
	}

	//vertical wins, second
	if currentBoard[0][1] == currentBoard[1][1] && currentBoard[1][1] == currentBoard[2][1] && currentBoard[0][1] != "_"  {
		return true
	}

	//vertical wins, third
	if currentBoard[0][2] == currentBoard[1][2] && currentBoard[1][2] == currentBoard[2][2] && currentBoard[0][2] != "_"  {
		return true
	}

	return false
}