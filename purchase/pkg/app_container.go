package pkg

import (
	"github.com/google/wire"
	"purchase/pkg/api/handler"
	"purchase/pkg/api/routers"
	"purchase/pkg/cloudevent"
)

var Container = wire.NewSet(
	handler.ApplicationHandlersSet,
	routers.RoutesSet,
	cloudevent.SetSenderService,
	ApplicationSet,
)
