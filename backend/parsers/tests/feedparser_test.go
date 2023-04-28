package parsers_test

import (
	"io/ioutil"
	"newsfeed/models"
	"newsfeed/parsers"
	"reflect"
	"testing"
)

var fields = []string{
	"Title",
	"Link",
	"Description",
	"Language",
	"Copyright",
	"ManagingEditor",
	"WebMaster",
	"PubDate",
	"LastBuildDate",
	"Category",
	"Generator",
	"Docs",
	"CloudDomain",
	"CloudPort",
	"CloudPath",
	"CloudRegisterProcedure",
	"CloudProtocol",
	"TTL",
	"ImageUrl",
	"ImageTitle",
	"ImageLink",
	"ImageWidth",
	"ImageHeight",
	"ImageDescription",
	"Rating",
	"TextInputTitle",
	"TextInputDescription",
	"TextInputName",
	"TextInputLink",
	"SkipHours",
	"SkipDays",
}

func getField(v *models.Feed, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func getFeed(path string) *models.Feed {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	feed, _ := parsers.ParseFeed(data)
	return feed
}

func TestFeedParserComplete(t *testing.T) {
	feed := getFeed("mock_complete.xml")
	expectedFeed := &models.Feed{
		Title:                  "Cash - \"Aktuell\" | News",
		Link:                   "https://www.cash.ch/feeds/latest/news",
		Description:            "Cash - \"Aktuell\" | News",
		Language:               "de",
		Copyright:              "Cash - Ringier Axel Springer Schweiz AG",
		ManagingEditor:         "editor@cash.ch",
		WebMaster:              "webmaster@cash.ch",
		PubDate:                "Thu, 27 Apr 2023 18:29:56 +0200",
		LastBuildDate:          "Thu, 27 Apr 2023 18:29:56 +0200",
		Category:               "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"},{\"domain\":\"http://www.cash.com/c2\",\"text\":\"C2\"},{\"domain\":\"http://www.cash.com/c3\",\"text\":\"C3\"}]",
		Generator:              "MightyInHouse Content System v2.3",
		Docs:                   "https://www.rssboard.org/rss-specification",
		CloudDomain:            "rpc.cash.ch",
		CloudPort:              80,
		CloudPath:              "/RPC2",
		CloudRegisterProcedure: "myCash.rssPleaseNotify",
		CloudProtocol:          "xml-rpc",
		TTL:                    60,
		ImageUrl:               "https://www.cash.ch/sites/default/files/media/field_image/2022-08/cash_favicon_96x96.png",
		ImageTitle:             "Cash - \"Aktuell\" | News",
		ImageLink:              "https://www.cash.ch/feeds/latest/news",
		ImageWidth:             96,
		ImageHeight:            96,
		ImageDescription:       "Cash Image",
		Rating:                 "Rating",
		TextInputTitle:         "Cash - \"Aktuell\" | News",
		TextInputDescription:   "Cash Input",
		TextInputName:          "Input",
		TextInputLink:          "https://www.cash.ch/feeds/latest/news",
		SkipHours:              "[0,1,2,3]",
		SkipDays:               "[\"Monday\",\"Tuesday\",\"Wednesday\"]",
	}
	for _, field := range fields {
		expected := getField(expectedFeed, field)
		is := getField(feed, field)
		if is != expected {
			t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
		}
	}
}

func TestFeedParserMissing(t *testing.T) {
	feed := getFeed("mock_missing.xml")
	expectedFeed := &models.Feed{
		Title:                  "Cash - \"Aktuell\" | News",
		Link:                   "https://www.cash.ch/feeds/latest/news",
		Description:            "Cash - \"Aktuell\" | News",
		Language:               "de",
		Copyright:              "Cash - Ringier Axel Springer Schweiz AG",
		ManagingEditor:         "",
		WebMaster:              "",
		PubDate:                "Thu, 27 Apr 2023 18:29:56 +0200",
		LastBuildDate:          "Thu, 27 Apr 2023 18:29:56 +0200",
		Category:               "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"}]",
		Generator:              "",
		Docs:                   "https://www.rssboard.org/rss-specification",
		CloudDomain:            "",
		CloudPort:              0,
		CloudPath:              "",
		CloudRegisterProcedure: "",
		CloudProtocol:          "",
		TTL:                    60,
		ImageUrl:               "https://www.cash.ch/sites/default/files/media/field_image/2022-08/cash_favicon_96x96.png",
		ImageTitle:             "Cash - \"Aktuell\" | News",
		ImageLink:              "https://www.cash.ch/feeds/latest/news",
		ImageWidth:             0,
		ImageHeight:            96,
		ImageDescription:       "Cash Image",
		Rating:                 "Rating",
		TextInputTitle:         "",
		TextInputDescription:   "",
		TextInputName:          "",
		TextInputLink:          "",
		SkipHours:              "[3]",
		SkipDays:               "null",
	}
	for _, field := range fields {
		expected := getField(expectedFeed, field)
		is := getField(feed, field)
		if is != expected {
			t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
		}
	}
}

func TestFeedParserOriginal(t *testing.T) {
	feed := getFeed("mock_org.xml")
	expectedFeed := &models.Feed{
		Title:                  "Cash - \"Aktuell\" | News",
		Link:                   "https://www.cash.ch/feeds/latest/news",
		Description:            "Cash - \"Aktuell\" | News",
		Language:               "de",
		Copyright:              "Cash - Ringier Axel Springer Schweiz AG",
		ManagingEditor:         "",
		WebMaster:              "",
		PubDate:                "Thu, 27 Apr 2023 18:29:56 +0200",
		LastBuildDate:          "",
		Category:               "null",
		Generator:              "",
		Docs:                   "",
		CloudDomain:            "",
		CloudPort:              0,
		CloudPath:              "",
		CloudRegisterProcedure: "",
		CloudProtocol:          "",
		TTL:                    60,
		ImageUrl:               "https://www.cash.ch/sites/default/files/media/field_image/2022-08/cash_favicon_96x96.png",
		ImageTitle:             "Cash - \"Aktuell\" | News",
		ImageLink:              "https://www.cash.ch/feeds/latest/news",
		ImageWidth:             0,
		ImageHeight:            96,
		ImageDescription:       "",
		Rating:                 "",
		TextInputTitle:         "",
		TextInputDescription:   "",
		TextInputName:          "",
		TextInputLink:          "",
		SkipHours:              "null",
		SkipDays:               "null",
	}
	for _, field := range fields {
		expected := getField(expectedFeed, field)
		is := getField(feed, field)
		if is != expected {
			t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
		}
	}
}
