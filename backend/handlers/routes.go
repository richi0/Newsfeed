package handlers

import (
	"newsfeed/config"
	"newsfeed/middleware"
	"newsfeed/models"

	"github.com/julienschmidt/httprouter"
)

func RegisterRouts(app *config.Application, data *models.Models) *httprouter.Router {
	feeds := NewFeedHandlers(app, data)
	news := NewNewsHandlers(app, data)
	router := httprouter.New()
	router.GET("/feeds", feeds.Read)
	router.POST("/feeds", middleware.BasicAuth(feeds.Create))
	router.POST("/feeds/create_by_url", middleware.BasicAuth(feeds.CreateByFeedUrl))
	router.PUT("/feeds/:id", middleware.BasicAuth(feeds.Update))
	router.DELETE("/feeds/:id", middleware.BasicAuth(feeds.Remove))
	router.DELETE("/feeds_drop", middleware.BasicAuth(feeds.Drop))
	router.GET("/news", news.Read)
	router.POST("/news", middleware.BasicAuth(news.Create))
	router.PUT("/news/:id", middleware.BasicAuth(news.Update))
	router.DELETE("/news/:id", middleware.BasicAuth(news.Remove))
	router.DELETE("/news_drop", middleware.BasicAuth(news.Drop))
	return router
}
