package handler

import (
	"context"
	. "email/internal/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/types"
	"net/http"
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
	Logger.Debug("**** Email Service **** ")

	status, _ := types.ToString(event.Extensions()["status"])

	// converting event data (payload) to a map json structure
	var purchase map[string]interface{}
	if err := event.DataAs(&purchase); err != nil {
		Logger.Errorw("Error while extracting cloud event data: ", err.Error())
		return &event, cloudevents.NewHTTPResult(http.StatusUnprocessableEntity, "Error while extracting cloud event data", err)
	}

	customer := purchase["customer"].(map[string]interface{})
	name, _ := types.ToString(customer["name"])
	msg := "Hello " + name + ". Your payment has been " + status + "."

	Logger.Debug("Mail message: " + msg)

	// omitted implementation ...

	Logger.Debug("Email has been sent!")

	// returning event to the flow
	return &event, nil
}
