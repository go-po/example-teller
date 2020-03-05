package views

import (
	"context"
	"github.com/go-po/po/examples/teller/commands"
	"github.com/go-po/po/stream"
)

type VariableNames struct {
	Names []string
}

func (view *VariableNames) Handle(ctx context.Context, message stream.Message) error {
	declared, ok := message.Data.(commands.DeclareCommand)
	if !ok {
		return nil
	}
	view.Names = append(view.Names, declared.Name)
	return nil
}
