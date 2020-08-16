package pkg

import (
	"purchase/pkg/api/routers"
)

type Application struct {
	SystemRoutes routers.SystemRoutes
}

func NewApplication(SystemRoutes routers.SystemRoutes) Application {
	return Application{SystemRoutes: SystemRoutes}
}
