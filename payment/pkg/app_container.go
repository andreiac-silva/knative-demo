package pkg

import (
	"github.com/google/wire"
	"payment/internal/cloud"
	"payment/pkg/handler"
)

var Container = wire.NewSet(
	cloud.CloudEventClientSet,
	handler.CloudEventHandlerSet,
	ApplicationSet,
)
