package handler

import (
	"encoding/json"
	"net/http"
	. "purchase/internal/logger"
	"purchase/pkg/api/render"
	"purchase/pkg/cloudevent"
	"purchase/pkg/models"
)

type PurchaseHandler struct {
	senderService cloudevent.SenderService
}

func NewHandler(senderService cloudevent.SenderService) PurchaseHandler {
	return PurchaseHandler{senderService: senderService}
}

func (handler *PurchaseHandler) Buy(w http.ResponseWriter, r *http.Request) {
	requestBody, err := handler.extractBody(r)
	if err != nil {
		render.ResponseError(w, err, http.StatusPreconditionFailed)
		return
	}

	err = handler.senderService.Send(requestBody)
	if err != nil {
		Logger.Errorw("Failed to create purchase event", "error", err)
		render.ResponseError(w, err, http.StatusInternalServerError)
		return
	}
	Logger.Debug("Purchase event has created successfully")
	render.Response(w, "", http.StatusCreated)
}

func (handler *PurchaseHandler) extractBody(r *http.Request) (models.Purchase, error) {
	var purchaseStruct models.Purchase
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()

	if err := decoder.Decode(&purchaseStruct); err != nil {
		Logger.Errorw("Error trying to parse payload", "error", err)
		return models.Purchase{}, err
	}

	return purchaseStruct, nil
}
