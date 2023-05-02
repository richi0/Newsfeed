package handlers_test

import "newsfeed/models"

var BASE_URL = "http://localhost:4000/"

var FEED_1 = &models.Feed{
	Url:              "https://www.cash.ch/1",
	Title:            "Cash - \"Aktuell\" | News_1",
	Link:             "https://www.cash.ch/feeds/latest/news_1",
	Description:      "Cash - \"Aktuell\" | News_1",
	Language:         "de_1",
	Copyright:        "Cash - Ringier Axel Springer Schweiz AG_1",
	PubDate:          "Thu, 27 Apr 2023 18:29:56 +0200_1",
	LastBuildDate:    "Thu, 27 Apr 2023 18:29:56 +0200_1",
	Category:         "[{\"domain\":\"http://www.cash.com/c1\",\"text\":\"C1\"}]",
	Docs:             "https://www.rssboard.org/rss-specification_1",
	ImageUrl:         "https://www.cash.ch/sites/default/files/media/field_image/2022-08/cash_favicon_96x96.png_1",
	ImageTitle:       "Cash - \"Aktuell\" | News_1",
	ImageLink:        "https://www.cash.ch/feeds/latest/news_1",
	ImageDescription: "Cash Image_1",
	Rating:           "Rating_1",
	SkipHours:        "[3]",
	SkipDays:         "null",
}

var FEED_2 = &models.Feed{
	Url:              "https://www.cash.ch/2",
	Title:            "Cash - \"Aktuell\" | News_2",
	Link:             "https://www.cash.ch/feeds/latest/news_2",
	Description:      "Cash - \"Aktuell\" | News_2",
	Language:         "de_2",
	Copyright:        "Cash - Ringier Axel Springer Schweiz AG_2",
	PubDate:          "Thu, 27 Apr 2023 18:29:56 +0200_2",
	LastBuildDate:    "Thu, 27 Apr 2023 18:29:56 +0200_2",
	Category:         "[{\"domain\":\"http://www.cash.com/c2\",\"text\":\"C2\"}]",
	Docs:             "https://www.rssboard.org/rss-specification_2",
	ImageUrl:         "https://www.cash.ch/sites/default/files/media/field_image/2022-08/cash_favicon_96x96.png_2",
	ImageTitle:       "Cash - \"Aktuell\" | News_2",
	ImageLink:        "https://www.cash.ch/feeds/latest/news_2",
	ImageDescription: "Cash Image_2",
	Rating:           "Rating_2",
	SkipHours:        "[3]",
	SkipDays:         "null",
}

var FEED_3 = &models.Feed{
	Url:              "https://www.cash.ch/3",
	Title:            "Cash - \"Aktuell\" | News_3",
	Link:             "https://www.cash.ch/feeds/latest/news_3",
	Description:      "Cash - \"Aktuell\" | News_3",
	Language:         "de_3",
	Copyright:        "Cash - Ringier Axel Springer Schweiz AG_3",
	PubDate:          "Thu, 27 Apr 2023 18:29:56 +0200_3",
	LastBuildDate:    "Thu, 27 Apr 2023 18:29:56 +0200_3",
	Category:         "[{\"domain\":\"http://www.cash.com/c3\",\"text\":\"C3\"}]",
	Docs:             "https://www.rssboard.org/rss-specification_3",
	ImageUrl:         "https://www.cash.ch/sites/default/files/media/field_image/2022-08/cash_favicon_96x96.png_3",
	ImageTitle:       "Cash - \"Aktuell\" | News_3",
	ImageLink:        "https://www.cash.ch/feeds/latest/news_3",
	ImageDescription: "Cash Image_3",
	Rating:           "Rating_3",
	SkipHours:        "[3]",
	SkipDays:         "null",
}

var NEWS_1 = &models.News{
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
}

var NEWS_2 = &models.News{
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
}

var NEWS_3 = &models.News{
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
}
