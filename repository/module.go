package repository

import (
	"github.com/Gearbox-protocol/third-eye/repository/handlers"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewRepository,
	NewDBClient,
	handlers.NewExtraRepo,
)
