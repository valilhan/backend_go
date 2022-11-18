package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gitlab.com/idoko/HyperSkill/db"
	"gitlab.com/idoko/HyperSkill/models"
)

func items(router chi.Router) {
	router.Post("/fund", addBalance)
	router.Get("/fund", getBalance)
	router.Post("/reserve", addReserveBalance)
	router.Post("/accept", addRevenue)
}

func addBalance(w http.ResponseWriter, r *http.Request) {
	item := &models.Balance{}
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := dbInstance.AddBalance(item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	item := &models.GetBalance{}

	// fetch the json to the item
	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	ans, err := dbInstance.GetBalanceById(item)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrorRenderer(err))
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &ans); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func addReserveBalance(w http.ResponseWriter, r *http.Request) {
	item := &models.ReserveBalance{}

	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := dbInstance.AddReserveBalance(item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func addRevenue(w http.ResponseWriter, r *http.Request) {
	item := &models.Revenue{}

	if err := render.Bind(r, item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := dbInstance.AddRevenue(item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}

}
