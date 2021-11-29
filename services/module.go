package services

import (
	"go.uber.org/fx"
)

var Module = fx.Option(
	fx.Provide(NewExecuteParser))