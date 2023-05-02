package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"newsfeed/models"
	"testing"
)

func TestReadEmptyNews(t *testing.T) {
	deleteAllNews()
	news := getAllNews()
	if len(news) != 0 {
		t.Error("DB should be empty")
	}
}

func TestReadNewsByID(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	feedID := getAllFeeds()[0].ID
	deleteAllNews()
	addNews(NEWS_1, feedID)
	addNews(NEWS_2, feedID)
	addNews(NEWS_3, feedID)
	id := getAllNews()[1].ID
	news := readNews(fmt.Sprint(id))
	if news.Title != NEWS_2.Title {
		t.Error("News2 title wrong")
	}
}

func TestReadNewsByFeedID(t *testing.T) {
	var news []models.News
	deleteAllFeed()
	addFeed(FEED_1)
	feedID := getAllFeeds()[0].ID
	deleteAllNews()
	addNews(NEWS_1, feedID)
	addNews(NEWS_2, feedID)
	addNews(NEWS_3, feedID)
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL+"news_by_feed_id/"+fmt.Sprint(feedID), nil)
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
	if len(news) != 3 {
		t.Error("Feed should contain three news items")
	}
}

func TestCreateNews(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	feedID := getAllFeeds()[0].ID
	deleteAllNews()
	addNews(NEWS_1, feedID)
	news := getAllNews()
	if len(news) != 1 {
		t.Error("DB should contain one feed")
	}
	newsItem := news[0]
	if newsItem.Title != NEWS_1.Title {
		t.Error("News title wrong")
	}
	if newsItem.Link != NEWS_1.Link {
		t.Error("News link wrong")
	}
	if newsItem.FeedID != NEWS_1.FeedID {
		t.Error("News feedID wrong")
	}
	addNews(NEWS_2, feedID)
	addNews(NEWS_3, feedID)
	news = getAllNews()
	if len(news) != 3 {
		t.Error("DB should contain three news items")
	}
	if news[2].Title != NEWS_1.Title {
		t.Error("News1 title wrong")
	}
	if news[1].Title != NEWS_2.Title {
		t.Error("News2 title wrong")
	}
	if news[0].Title != NEWS_3.Title {
		t.Error("News3 title wrong")
	}
}

func TestUpdateNews(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	feedID := getAllFeeds()[0].ID
	deleteAllNews()
	addNews(NEWS_1, feedID)
	news := getAllNews()[0]
	if news.Title != NEWS_1.Title {
		t.Error("News1 title wrong")
	}
	news.Title = "New Title"
	data, err := json.Marshal(news)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", BASE_URL+"news/"+fmt.Sprint(news.ID), bytes.NewBuffer(data))
	req.SetBasicAuth("user", "password")
	req.Header.Add("Content-Type", "application/json")
	client.Do(req)
	news = getAllNews()[0]
	if news.Title != "New Title" {
		t.Error("Feed1 title wrong")
	}
}

func TestRemoveNews(t *testing.T) {
	deleteAllFeed()
	addFeed(FEED_1)
	feedID := getAllFeeds()[0].ID
	deleteAllNews()
	addNews(NEWS_1, feedID)
	addNews(NEWS_2, feedID)
	addNews(NEWS_3, feedID)
	news := getAllNews()
	if len(news) != 3 {
		t.Error("DB should contain three news")
	}
	deleteNews(fmt.Sprint(news[1].ID))
	news = getAllNews()
	if len(news) != 2 {
		t.Error("DB should contain two news")
	}
	if news[0].Title != NEWS_3.Title {
		t.Error("News1 title wrong")
	}
	if news[1].Title != NEWS_1.Title {
		t.Error("News2 title wrong")
	}
}
