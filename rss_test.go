package rss

import (
	"testing"
)

func TestReadRss(t *testing.T) {
	var channel *Channel

	channel, err := ReadRss("http://www.joelonsoftware.com/rss.xml")
	if err != nil {
		t.Error(err)
	}

	if channel.Title != "Joel on Software" {
		t.Error("Title:" + channel.Title + " != Joel on Software")
	}
}
