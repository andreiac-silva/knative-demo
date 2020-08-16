package cloudevent

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
	. "purchase/internal/logger"
	"purchase/pkg/models"
)

type SenderService struct {
	cloudConfig CloudEventConfig
	client      client.Client
}

func NewSenderService(client cloudevents.Client, cloudConfig CloudEventConfig) SenderService {
	return SenderService{
		cloudConfig: cloudConfig,
		client:      client,
	}
}

type Data struct {
	CustomerId string      `json:"customerId"`
	Data       interface{} `json:"data"`
	Headers    string      `json:"headers"`
}

func (s SenderService) Send(purchaseData models.Purchase) (err error) {
	eventObj := cloudevents.NewEvent(event.CloudEventsVersionV1)
	eventObj.SetType(s.cloudConfig.EventType)
	eventObj.SetID(uuid.New().String())

	Logger.Debugf("eventSource: %s", s.cloudConfig.EventSource)
	eventObj.SetSource(s.cloudConfig.EventSource)

	if err := eventObj.SetData(cloudevents.ApplicationJSON, purchaseData); err != nil {
		Logger.Error("failed to set cloud events data: %s", err.Error())
		return err
	}

	if result := s.client.Send(context.TODO(), eventObj); cloudevents.IsUndelivered(result) {
		Logger.Errorf("failed to send, %v", result)
		return result
	}

	defer Logger.Debugf("eventObj sent: %+v", eventObj)
	return nil
}