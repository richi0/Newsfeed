package models

import (
	"newsfeed/config"
	"time"

	"gorm.io/gorm"
)

type FeedModel struct {
	app *config.Application
}

type Feed struct {
	gorm.Model
	ID                     uint   `gorm:"primaryKey"`
	Url                    string `gorm:"notNull;unique"`
	Title                  string
	Link                   string
	Description            string
	Language               string
	Copyright              string
	ManagingEditor         string
	WebMaster              string
	PubDate                string
	LastBuildDate          string
	Category               string
	Generator              string
	Docs                   string
	CloudDomain            string
	CloudPort              uint
	CloudPath              string
	CloudRegisterProcedure string
	CloudProtocol          string
	TTL                    uint
	ImageUrl               string
	ImageTitle             string
	ImageLink              string
	ImageWidth             uint
	ImageHeight            uint
	ImageDescription       string
	Rating                 string
	TextInputTitle         string
	TextInputDescription   string
	TextInputName          string
	TextInputLink          string
	SkipHours              string
	SkipDays               string
	LastFetched            time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              gorm.DeletedAt
}

func (f *FeedModel) Create(feed *Feed) *Feed {
	var newFeed Feed
	f.app.DB.Create(feed)
	f.app.DB.First(&newFeed, feed.ID)
	return &newFeed
}

func (f *FeedModel) Read(limit int, offset int) []*Feed {
	var feeds []*Feed
	f.app.DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&feeds)
	return feeds
}

func (f *FeedModel) Update(feed *Feed) *Feed {
	f.app.DB.Save(feed)
	return feed
}

func (f *FeedModel) Remove(id string) {
	f.app.DB.Delete(&Feed{}, id)
}

func (f *FeedModel) Drop() {
	f.app.DB.Migrator().DropTable(&Feed{})
}
