package pkg

import (
	"github.com/google/wire"
	"email/internal/cloud"
	"email/pkg/handler"
)

var Container = wire.NewSet(
	cloud.CloudEventClientSet,
	handler.CloudEventHandlerSet,
	ApplicationSet,
)
