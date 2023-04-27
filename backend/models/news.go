package models

import (
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
	CategoryDomain  string
	CategoryText    string
	Comments        string
	EnclosureUrl    string
	EnclosureLength uint
	EnclosureType   string
	Guid            string
	SourceUrl       string
	SourceText      string
	PubDate         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (n *NewsModel) Create(news *News) *News {
	var newNews News
	n.app.DB.Create(news)
	n.app.DB.First(&newNews, news.ID)
	return &newNews
}

func (n *NewsModel) Read(limit int, offset int) []*News {
	var newsList []*News
	n.app.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&newsList)
	return newsList
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
