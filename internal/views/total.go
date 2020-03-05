package views

import (
	"context"
	"github.com/go-po/po/examples/teller/events"
	"github.com/go-po/po/stream"
)

type VariableTotal struct {
	Total int64
}

func (view *VariableTotal) Handle(ctx context.Context, msg stream.Message) error {
	switch event := msg.Data.(type) {
	case events.SubtractedEvent:
		view.Total = view.Total - event.Value
	case events.AddedEvent:
		view.Total = view.Total + event.Value
	}
	return nil
}
