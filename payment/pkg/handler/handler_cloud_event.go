package handler

import (
	"context"
	"github.com/Pallinder/go-randomdata"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"net/http"
	. "payment/internal/logger"
	"strings"
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
	Logger.Debug("**** Payment Service **** ")

	// converting event data (payload) to a map json structure
	var purchase map[string]interface{}
	if err := event.DataAs(&purchase); err != nil {
		Logger.Errorw("Error while extracting cloud event data: ", err.Error())
		return &event, cloudevents.NewHTTPResult(http.StatusUnprocessableEntity, "Error while extracting cloud event data", err)
	}

	// completing event payload with some random data
	loadCustomerInfo(purchase["customer"].(map[string]interface{}))

	// extensions are headers
	event.SetExtension("status", "approved")

	// setting new event payload
	if err := event.SetData(cloudevents.ApplicationJSON, purchase); err != nil {
		Logger.Errorw("Failed setting event data: ", err.Error())
		return &event, cloudevents.NewHTTPResult(http.StatusInternalServerError, "Failed setting event data", err)
	}

	Logger.Debug("Payment approved!")

	// returning event to the flow
	return &event, nil
}

func loadCustomerInfo(customer map[string]interface{}) {
	name := randomdata.FullName(randomdata.RandomGender)
	customer["name"] = name
	customer["email"] = strings.ToLower(strings.Replace(name, " ", ".", -1)) + "@gophers.com"

	Logger.Debug("Processing payment for customer " + name + " ...")
}
