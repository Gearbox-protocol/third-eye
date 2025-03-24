package debts

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(NewDebtEngine)
