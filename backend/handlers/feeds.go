package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"newsfeed/config"
	"newsfeed/models"
	"newsfeed/parsers"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type FeedHandlers struct {
	app  *config.Application
	data *models.Models
}

func NewFeedHandlers(app *config.Application, data *models.Models) *FeedHandlers {
	return &FeedHandlers{app, data}
}

func (f *FeedHandlers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var feed models.Feed
	err := readJson(w, r, &feed)
	if err != nil {
		return
	}
	dbFeed, err := f.data.Feeds.Create(&feed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, dbFeed)
}

func (f *FeedHandlers) CreateByFeedUrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := getQueryParamString("url", r)
	if url == "" {
		http.Error(w, "feed url not provided", http.StatusInternalServerError)
		return
	}
	res, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("feed fetch failed. %s", err), http.StatusInternalServerError)
		return
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not reed feed body. %s", err), http.StatusInternalServerError)
		return
	}
	feed, _ := parsers.ParseFeed(content)
	feed.Url = url
	dbFeed, err := f.data.Feeds.Create(feed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, dbFeed)
}

func (f *FeedHandlers) Read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	limit := getQueryParamInt("limit", 100, r)
	offset := getQueryParamInt("offset", 0, r)
	dbFeeds := f.data.Feeds.Read(limit, offset)
	writeJson(w, dbFeeds)
}

func (f *FeedHandlers) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var feed models.Feed
	err := readJson(w, r, &feed)
	if err != nil {
		return
	}
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "feed id not provided", http.StatusInternalServerError)
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		http.Error(w, "feed id not a number", http.StatusInternalServerError)
		return
	}
	feed.ID = uint(idUint)
	dbFeed := f.data.Feeds.Update(&feed)
	writeJson(w, dbFeed)
}

func (f *FeedHandlers) Remove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "feed id not provided", http.StatusInternalServerError)
		return
	}
	f.data.Feeds.Remove(id)
	w.Write([]byte(fmt.Sprintf("feed %s deleted", id)))
}

func (f *FeedHandlers) Drop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	f.data.Feeds.Drop()
	w.Write([]byte("dropped feed table"))
}
