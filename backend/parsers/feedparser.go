package parsers

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"newsfeed/models"
	"strings"
)

func getCategoryString(category []Category) string {
	data, err := json.Marshal(category)
	if err != nil {
		log.Panicf("Cannot marshal category. Error: %s", err)
	}
	categoryString := string(data)
	categoryString = strings.ReplaceAll(categoryString, "\"Domain\"", "\"domain\"")
	categoryString = strings.ReplaceAll(categoryString, "\"CharData\"", "\"text\"")
	return categoryString
}

func getSkipHoursString(hours SkipHours) string {
	data, err := json.Marshal(hours.Hour)
	if err != nil {
		log.Panicf("Cannot marshal skip hours. Error: %s", err)
	}
	hoursString := string(data)
	return hoursString
}

func getSkipDaysString(days SkipDays) string {
	data, err := json.Marshal(days.Day)
	if err != nil {
		log.Panicf("Cannot marshal skip days. Error: %s", err)
	}
	daysString := string(data)
	return daysString
}

func ParseFeed(data []byte) (*models.Feed, []*models.News) {
	var rss Rss
	xml.Unmarshal(data, &rss)
	news := make([]*models.News, 0)
	for _, item := range rss.Channel.Items {
		newsItem := &models.News{
			Title:           item.Title,
			Link:            item.Link,
			Description:     item.Description,
			PubDate:         item.PubDate,
			Category:        getCategoryString(item.Category),
			Author:          item.Author,
			Comments:        item.Comments,
			EnclosureUrl:    item.Enclosure.Url,
			EnclosureLength: item.Enclosure.Length,
			EnclosureType:   item.Enclosure.Type,
			GuidUrl:         item.Guid.CharData,
			GuidIsPermaLink: item.Guid.IsPermaLink,
			SourceUrl:       item.Source.Url,
			SourceText:      item.Source.CharData,
		}
		news = append(news, newsItem)
	}
	feed := &models.Feed{
		Title:                  rss.Channel.Title,
		Link:                   rss.Channel.Link,
		Description:            rss.Channel.Description,
		Language:               rss.Channel.Language,
		Copyright:              rss.Channel.Copyright,
		ManagingEditor:         rss.Channel.ManagingEditor,
		WebMaster:              rss.Channel.WebMaster,
		PubDate:                rss.Channel.PubDate,
		LastBuildDate:          rss.Channel.LastBuildDate,
		Category:               getCategoryString(rss.Channel.Category),
		Generator:              rss.Channel.Generator,
		Docs:                   rss.Channel.Docs,
		CloudDomain:            rss.Channel.Cloud.Domain,
		CloudPort:              rss.Channel.Cloud.Port,
		CloudPath:              rss.Channel.Cloud.Path,
		CloudRegisterProcedure: rss.Channel.Cloud.RegisterProcedure,
		CloudProtocol:          rss.Channel.Cloud.Protocol,
		TTL:                    rss.Channel.Ttl,
		ImageUrl:               rss.Channel.Image.Url,
		ImageTitle:             rss.Channel.Image.Title,
		ImageLink:              rss.Channel.Image.Link,
		ImageWidth:             rss.Channel.Image.Width,
		ImageHeight:            rss.Channel.Image.Height,
		ImageDescription:       rss.Channel.Image.Description,
		Rating:                 rss.Channel.Rating,
		TextInputTitle:         rss.Channel.TextInput.Title,
		TextInputDescription:   rss.Channel.TextInput.Description,
		TextInputName:          rss.Channel.TextInput.Name,
		TextInputLink:          rss.Channel.TextInput.Link,
		SkipHours:              getSkipHoursString(rss.Channel.SkipHours),
		SkipDays:               getSkipDaysString(rss.Channel.SkipDays),
	}
	return feed, news
}
