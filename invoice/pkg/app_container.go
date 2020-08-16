package pkg

import (
	"invoice/internal/cloud"
	"invoice/pkg/handler"
	"github.com/google/wire"
)

var Container = wire.NewSet(
	cloud.CloudEventClientSet,
	handler.CloudEventHandlerSet,
	ApplicationSet,
)
