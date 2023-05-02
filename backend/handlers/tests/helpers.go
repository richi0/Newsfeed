package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"newsfeed/config"
	"newsfeed/handlers"
	"newsfeed/models"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func startServerInBackground(app *config.Application, data *models.Models) {
	mux := handlers.RegisterRouts(app, data)
	log.Println("Start server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func init() {
	os.Setenv("BASIC_USER", "user")
	os.Setenv("BASIC_PASSWORD", "password")

	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	app := config.NewApplication(nil, nil, db)
	db.AutoMigrate(&models.Feed{}, &models.News{})

	data := models.NewModels(app)
	go startServerInBackground(app, data)
	time.Sleep(200 * time.Millisecond)
}

func getAllFeeds() []models.Feed {
	var feeds []models.Feed
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"feeds", nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("user", "password")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &feeds)
	return feeds
}

func addFeed(feed *models.Feed) {
	data, err := json.Marshal(feed)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", BASE_URL+"feeds", bytes.NewBuffer(data))
	req.SetBasicAuth("user", "password")
	req.Header.Add("Content-Type", "application/json")
	client.Do(req)
}

func deleteAllFeed() {
	feeds := getAllFeeds()
	for _, feed := range feeds {
		deleteFeed(fmt.Sprint(feed.ID))
	}
}

func deleteFeed(id string) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", BASE_URL+"feeds/"+id, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("user", "password")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func readFeed(id string) models.Feed {
	var feed models.Feed
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"feeds/"+id, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("user", "password")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resData, &feed)
	if err != nil {
		panic(err)
	}
	return feed
}

func getAllNews() []models.News {
	var news []models.News
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"news", nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("user", "password")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &news)
	return news
}

func deleteAllNews() {
	news := getAllNews()
	for _, newsItem := range news {
		deleteNews(fmt.Sprint(newsItem.ID))
	}
}

func deleteNews(id string) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", BASE_URL+"news/"+id, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("user", "password")
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

func addNews(news *models.News, feedID uint) {
	news.FeedID = feedID
	data, err := json.Marshal(news)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("POST", BASE_URL+"news", bytes.NewBuffer(data))
	req.SetBasicAuth("user", "password")
	req.Header.Add("Content-Type", "application/json")
	client.Do(req)
}

func readNews(id string) models.News {
	var news models.News
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"news/"+id, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("user", "password")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resData, &news)
	if err != nil {
		panic(err)
	}
	return news
}
