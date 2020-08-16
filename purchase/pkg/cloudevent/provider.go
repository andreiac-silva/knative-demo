package cloudevent

import (
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	. "purchase/internal/logger"
	"log"
)

var SetSenderService = wire.NewSet(NewSenderService, NewCloudEventsClient, NewCloudEventConfig)

type CloudEventConfig struct {
	EventSource string `envconfig:"EVENT_SOURCE" default:"dev.knative.eventing.demo.purchase"`
	EventType   string `envconfig:"EVENT_TYPE" default:"dev.knative.eventing.demo.purchase"`
	Sink        string `envconfig:"K_SINK"`
	Name        string `envconfig:"POD_NAME" required:"true"`
	Namespace   string `envconfig:"POD_NAMESPACE" required:"true"`
	label       string
}

func NewCloudEventConfig() (config CloudEventConfig) {
	if err := envconfig.Process("", &config); err != nil {
		Logger.Fatalf("[ERROR] Failed to process env var: %s", err)
	}

	if config.EventSource == "" {
		config.EventSource = fmt.Sprintf("http://demo.com/purchase/#%s/%s", config.Namespace, config.Name)
	}

	Logger.Debugf("Service eventSource: %s, eventType: %s, sink: %s ", config.EventSource, config.EventType, config.Sink)
	return config
}

func NewCloudEventsClient(cloudEventConfig CloudEventConfig) cloudevents.Client {
	p, err := cloudevents.NewHTTP(cloudevents.WithTarget(cloudEventConfig.Sink))
	if err != nil {
		log.Fatalf("failed to create http protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p, cloudevents.WithUUIDs(), cloudevents.WithTimeNow())
	if err != nil {
		Logger.Fatalf("failed to create cloud events client: %s", err.Error())
	}
	return c
}
