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
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var BASE_URL = "http://localhost:4000/"

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

func deleteAllFeed() {
	feeds := getAllFeeds()
	for _, feed := range feeds {
		deleteFeed(fmt.Sprint(feed.ID))
	}
	feeds = getAllFeeds()
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

func TestReadEmptyFeed(t *testing.T) {
	deleteAllFeed()
	feeds := getAllFeeds()
	if len(feeds) != 0 {
		t.Error("DB should be empty")
	}
}

func TestReadFeedByID(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	addFeed(FEED_2)
	addFeed(FEED_3)
	feed := readFeed("2")
	if feed.Title != FEED_2.Title {
		t.Error("Feed1 title wrong")
	}
}

func TestCreateFeed(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	feeds := getAllFeeds()
	if len(feeds) != 1 {
		t.Error("DB should contain one feed")
	}
	feed := feeds[0]
	if feed.Title != FEED_1.Title {
		t.Error("Feed title wrong")
	}
	if feed.Link != FEED_1.Link {
		t.Error("Feed link wrong")
	}
	if feed.TTL != FEED_1.TTL {
		t.Error("Feed ttl wrong")
	}
	addFeed(FEED_2)
	addFeed(FEED_3)
	feeds = getAllFeeds()
	if len(feeds) != 3 {
		t.Error("DB should contain three feeds")
	}
	if feeds[0].Title != FEED_3.Title {
		t.Error("Feed1 title wrong")
	}
	if feeds[1].Title != FEED_2.Title {
		t.Error("Feed2 title wrong")
	}
	if feeds[2].Title != FEED_1.Title {
		t.Error("Feed3 title wrong")
	}
}

func TestBadPassword(t *testing.T) {
	res, err := http.Post(BASE_URL+"feeds", "", nil)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Error("No password should return 401")
	}
}

func TestUpdateFeed(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	feed := getAllFeeds()[0]
	if feed.Title != FEED_1.Title {
		t.Error("Feed1 title wrong")
	}
	feed.Title = "New Title"
	data, err := json.Marshal(feed)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", BASE_URL+"feeds/"+fmt.Sprint(feed.ID), bytes.NewBuffer(data))
	req.SetBasicAuth("user", "password")
	req.Header.Add("Content-Type", "application/json")
	client.Do(req)
	feed = getAllFeeds()[0]
	if feed.Title != "New Title" {
		t.Error("Feed1 title wrong")
	}
}

func TestRemoveFeed(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	addFeed(FEED_2)
	addFeed(FEED_3)
	feeds := getAllFeeds()
	if len(feeds) != 3 {
		t.Error("DB should contain three feeds")
	}
	deleteFeed("2")
	feeds = getAllFeeds()
	if len(feeds) != 2 {
		t.Error("DB should contain two feeds")
	}
	if feeds[0].Title != FEED_3.Title {
		t.Error("Feed1 title wrong")
	}
	if feeds[1].Title != FEED_1.Title {
		t.Error("Feed2 title wrong")
	}
}
