package rss

import (
	"encoding/xml"
	"net/http"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title          string    `xml:"title"`
	Link           string    `xml:"link"`
	Description    string    `xml:"description"`
	Language       string    `xml:"language"`
	PublishDate    Date      `xml:"pubDate"`       // RFC 822
	LastBuildDate  string    `xml:"lastBuildDate"` // RFC 822
	Category       []string  `xml:"category"`
	Copyright      string    `xml:"copyright"`
	ManagingEditor string    `xml:"managingEditor"`
	WebMaster      string    `xml:"webMaster"`
	Generator      string    `xml:"generator"`
	Docs           string    `xml:"docs"`
	TTL            int       `xml:"ttl"`
	Image          Image     `xml:"image"`
	Rating         string    `xml:"rating"`
	SkipDays       SkipDays  `xml:"skipDays"`
	SkipHours      SkipHours `xml:"skipHours"`
	Items          []Item    `xml:"item"`
}

type Date string

type Image struct {
	// Is the URL of a GIF, JPEG or PNG image that represents the channel. 
	Url string `xml:"url"`

	// Describes the image, it's used in the ALT attribute of the HTML <img> tag
	// when the channel is rendered in HTML. 	
	Title string `xml:"title"`

	// The URL of the site, when the channel is rendered, the image is a link to
	// the site. (Note, in practice the image <title> and <link> should have the
	// same value as the channel's <title> and <link>. 
	Link string `xml:"link"`

	// Maximum value for width is 144, default value is 88.
	Width uint `xml:"width"`

	// Maximum value for height is 400, default value is 31.
	Height uint `xml:"height"`
}

type SkipDays struct {
	// An XML element that contains up to seven <day> sub-elements whose value 
	// is Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday. 
	// Aggregators may not read the channel during days listed in the skipDays 
	// element. 
	Days []string
}

type SkipHours struct {
	// An XML element that contains up to 24 <hour> sub-elements whose 
	// value is a number between 0 and 23, representing a time in GMT, 
	// when aggregators, if they support the feature, may not read the 
	// channel on hours listed in the skipHours element.
	// The hour beginning at midnight is hour zero.
	Hours []int
}

type Item struct {
	Guid        string    `xml:"guid"`        // A string that uniquely identifies the item.
	Title       string    `xml:"title"`       // The title of the item.
	Link        string    `xml:"link"`        // The URL of the item.
	Description string    `xml:"description"` // The item synopsis.
	Author      string    `xml:"author"`      // Email address of the author of the item.
	Category    []string  `xml:"category"`    // Includes the item in one or more categories.
	Comments    string    `xml:"comments"`    // URL of a page for comments relating to the item.
	Enclosure   Enclosure `xml:"enclosure"`   // Describes a media object that is attached to the item.
	Pubdate     Date      `xml:"pubDate"`     // Indicates when the item was published.
	Source      string    `xml:"source"`      // The RSS channel that the item came from.
}

type Enclosure struct {
	Url    string `xml:"url"`
	Length uint   `xml:"length"`
	Type   string `xml:"type"` // Standard MIME-type
}

func ReadRss(url string) (*Channel, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	xmlDecoder := xml.NewDecoder(r.Body)
	// TODO: xmlDecoder.CharsetReader = ?

	var rss Rss
	err = xmlDecoder.Decode(&rss)
	if err != nil {
		return nil, err
	}

	return &rss.Channel, nil
}
