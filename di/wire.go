//go:build wireinject

package di

import (
	"go-wire-tutorial/model"

	"github.com/google/wire"
)

func InitializeEvent() model.Event {
	wire.Build(model.NewEvent, model.NewGreeter, model.NewMessage)
	return model.Event{}
}
