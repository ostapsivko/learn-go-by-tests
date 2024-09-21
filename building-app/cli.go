package poker

import (
	"io"
)

type CLI struct {
	store PlayerStore
	input io.Reader
}

func (c *CLI) PlayPoker() {
	c.store.RecordWin("Azdab")
}
