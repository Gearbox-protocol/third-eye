/*
* Gearbox monitoring
* Copyright (c) 2021. Mikael Lazarev
*
*/

package models

import (
	"go.uber.org/fx"
	"github.com/Gearbox-protocol/gearscan/models/address_provider" 
)
 
var Module = fx.Option(
	fx.Provide(address_provider.NewAddressProvider),
)
 