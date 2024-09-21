package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

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
	c.alerter.ScheduleAlertAt(5*time.Second, 100)
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
