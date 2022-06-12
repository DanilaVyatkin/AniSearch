package cartoon

import (
	"MakeAnAPI/internal/cartoon/apperror"
	"MakeAnAPI/internal/handlers"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	CartoonsID = "/anime"
	cartoonID  = "/anime/:id"
)

type handler struct {
	Storage Storage
}

func NewHandler(Storage Storage) handlers.Handler {
	return &handler{
		Storage: Storage,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, CartoonsID, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, cartoonID, apperror.Middleware(h.GetListByID))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	all, err := h.Storage.FindAll(context.TODO())

	if err != nil {
		return err
	}

	allBytes, err := json.Marshal(all)
	if err != nil {
		log.Println(err)
	}

	w.Write(allBytes)

	return nil
}

func (h *handler) GetListByID(w http.ResponseWriter, e *http.Request) error {
	findOne, err := h.Storage.FindOne(context.TODO(), cartoonID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(findOne)
	return nil
}
