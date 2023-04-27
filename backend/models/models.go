package models

import (
	"newsfeed/config"
)

type Models struct {
	Feeds *FeedModel
	News  *NewsModel
}

func NewModels(app *config.Application) *Models {
	return &Models{Feeds: &FeedModel{app}, News: &NewsModel{app}}
}
