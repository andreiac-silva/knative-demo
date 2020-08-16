package routers

import "github.com/google/wire"

var RoutesSet = wire.NewSet(NewSystemRoutes)
