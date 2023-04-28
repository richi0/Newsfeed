package parsers

import (
	"encoding/xml"
	"newsfeed/models"
)

func ParseFeed(data []byte) (*models.Feed, []*models.News) {
	var rss Rss
	xml.Unmarshal(data, &rss)
	news := make([]*models.News, 0)
	for _, item := range rss.Channel.Items {
		news = append(news, &models.News{Title: item.Title})
	}
	return &models.Feed{Title: rss.Channel.Title}, news
}
