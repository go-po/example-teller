package views

import (
	"context"
	"github.com/go-po/example-teller/internal/domain"
	"github.com/go-po/po/stream"
)

type VariableNames struct {
	Names []string
}

func (view *VariableNames) Handle(ctx context.Context, message stream.Message) error {
	declared, ok := message.Data.(domain.DeclareCommand)
	if !ok {
		return nil
	}
	view.Names = append(view.Names, declared.Name)
	return nil
}
