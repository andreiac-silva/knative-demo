package cloud

import "github.com/google/wire"

// to provide dependency injection which is using wire
var CloudEventClientSet = wire.NewSet(NewCloudEventsClient)
