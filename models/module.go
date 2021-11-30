/*
* Gearbox monitoring
* Copyright (c) 2021. Mikael Lazarev
*
 */

package models

import (
	"github.com/Gearbox-protocol/third-eye/models/address_provider"
	"go.uber.org/fx"
)

var Module = fx.Option(
	fx.Provide(address_provider.NewAddressProvider),
)
