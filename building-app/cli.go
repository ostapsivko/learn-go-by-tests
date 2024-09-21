package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	store   PlayerStore
	input   *bufio.Scanner
	alerter BlindAlerter
}

func NewCLI(store PlayerStore, input io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		store:   store,
		input:   bufio.NewScanner(input),
		alerter: alerter,
	}
}

func (c *CLI) PlayPoker() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Second
	}

	userInput := c.readLine()
	c.store.RecordWin(extractWinner(userInput))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.input.Scan()
	return c.input.Text()
}
