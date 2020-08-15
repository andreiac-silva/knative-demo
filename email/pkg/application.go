package pkg

import (
	"email/pkg/handler"
)

type Application struct {
	CloudEventHandler handler.CloudEventHandler
}

func NewApplication(CloudEventHandler handler.CloudEventHandler) Application {
	return Application{CloudEventHandler: CloudEventHandler}
}
