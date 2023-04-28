package parsers

type Rss struct {
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title          string     `xml:"title"`
	Link           *Link      `xml:"link"`
	Description    string     `xml:"description"`
	Language       string     `xml:"language"`
	Copyright      string     `xml:"copyright"`
	ManagingEditor string     `xml:"managingEditor"`
	WebMaster      string     `xml:"webMaster"`
	PubDate        string     `xml:"pubDate"`
	LastBuildDate  string     `xml:"lastBuildDate"`
	Category       *Category  `xml:"category"`
	Generator      string     `xml:"generator"`
	Docs           string     `xml:"docs"`
	Cloud          *Cloud     `xml:"cloud"`
	Ttl            string     `xml:"ttl"`
	Image          *Image     `xml:"image"`
	Rating         string     `xml:"rating"`
	TextInput      *TextInput `xml:"textInput"`
	SkipHours      string     `xml:"skipHours"`
	SkipDays       string     `xml:"skipDays"`
	Items          []*Item    `xml:"item"`
}

type Content struct {
	Expression string `xml:"expression,attr"`
	Height     int    `xml:"height,attr"`
	Medium     string `xml:"medium,attr"`
	Type       string `xml:"type,attr"`
	Url        string `xml:"url,attr"`
	Width      int    `xml:"width,attr"`
}

type Cloud struct {
	Domain            string `xml:"domain,attr"`
	Port              int    `xml:"port,attr"`
	RegisterProcedure string `xml:"registerProcedure,attr"`
	Protocol          string `xml:"protocol,attr"`
}

type TextInput struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        *Link  `xml:"link"`
}

type Group struct {
	Content []Content `xml:"content"`
}

type Guid struct {
	IsPermaLink bool   `xml:"isPermaLink,attr"`
	CharData    string `xml:",chardata"`
}

type Item struct {
	Category    *Category `xml:"category"`
	Content     *Content  `xml:"content"`
	Description string    `xml:"description"`
	Enclosure   Enclosure `xml:"enclosure"`
	Encoded     string    `xml:"encoded"`
	Group       *Group    `xml:"group"`
	Guid        *Guid     `xml:"guid"`
	Link        *Link     `xml:"link"`
	PubDate     string    `xml:"pubDate"`
	Title       string    `xml:"title"`
}

type Category struct {
	Domain   string `xml:"domain,attr"`
	CharData string `xml:",chardata"`
}

type Enclosure struct {
	Length int    `xml:"length,attr"`
	Type   string `xml:"type,attr"`
	Url    string `xml:"url,attr"`
}

type Image struct {
	Description string `xml:"description"`
	Height      int    `xml:"height"`
	Link        Link   `xml:"link"`
	Title       string `xml:"title"`
	Url         string `xml:"url"`
	Width       int    `xml:"width"`
}

type Link struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	CharData string `xml:",chardata"`
}
