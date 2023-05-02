package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

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
		t.Error("Feed2 title wrong")
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
