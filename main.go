package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var game *Game

	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		if len(input) == 0 {
			continue
		}

		switch input[0] {
		case "start_game":
			if len(input) != 4 {
				fmt.Println("Invalid command. Correct format: start_game <no_of_players> <no_of_rows> <no_of_columns>")
				continue
			}
			// Additional error checking for input[2] and input[3] needed
			// Inside the case "start_game":
			noOfPlayers, err := parseInt(input[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			noOfRows, err := parseInt(input[2])
			if err != nil {
				fmt.Println(err)
				continue
			}

			noOfColumns, err := parseInt(input[3])
			if err != nil {
				fmt.Println(err)
				continue
			}

			game = NewTicTakToeGame(noOfPlayers, noOfRows, noOfColumns)
			fmt.Printf("Created TicTakToe Game with %d rows and %d colums\n", noOfRows, noOfColumns)

			// Take player names and marks
			game.Players = make([]*Player, noOfPlayers)
			for i := 0; i < noOfPlayers; i++ {
				fmt.Printf("Enter name and mark for player %d (separated by a space): ", i+1)
				if scanner.Scan() {
					playerInput := strings.Fields(scanner.Text())
					if len(playerInput) != 2 {
						fmt.Println("Invalid input. Enter name and mark separated by a space.")
						i-- // Decrement i to retry the input for the current player
						continue
					}
					playerName := strings.TrimSpace(playerInput[0])
					playerMark := strings.TrimSpace(playerInput[1])
					game.Players[i] = &Player{name: playerName, mark: playerMark}
				}
			}

		case "make_move":
			if len(input) != 3 {
				fmt.Println("Invalid command. Correct format: make_move <row> <col>")
				continue
			}
			// Additional error checking for input[2] and input[3] needed
			// Inside the case "start_game":
			row, err := parseInt(input[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			col, err := parseInt(input[2])
			if err != nil {
				fmt.Println(err)
				continue
			}

			game.MakeMove(row, col)

		case "display_grid":
			game.DisplayGrid()

		case "display_players":
			game.DisplayPlayers()

		case "display_current_player_index":
			game.DisplayCurrentPlayerIndex()

		case "exit":
			fmt.Println("Exiting")
			return

		default:
			fmt.Println("Invalid command")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func parseInt(s string) (int, error) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error parsing '%s': %v", s, err)
	}
	return value, nil
}
