package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	input  *bufio.Scanner
	output io.Writer
	game   Game
}

func NewCLI(input io.Reader, output io.Writer, game Game) *CLI {
	return &CLI{
		input:  bufio.NewScanner(input),
		output: output,
		game:   game,
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.output, PlayerPrompt)

	numberOfPlayersInput := c.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	c.game.Finish(extractWinner(winnerInput))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.input.Scan()
	return c.input.Text()
}
