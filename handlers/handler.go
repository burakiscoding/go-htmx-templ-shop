package handlers

import (
	"net/http"
	"strconv"

	"github.com/burakiscoding/go-htmx-templ/stores"
	"github.com/burakiscoding/go-htmx-templ/templates"
	"github.com/burakiscoding/go-htmx-templ/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *stores.Store
}

func NewHandler(store *stores.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) error {
	products, err := h.store.GetAll()
	if err != nil {
		return err
	}
	return utils.Render(templates.HomePage(products), w, r)
}

func (h *Handler) HandleProduct(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}

	product, err := h.store.GetById(id)
	if err != nil {
		return err
	}

	return utils.Render(templates.Product(product), w, r)
}

func (h *Handler) HandleEdit(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}

	product, err := h.store.GetById(id)
	if err != nil {
		return err
	}

	return utils.Render(templates.ProductEditForm(product), w, r)
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}

	name := r.FormValue("name")
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		return err
	}

	unitPrice, err := strconv.ParseFloat(r.FormValue("unit_price"), 32)
	if err != nil {
		return err
	}

	product, err := h.store.Update(id, name, quantity, float32(unitPrice))
	if err != nil {
		return err
	}

	return utils.Render(templates.Product(product), w, r)
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err
	}

	if err := h.store.Delete(id); err != nil {
		return err
	}

	return nil
}

func (h *Handler) HandleShowAddForm(w http.ResponseWriter, r *http.Request) error {
	return utils.Render(templates.ProductAddForm(), w, r)
}

func (h *Handler) HandleAdd(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		return err
	}

	unitPrice, err := strconv.ParseFloat(r.FormValue("unit_price"), 32)
	if err != nil {
		return err
	}

	product, err := h.store.Add(name, quantity, float32(unitPrice))
	if err != nil {
		return err
	}

	return utils.Render(templates.Product(product), w, r)
}

func (h *Handler) HandleNoContent(w http.ResponseWriter, r *http.Request) error {
	return utils.Render(templates.NoContent(), w, r)
	// w.WriteHeader(http.StatusNoContent)
	// return nil
}
