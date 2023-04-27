package parsers

type Channel struct {
	Copyright       *string   `xml:"copyright"`
	Description     string    `xml:"description"`
	Docs            *string   `xml:"docs"`
	Generator       *string   `xml:"generator"`
	Image           *Image    `xml:"image"`
	Language        *string   `xml:"language"`
	LastBuildDate   *string   `xml:"lastBuildDate"`
	Link            []Link    `xml:"link"`
	PubDate         *string   `xml:"pubDate"`
	Title           string    `xml:"title"`
	Ttl             *string   `xml:"ttl"`
	UpdateFrequency *bool     `xml:"updateFrequency"`
	UpdatePeriod    *string   `xml:"updatePeriod"`
	ManagingEditor  *string   `xml:"managingEditor"`
	WebMaster       *string   `xml:"webMaster"`
	Category        *Category `xml:"category"`
}

type Category struct {
	Domain   *string `xml:"domain,attr"`
	CharData string  `xml:",chardata"`
}

type Enclosure struct {
	Length int    `xml:"length,attr"`
	Type   string `xml:"type,attr"`
	Url    string `xml:"url,attr"`
}

type Image struct {
	Description *string `xml:"description"`
	Height      *int    `xml:"height"`
	Link        Link    `xml:"link"`
	Title       string  `xml:"title"`
	Url         string  `xml:"url"`
	Width       *int    `xml:"width"`
}

type Link struct {
	Href     *string `xml:"href,attr"`
	Rel      *string `xml:"rel,attr"`
	Type     *string `xml:"type,attr"`
	CharData string  `xml:",chardata"`
}
