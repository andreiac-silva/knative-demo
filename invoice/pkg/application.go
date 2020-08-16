package pkg

import (
	"invoice/pkg/handler"
)

type Application struct {
	CloudEventHandler handler.CloudEventHandler
}

func NewApplication(CloudEventHandler handler.CloudEventHandler) Application {
	return Application{CloudEventHandler: CloudEventHandler}
}
