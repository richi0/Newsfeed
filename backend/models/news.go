package models

import (
	"errors"
	"newsfeed/config"
	"time"

	"gorm.io/gorm"
)

type NewsModel struct {
	app *config.Application
}

type News struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	Title           string
	Link            string
	Description     string
	Author          string
	Category        string
	Comments        string
	EnclosureUrl    string
	EnclosureLength uint
	EnclosureType   string
	GuidUrl         string `gorm:"notNull;unique"`
	GuidIsPermaLink string
	SourceUrl       string
	SourceText      string
	PubDate         string
	FeedID          uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (n *NewsModel) Create(news *News) (*News, error) {
	var newNews News
	url := news.GuidUrl
	if url == "" {
		return nil, errors.New("guid url cannot be empty")
	}
	if n.Exists(url) {
		return nil, errors.New("guid url must be unique")
	}
	n.app.DB.Create(news)
	n.app.DB.First(&newNews, news.ID)
	return &newNews, nil
}

func (n *NewsModel) CreateBulk(news []*News, feedID uint) []*News {
	for _, item := range news {
		item.FeedID = feedID
		n.Create(item)
	}
	return news
}

func (n *NewsModel) Read(limit int, offset int) []*News {
	var newsList []*News
	n.app.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&newsList)
	return newsList
}

func (n *NewsModel) ByUrl(url string) *News {
	var news News
	err := n.app.DB.Where("guid_url = ?", url).First(&news).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &news
}

func (n *NewsModel) ByID(id string) *News {
	var news News
	err := n.app.DB.First(&news, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &news
}

func (n *NewsModel) ByFeedID(id string) []News {
	var news []News
	n.app.DB.Where("feed_id = ?", id).Find(&news)
	return news
}

func (n *NewsModel) Update(news *News) *News {
	n.app.DB.Save(news)
	return news
}

func (n *NewsModel) Remove(id string) {
	n.app.DB.Delete(&News{}, id)
}

func (n *NewsModel) Drop() {
	n.app.DB.Migrator().DropTable(&News{})
}

func (n *NewsModel) Exists(url string) bool {
	var news News
	var count int64
	n.app.DB.Model(&news).Where("guid_url = ?", url).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
