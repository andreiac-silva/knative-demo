package handler

import (
	"context"
	. "invoice/internal/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CloudEventHandler struct {
	cloudEventClient cloudevents.Client
}

func NewCloudEventHandler(cloudEventsClient cloudevents.Client) CloudEventHandler {
	return CloudEventHandler{cloudEventClient: cloudEventsClient}
}

func (handler *CloudEventHandler) StartReceiver() {
	if err := handler.cloudEventClient.StartReceiver(context.Background(), handler.receive); err != nil {
		Logger.Panic("Something went wrong starting cloud events receiver", "error", err.Error())
	}
}

func (handler *CloudEventHandler) receive(event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	Logger.Debug("**** Invoice Service **** ")
	Logger.Debug("Processing invoice")

	// omitted implementation ...

	Logger.Debug("Invoice successfully issued!")

	// returning event to the flow
	return &event, nil
}
