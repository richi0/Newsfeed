package handlers

import (
	"fmt"
	"net/http"
	"newsfeed/config"
	"newsfeed/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type NewsHandlers struct {
	app  *config.Application
	data *models.Models
}

func NewNewsHandlers(app *config.Application, data *models.Models) *NewsHandlers {
	return &NewsHandlers{app, data}
}

func (n *NewsHandlers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var news models.News
	err := readJson(w, r, &news)
	if err != nil {
		return
	}
	dbNews := n.data.News.Create(&news)
	writeJson(w, dbNews)
}

func (n *NewsHandlers) Read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	limit := getQueryParamInt("limit", 100, r)
	offset := getQueryParamInt("offset", 0, r)
	dbNews := n.data.News.Read(limit, offset)
	writeJson(w, dbNews)
}

func (n *NewsHandlers) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var news models.News
	err := readJson(w, r, &news)
	if err != nil {
		return
	}
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "news id not provided", http.StatusInternalServerError)
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "news id not a number", http.StatusInternalServerError)
		return
	}
	news.ID = uint(idUint)
	dbNews := n.data.News.Update(&news)
	writeJson(w, dbNews)
}

func (n *NewsHandlers) Remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "news id not provided", http.StatusInternalServerError)
		return
	}
	n.data.News.Remove(id)
	w.Write([]byte(fmt.Sprintf("news %s deleted", id)))
}

func (n *NewsHandlers) Drop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	n.data.News.Drop()
	w.Write([]byte("dropped news table"))
}
