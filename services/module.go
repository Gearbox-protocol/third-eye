package services

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(NewAMQPService),
	fx.Provide(NewExecuteParser),
)
