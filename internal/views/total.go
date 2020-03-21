package views

import (
	"context"
	"github.com/go-po/example-teller/internal/domain"
	"github.com/go-po/po/stream"
)

type VariableTotal struct {
	Total int64
}

func (view *VariableTotal) Handle(ctx context.Context, msg stream.Message) error {
	switch event := msg.Data.(type) {
	case domain.SubtractedEvent:
		view.Total = view.Total - event.Value
	case domain.AddedEvent:
		view.Total = view.Total + event.Value
	}
	return nil
}
