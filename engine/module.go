/*
* Gearbox monitoring
* Copyright (c) 2021. Mikael Lazarev
*
 */

package engine

import (
	"go.uber.org/fx"
)

var Module = fx.Option(
	fx.Provide(NewEngine))
