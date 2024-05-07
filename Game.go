package main

import (
    "fmt"
    "os"
    "os/exec"
)

var board [3][3]string
var currentPlayer string

func main() {
    initBoard()
    currentPlayer = "X"
    clearScreen()
    printBoard()

    for {
        fmt.Printf("Player %s's turn\n", currentPlayer)
        fmt.Print("Enter row and column (e.g., 1 2): ")
        var row, col int
        fmt.Scanln(&row, &col)

        if isValidMove(row-1, col-1) {
            board[row-1][col-1] = currentPlayer
            clearScreen()
            printBoard()

            if checkWin() {
                fmt.Printf("Player %s wins!\n", currentPlayer)
                break
            }

            if checkDraw() {
                fmt.Println("It's a draw!")
                break
            }

            switchPlayer()
        } else {
            fmt.Println("Invalid move. Please try again.")
        }
    }
}

func initBoard() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = " "
        }
    }
}

func printBoard() {
    for i := 0; i < 3; i++ {
        fmt.Printf(" %s | %s | %s\n", board[i][0], board[i][1], board[i][2])
        if i < 2 {
            fmt.Println("---+---+---")
        }
    }
}

func isValidMove(row, col int) bool {
    return row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == " "
}

func checkWin() bool {
    for i := 0; i < 3; i++ {
        if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
            return true
        }
        if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
            return true
        }
    }
    if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
        return true
    }
    if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
        return true
    }
    return false
}

func checkDraw() bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == " " {
                return false
            }
        }
    }
    return true
}

func switchPlayer() {
    if currentPlayer == "X" {
        currentPlayer = "O"
    } else {
        currentPlayer = "X"
    }
}

func clearScreen() {
    cmd := exec.Command("clear") // for Linux/OS X
    cmd.Stdout = os.Stdout
    cmd.Run()
}
