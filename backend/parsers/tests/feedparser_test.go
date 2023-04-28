package parsers_test

import (
	"io/ioutil"
	"newsfeed/parsers"
	"testing"
)

func TestFeedParser(t *testing.T) {
	testFileName := "mock_complete.xml"
	data, err := ioutil.ReadFile(testFileName)
	if err != nil {
		panic(err)
	}
	feed, news := parsers.ParseFeed(data)
	t.Log(feed.Title)
	t.Log(news[0].Title)
	if len(news) != 3 {
		t.Errorf("Expected 3 news items, got %d", len(news))
	}
}
