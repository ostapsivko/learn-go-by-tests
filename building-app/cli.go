package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PlayerPrompt         = "Please enter the number of players: "
	BadPlayerInputErrMsg = "bad value received for number of players, please try again with a number"
	BadWinnerInputErrMsg = "bad value received for a winner, please try again in a format '{Player} wins'"
	winnerTrimmingFormat = " wins"
)

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
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprint(c.output, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(c.output, BadWinnerInputErrMsg)
		return
	}

	c.game.Finish(winner)
}

func extractWinner(input string) (string, error) {
	if !strings.Contains(input, winnerTrimmingFormat) {
		return "", errors.New(BadWinnerInputErrMsg)
	}

	winner := strings.Replace(input, winnerTrimmingFormat, "", 1)
	return winner, nil
}

func (c *CLI) readLine() string {
	c.input.Scan()
	return c.input.Text()
}
