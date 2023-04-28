package parsers_test

import (
	"fmt"
	"io/ioutil"
	"newsfeed/models"
	"newsfeed/parsers"
	"reflect"
	"testing"
)

var feedFields = []string{
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

var newsFields = []string{
	"Title",
	"Link",
	"Description",
	"Author",
	"Category",
	"Comments",
	"EnclosureUrl",
	"EnclosureLength",
	"EnclosureType",
	"GuidUrl",
	"GuidIsPermaLink",
	"SourceUrl",
	"SourceText",
	"PubDate",
}

func getFeedField(v *models.Feed, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func getNewsField(v *models.News, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	if field == "EnclosureLength" {
		return fmt.Sprint(f.Uint())
	}
	return f.String()
}

func getFeed(path string) (*models.Feed, []*models.News) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	feed, news := parsers.ParseFeed(data)
	return feed, news
}

func TestFeedParserComplete(t *testing.T) {
	feed, _ := getFeed("mock_complete.xml")
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
	for _, field := range feedFields {
		expected := getFeedField(expectedFeed, field)
		is := getFeedField(feed, field)
		if is != expected {
			t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
		}
	}
}

func TestFeedParserMissing(t *testing.T) {
	feed, _ := getFeed("mock_missing.xml")
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
	for _, field := range feedFields {
		expected := getFeedField(expectedFeed, field)
		is := getFeedField(feed, field)
		if is != expected {
			t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
		}
	}
}

func TestFeedParserOriginal(t *testing.T) {
	feed, _ := getFeed("mock_org.xml")
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
	for _, field := range feedFields {
		expected := getFeedField(expectedFeed, field)
		is := getFeedField(feed, field)
		if is != expected {
			t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
		}
	}
}

func TestNewsParserComplete(t *testing.T) {
	_, newsItems := getFeed("mock_complete.xml")
	expectedNews := []*models.News{
		{
			Title:           "1 Börsen-Ticker :  SMI schlägt Seitwärts-Kurs ein",
			Link:            "1 https://www.cash.ch/news/top-news/borsen-ticker-smi-schlagt-seitwarts-kurs-ein-596719",
			Description:     "1 +++Märkte+++ - Die Schweizer Börse hat am Donnerstag weder nachgegeben, noch dazu gewonnen.",
			PubDate:         "1 Thu, 27 Apr 2023 06:09:20 +0200",
			Category:        "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"},{\"domain\":\"http://www.cash.com/c2\",\"text\":\"C2\"},{\"domain\":\"http://www.cash.com/c3\",\"text\":\"C3\"}]",
			Author:          "1 author@cash.ch",
			Comments:        "1 http://www.cash.ch/comments",
			EnclosureUrl:    "1 https://www.cash.ch/fp/1200/675/1910/1089/sites/default/files/images/library/six_boerse_zuerich_aktien.jpg",
			EnclosureLength: 11816035,
			EnclosureType:   "1 image/jpeg",
			GuidUrl:         "1 https://www.cash.ch/node/573556",
			GuidIsPermaLink: "true",
			SourceUrl:       "1 http://www.cash.ch/links2.xml",
			SourceText:      "1 Cash Rss",
		},
		{
			Title:           "2 Börsen-Ticker :  SMI schlägt Seitwärts-Kurs ein",
			Link:            "2 https://www.cash.ch/news/top-news/borsen-ticker-smi-schlagt-seitwarts-kurs-ein-596719",
			Description:     "2 +++Märkte+++ - Die Schweizer Börse hat am Donnerstag weder nachgegeben, noch dazu gewonnen.",
			PubDate:         "2 Thu, 27 Apr 2023 06:09:20 +0200",
			Category:        "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"},{\"domain\":\"http://www.cash.com/c2\",\"text\":\"C2\"},{\"domain\":\"http://www.cash.com/c3\",\"text\":\"C3\"}]",
			Author:          "2 author@cash.ch",
			Comments:        "2 http://www.cash.ch/comments",
			EnclosureUrl:    "2 https://www.cash.ch/fp/1200/675/1910/1089/sites/default/files/images/library/six_boerse_zuerich_aktien.jpg",
			EnclosureLength: 21816035,
			EnclosureType:   "2 image/jpeg",
			GuidUrl:         "2 https://www.cash.ch/node/573556",
			GuidIsPermaLink: "true",
			SourceUrl:       "2 http://www.cash.ch/links2.xml",
			SourceText:      "2 Cash Rss",
		},
		{
			Title:           "3 Börsen-Ticker :  SMI schlägt Seitwärts-Kurs ein",
			Link:            "3 https://www.cash.ch/news/top-news/borsen-ticker-smi-schlagt-seitwarts-kurs-ein-596719",
			Description:     "3 +++Märkte+++ - Die Schweizer Börse hat am Donnerstag weder nachgegeben, noch dazu gewonnen.",
			PubDate:         "3 Thu, 27 Apr 2023 06:09:20 +0200",
			Category:        "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"},{\"domain\":\"http://www.cash.com/c2\",\"text\":\"C2\"},{\"domain\":\"http://www.cash.com/c3\",\"text\":\"C3\"}]",
			Author:          "3 author@cash.ch",
			Comments:        "3 http://www.cash.ch/comments",
			EnclosureUrl:    "3 https://www.cash.ch/fp/1200/675/1910/1089/sites/default/files/images/library/six_boerse_zuerich_aktien.jpg",
			EnclosureLength: 31816035,
			EnclosureType:   "3 image/jpeg",
			GuidUrl:         "3 https://www.cash.ch/node/573556",
			GuidIsPermaLink: "false",
			SourceUrl:       "3 http://www.cash.ch/links2.xml",
			SourceText:      "3 Cash Rss",
		},
	}
	for _, field := range newsFields {
		for i, item := range newsItems {
			expected := getNewsField(item, field)
			is := getNewsField(expectedNews[i], field)
			if is != expected {
				t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
			}
		}
	}
}

func TestNewsParserMissing(t *testing.T) {
	_, newsItems := getFeed("mock_missing.xml")
	expectedNews := []*models.News{
		{
			Title:           "1 Börsen-Ticker :  SMI schlägt Seitwärts-Kurs ein",
			Link:            "1 https://www.cash.ch/news/top-news/borsen-ticker-smi-schlagt-seitwarts-kurs-ein-596719",
			Description:     "1 +++Märkte+++ - Die Schweizer Börse hat am Donnerstag weder nachgegeben, noch dazu gewonnen.",
			PubDate:         "1 Thu, 27 Apr 2023 06:09:20 +0200",
			Category:        "null",
			Author:          "1 author@cash.ch",
			Comments:        "",
			EnclosureUrl:    "",
			EnclosureLength: 11,
			EnclosureType:   "1 image/jpeg",
			GuidUrl:         "1 https://www.cash.ch/node/573556",
			GuidIsPermaLink: "true",
			SourceUrl:       "1 http://www.cash.ch/links2.xml",
			SourceText:      "1 Cash Rss",
		},
		{
			Title:           "2 Börsen-Ticker :  SMI schlägt Seitwärts-Kurs ein",
			Link:            "2 https://www.cash.ch/news/top-news/borsen-ticker-smi-schlagt-seitwarts-kurs-ein-596719",
			Description:     "2 +++Märkte+++ - Die Schweizer Börse hat am Donnerstag weder nachgegeben, noch dazu gewonnen.",
			PubDate:         "2 Thu, 27 Apr 2023 06:09:20 +0200",
			Category:        "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"},{\"domain\":\"http://www.cash.com/c2\",\"text\":\"C2\"},{\"domain\":\"http://www.cash.com/c3\",\"text\":\"C3\"}]",
			Author:          "2 author@cash.ch",
			Comments:        "2 http://www.cash.ch/comments",
			EnclosureUrl:    "2 https://www.cash.ch/fp/1200/675/1910/1089/sites/default/files/images/library/six_boerse_zuerich_aktien.jpg",
			EnclosureLength: 21816035,
			EnclosureType:   "2 image/jpeg",
			GuidUrl:         "2 https://www.cash.ch/node/573556",
			GuidIsPermaLink: "",
			SourceUrl:       "2 http://www.cash.ch/links2.xml",
			SourceText:      "2 Cash Rss",
		},
		{
			Title:           "3 Börsen-Ticker :  SMI schlägt Seitwärts-Kurs ein",
			Link:            "3 https://www.cash.ch/news/top-news/borsen-ticker-smi-schlagt-seitwarts-kurs-ein-596719",
			Description:     "3 +++Märkte+++ - Die Schweizer Börse hat am Donnerstag weder nachgegeben, noch dazu gewonnen.",
			PubDate:         "3 Thu, 27 Apr 2023 06:09:20 +0200",
			Category:        "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"},{\"domain\":\"http://www.cash.com/c2\",\"text\":\"C2\"},{\"domain\":\"http://www.cash.com/c3\",\"text\":\"C3\"}]",
			Author:          "3 author@cash.ch",
			Comments:        "3 http://www.cash.ch/comments",
			EnclosureUrl:    "3 https://www.cash.ch/fp/1200/675/1910/1089/sites/default/files/images/library/six_boerse_zuerich_aktien.jpg",
			EnclosureLength: 31816035,
			EnclosureType:   "3 image/jpeg",
			GuidUrl:         "3 https://www.cash.ch/node/573556",
			GuidIsPermaLink: "false",
			SourceUrl:       "3 http://www.cash.ch/links2.xml",
			SourceText:      "3 Cash Rss",
		},
	}
	for _, field := range newsFields {
		for i, item := range newsItems {
			expected := getNewsField(item, field)
			is := getNewsField(expectedNews[i], field)
			if is != expected {
				t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
			}
		}
	}
}

func TestNewsParserOriginal(t *testing.T) {
	_, newsItems := getFeed("mock_org.xml")
	expectedNews := []*models.News{
		{
			Title:           "Börsen-Ticker :  SMI schlägt Seitwärts-Kurs ein",
			Link:            "https://www.cash.ch/news/top-news/borsen-ticker-smi-schlagt-seitwarts-kurs-ein-596719",
			Description:     "+++Märkte+++ - Die Schweizer Börse hat am Donnerstag weder nachgegeben, noch dazu gewonnen.",
			PubDate:         "Thu, 27 Apr 2023 06:09:20 +0200",
			Author:          "",
			Comments:        "",
			GuidUrl:         "https://www.cash.ch/node/573556",
			GuidIsPermaLink: "",
			EnclosureUrl:    "https://www.cash.ch/fp/1200/675/1910/1089/sites/default/files/images/library/six_boerse_zuerich_aktien.jpg",
			EnclosureLength: 1816035,
			EnclosureType:   "image/jpeg",
			Category:        "[{\"domain\":\"\",\"text\":\"News\"}]",
			SourceUrl:       "",
			SourceText:      "",
		},
		{
			Title:           "Aktien Schweiz Schluss: Kaum verändert - Warten auf Zentralbanken",
			Link:            "https://www.cash.ch/borsen-ticker/borsen-kommentar/aktien-schweiz-schluss-kaum-verandert-warten-auf-zentralbanken-597164",
			Description:     "- Der Schweizer Aktienmarkt hat am Donnerstag nahezu unverändert geschlossen. Dabei bewegte sich der Leitindex SMI bis kurz vor Schluss in einer engen Spanne leicht über dem Vortagesschluss. Doch war der Markt keineswegs homogen. Unter der Oberfläche zeigte sich vielmehr ein uneinheitliches Bild mit Verlierern und Gewinnern. Vor den mit Spannung erwarteten Zinsentscheidungen der US-Notenbank Fed und der Europäischen Zentralbank (EZB) in der kommenden Woche und angesichts unterschiedlicher Firmenergebnisse hätten sich die Anleger zurückgehalten, hiess es am Markt.",
			PubDate:         "Thu, 27 Apr 2023 18:15:10 +0200",
			Author:          "",
			Comments:        "",
			GuidUrl:         "https://www.cash.ch/node/574000",
			GuidIsPermaLink: "",
			EnclosureUrl:    "https://www.cash.ch/fp/1200/675/2592/1727/sites/default/files/images/library/six_borse_schweiz_aktien.jpg",
			EnclosureLength: 949745,
			EnclosureType:   "image/jpeg",
			Category:        "[{\"domain\":\"\",\"text\":\"News\"}]",
			SourceUrl:       "",
			SourceText:      "",
		},
		{
			Title:           "«Eine Kurskorrektur wäre ein spannender Moment, um Aktienpositionen aufzubauen»",
			Link:            "https://www.cash.ch/news/top-news/eine-kurskorrektur-ware-ein-spannender-moment-um-aktienpositionen-aufzubauen-596555",
			Description:     "BLKB-Anlagechefin Fabienne Hockenjos-Erni - Für Fabienne Hockenjos-Erni, Anlagechefin der BLKB, bleibt das Marktumfeld volatil. Im Interview sagt sie, welche Aktien für das restliche Jahr in der Pole-Position sind, auf welche Dividendentitel sie setzt und wo Gefahren lauern.",
			PubDate:         "Thu, 27 Apr 2023 18:00:01 +0200",
			Author:          "",
			Comments:        "",
			GuidUrl:         "https://www.cash.ch/node/573392",
			GuidIsPermaLink: "",
			EnclosureUrl:    "https://www.cash.ch/fp/1200/675/1250/793/sites/default/files/media/field_image/2023-04/fabienne_hockenjos-erni_blkb.jpg",
			EnclosureLength: 2056836,
			EnclosureType:   "image/jpeg",
			Category:        "[{\"domain\":\"\",\"text\":\"News\"}]",
			SourceUrl:       "",
			SourceText:      "",
		},
	}
	for _, field := range newsFields {
		for i, item := range newsItems {
			expected := getNewsField(item, field)
			is := getNewsField(expectedNews[i], field)
			if is != expected {
				t.Errorf("Expected field %s to be %s but got %s", field, expected, is)
			}
		}
	}
}

func TestNewsParserNoNews(t *testing.T) {
	_, newsItems := getFeed("mock_no_news.xml")
	if len(newsItems) != 0 {
		t.Error("Expected news to be empty")
	}
}
