package cloud

import (
	 . "invoice/internal/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func NewCloudEventsClient() (cloudevents.Client, error) {
	client, err := cloudevents.NewDefaultClient()

	if err != nil {
		Logger.Errorw("Something went wrong creating cloud events client", "error", err.Error())
		return nil, err
	}

	Logger.Info("Cloud events client was created")
	return client, nil
}
