package api

import (
	"net/http"

	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/services"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/square/go-jose.v2/json"
)

type Api struct {
	finderService services.Finder
}

func (api *Api) Handlers(router *httprouter.Router) {
	router.POST("/origins/:origin/destinations/:destination/values/:value", api.insertRoute)
	router.GET("/origins/:origin/destinations/:destination", api.bestPrice)
}

func (api *Api) insertRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !validateParams(ps) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	origin := ps.ByName("origin")
	destination := ps.ByName("destination")
	value := ps.ByName("value")

	record := entity.Record{
		Origin:      origin,
		Destination: destination,
		Value:       value,
	}
	if err := api.finderService.Save(record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (api *Api) bestPrice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	origin := ps.ByName("origin")
	destination := ps.ByName("destination")

	if !validateParams(ps) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cheapestRoute, err := api.finderService.GetSmallestPriceAndRoute(origin, destination)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, err := json.Marshal(cheapestRoute)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func validateParams(params httprouter.Params) bool {
	for _, param := range params {
		if param.Value == "" {
			return false
		}
	}
	return true
}

func (api *Api) Start() {
	router := httprouter.New()
	api.Handlers(router)
	http.ListenAndServe(":8080", router)
}

func New(finderService services.Finder) *Api {
	return &Api{
		finderService: finderService,
	}
}
