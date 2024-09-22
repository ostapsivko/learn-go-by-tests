package poker

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	store   PlayerStore
	input   *bufio.Scanner
	alerter BlindAlerter
	output  io.Writer
}

func NewCLI(store PlayerStore, input io.Reader, output io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		store:   store,
		input:   bufio.NewScanner(input),
		alerter: alerter,
		output:  output,
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.output, PlayerPrompt)
	c.scheduleBlindAlerts()
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

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}
