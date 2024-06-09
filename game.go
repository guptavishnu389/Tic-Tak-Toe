package main

import "fmt"

type Game struct {
	Grid               [][]string
	Players            []*Player
	CurrentPlayerIndex int
}

// NewTicTakToeGame creates a new Game with given rows and slots per column.
func NewTicTakToeGame(noOfPlayers, numberOfRows, numberOfColumns int) *Game {
	// Initialize the grid with empty strings
	grid := make([][]string, numberOfRows)
	for i := range grid {
		row := make([]string, numberOfColumns)
		for j := 0; j < numberOfColumns; j++ {
			row[j] = "_"
		}
		grid[i] = row
	}

	players := make([]*Player, noOfPlayers)
	game := &Game{
		Grid:    grid,
		Players: players,
	}

	return game
}

func (g *Game) DisplayGrid() {
	fmt.Println("Current Grid:")
	for _, row := range g.Grid {
		fmt.Println(row)
	}
}

func (g *Game) DisplayPlayers() {
	fmt.Println("Current Players:")
	for _, p := range g.Players {
		fmt.Println(p.name, " ", p.mark)
	}
}

func (g *Game) DisplayCurrentPlayerIndex() {
	fmt.Println("Current PlayerIndex:")
	fmt.Println(g.CurrentPlayerIndex)
}

func (g *Game) MakeMove(row, col int) {
	if row < 0 || row >= len(g.Grid) || col < 0 || col >= len(g.Grid[0]) || g.Grid[row][col] != "_" {
		fmt.Println("invalid move. Try again")
		return
	}
	g.Grid[row][col] = g.Players[g.CurrentPlayerIndex].mark

	if g.CheckWin() {
		fmt.Printf("%s has won the game", g.Players[g.CurrentPlayerIndex].name)
	}

	if g.CurrentPlayerIndex == len(g.Players)-1 {
		g.CurrentPlayerIndex = 0
	} else {
		g.CurrentPlayerIndex++
	}

	g.DisplayGrid()

	return
}

// CheckWin checks if a player has won the game.
// winning logic for 3x3 grid game only
func (g *Game) CheckWin() bool {
	// Check rows and columns for winning combinations
	for i := 0; i < len(g.Grid); i++ {
		if g.Grid[i][0] != "_" && g.Grid[i][0] == g.Grid[i][1] && g.Grid[i][1] == g.Grid[i][2] {
			return true // Winning row
		}
		if g.Grid[0][i] != "_" && g.Grid[0][i] == g.Grid[1][i] && g.Grid[1][i] == g.Grid[2][i] {
			return true // Winning column
		}
	}

	// Check diagonals for winning combinations
	if g.Grid[0][0] != "_" && g.Grid[0][0] == g.Grid[1][1] && g.Grid[1][1] == g.Grid[2][2] {
		return true // Diagonal from top-left to bottom-right
	}
	if g.Grid[0][2] != "_" && g.Grid[0][2] == g.Grid[1][1] && g.Grid[1][1] == g.Grid[2][0] {
		return true // Diagonal from top-right to bottom-left
	}

	return false
}
