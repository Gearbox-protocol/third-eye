package watcher

import (
	"go.uber.org/fx"
)

var Module = fx.Option(
	fx.Invoke(newMetEngine),
)
