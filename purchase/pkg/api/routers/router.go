package routers

import (
	"github.com/gorilla/mux"
	. "purchase/internal/logger"
	"purchase/pkg/api/handler"
	"net/http"
)

type SystemRoutes struct {
	purchaseHandler handler.PurchaseHandler

}

func (sys *SystemRoutes) SetupHandler() http.Handler {
	r := mux.NewRouter().PathPrefix("/purchase/api/v1").Subrouter()
	r.HandleFunc("/buy", sys.purchaseHandler.Buy).Methods(http.MethodPost, http.MethodOptions)
	return r
}

func NewSystemRoutes(purchaseHandler handler.PurchaseHandler,) SystemRoutes {
	Logger.Info("Creating System Main Routers")
	return SystemRoutes{
		purchaseHandler:    purchaseHandler,
	}
}
