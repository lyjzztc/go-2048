package main

import (
	"fmt"
	"math/rand"
	"time"
)

const gridSize = 4

type Game struct {
	board [gridSize][gridSize]int
	score int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	game.Run()
}

func NewGame() *Game {
	g := &Game{}
	g.AddTile()
	g.AddTile()
	return g
}

func (g *Game) Run() {
	for {
		g.PrintBoard()
		var move string
		fmt.Print("Enter move (W/A/S/D): ")
		fmt.Scan(&move)
		switch move {
		case "w", "W":
			g.MoveUp()
		case "s", "S":
			g.MoveDown()
		case "a", "A":
			g.MoveLeft()
		case "d", "D":
			g.MoveRight()
		default:
			fmt.Println("Invalid move! Please use W, A, S, or D.")
			continue
		}
		g.AddTile()
		if g.IsGameOver() {
			g.PrintBoard()
			fmt.Println("Game Over!")
			break
		}
	}
}

// MoveUp, MoveDown, MoveLeft, MoveRight handle the moves and merging.
func (g *Game) MoveUp()   { g.board = transpose(g.board); g.MoveLeft(); g.board = transpose(g.board) }
func (g *Game) MoveDown() { g.board = transpose(g.board); g.MoveRight(); g.board = transpose(g.board) }
func (g *Game) MoveLeft() {
	for i := 0; i < gridSize; i++ {
		g.board[i] = g.merge(g.board[i])
	}
}
func (g *Game) MoveRight() {
	for i := 0; i < gridSize; i++ {
		reverse(&g.board[i])
		g.board[i] = g.merge(g.board[i])
		reverse(&g.board[i])
	}
}

func (g *Game) AddTile() {
	empty := [][2]int{}
	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			if g.board[r][c] == 0 {
				empty = append(empty, [2]int{r, c})
			}
		}
	}
	if len(empty) > 0 {
		pos := empty[rand.Intn(len(empty))]
		g.board[pos[0]][pos[1]] = 2 * (rand.Intn(2) + 1)
	}
}

func (g *Game) IsGameOver() bool {
	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			if g.board[r][c] == 0 {
				return false
			}
			if r < gridSize-1 && g.board[r][c] == g.board[r+1][c] {
				return false
			}
			if c < gridSize-1 && g.board[r][c] == g.board[r][c+1] {
				return false
			}
		}
	}
	return true
}

func (g *Game) PrintBoard() {
	fmt.Println("Score:", g.score)
	for _, row := range g.board {
		for _, val := range row {
			if val == 0 {
				fmt.Print(".\t")
			} else {
				fmt.Printf("%d\t", val)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Game) merge(row [gridSize]int) [gridSize]int {
	newRow := [gridSize]int{}
	j := 0
	for i := 0; i < gridSize; i++ {
		if row[i] != 0 {
			if j < gridSize-1 && newRow[j] == 0 {
				newRow[j] = row[i]
			} else if newRow[j] == row[i] {
				newRow[j] *= 2
				g.score += newRow[j]
				j++
			} else {
				j++
				if j < gridSize {
					newRow[j] = row[i]
				}
			}
		}
	}
	return newRow
}

func reverse(row *[gridSize]int) {
	for i := 0; i < gridSize/2; i++ {
		(*row)[i], (*row)[gridSize-1-i] = (*row)[gridSize-1-i], (*row)[i]
	}
}

func transpose(board [gridSize][gridSize]int) [gridSize][gridSize]int {
	var newBoard [gridSize][gridSize]int
	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			newBoard[c][r] = board[r][c]
		}
	}
	return newBoard
}
