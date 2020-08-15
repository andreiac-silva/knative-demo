package pkg

import (
	"payment/pkg/handler"
)

type Application struct {
	CloudEventHandler handler.CloudEventHandler
}

func NewApplication(CloudEventHandler handler.CloudEventHandler) Application {
	return Application{CloudEventHandler: CloudEventHandler}
}
