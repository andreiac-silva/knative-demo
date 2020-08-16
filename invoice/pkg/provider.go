package pkg

import "github.com/google/wire"

// to provide dependency injection which is using wire
var ApplicationSet = wire.NewSet(NewApplication)
