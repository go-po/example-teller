package views

import (
	"context"
	"github.com/go-po/po/stream"
)

type CommandCount struct {
	Count int64
}

func (view *CommandCount) Handle(ctx context.Context, message stream.Message) error {
	view.Count = view.Count + 1
	return nil
}
