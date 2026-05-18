package mail

import (
	"github.com/0kaba0hub/go-message"
)

// A InlineHeader represents a message text header.
type InlineHeader struct {
	message.Header
}

var _ PartHeader = (*InlineHeader)(nil)

func (*InlineHeader) partHeader() {}
