package main

import (
	"fmt"
)

func main() {
	board := [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}} // hamne ek board Name ka variable initialize kiya hai and vo ek array jisme hamne haar ek cell ko "-" se fill kiya hai
	currentPlayer := "X" // fir maine default value ya default player ko "X" maan liya hai 
	fmt.Println("Welcome to Tic Tac Toe!")

	for { // for ka loop tab end hga jab koi bhi player jeet jaye ya tuo draw ho jaye or ya tuo player invalid chezz inter kare
		fmt.Printf("Player %s's turn: ", currentPlayer) // ye line print karegi ki kiski turn hai "x"/"O"
		var row, col int // Maine yaha per 2 variable ko declare kiya hai 
		fmt.Scanln(&row, &col) //Scanln se ham user se input lenge row/col ka

		if row < 1 || row > 3 || col < 1 || col > 3 || board[row-1][col-1] != "_" { // Yaha per ham check kar rhe hai ki inputs valid hai ki nahi 
			fmt.Println("Invalid move. Please try again.")
			continue
		}

		board[row-1][col-1] = currentPlayer // yaha per ham tictactoe board per jate hai and ham Player ka symbol input karte hai "X/O"
		fmt.Println("  1 2 3") // Ise line ke through ham Har ek column ko mark kar rhe hai from 1,2,3
		for i := 0; i < 3; i++ { // ise loop ko use kar rhe hai because hame rows ko bhi mark karna hai
			fmt.Printf("%d ", i+1) // ye line print kar rhi / mark kar rhi hai har ek row ko
			for j := 0; j < 3; j++ { // we are iterating over each column because abhi hamne symbol "x"/"O" ko print karwana hai board per 
				fmt.Printf("%s ", board[i][j])//this is how we print the symbol
			}
			fmt.Println()
		}


		//3 conditions jab players jeet ta hai row mein same symbol align hue / column mein / diagnolly 
		if (board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer) ||
			(board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer) {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}//diagonally checking

		for i := 0; i < 3; i++ {
			if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
				fmt.Printf("Player %s wins!\n", currentPlayer)
				return
			}
			if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
				fmt.Printf("Player %s wins!\n", currentPlayer)
				return
			}
		}// rows and column check kar rhe hai 

		draw := true //initially draw ko true rakha hai because agar last mein koi space khali bachi tuo draw ni hga tan mujhe draw ki value ko false karna hoga haar ek cell check hora ki space " " tuo ni hai 
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == "_" {
					draw = false
					break
				}
			}
		}
		if draw { // Agar draw ki value check karne ke badh true hue tuo ye chalega varna nahi 
			fmt.Println("It's a draw!")
			break
		}

		if currentPlayer == "X" { // Yaha per Players ko switch kar rha hu agar "X" ki chance thi tuo ab "O" ki chance hogi
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
