package handler

import "github.com/google/wire"

// to provide dependency injection which is using wire
var CloudEventHandlerSet = wire.NewSet(NewCloudEventHandler)
