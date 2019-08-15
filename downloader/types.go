package downloader

//URL for url infomation
type URL struct {
	URL  string `json:"url"`
	Size int    `json:"size"`
	Ext  string `json:"ext"`
}

//Stream data
type Stream struct {
	URLs    []URL  `json:"urls"`
	Size    int    `json:"size"`
	Quality string `json:"quality"`
	Name    string `json:"name"`
}

//Data download infomation
type Data struct {
	Title string `json:"title"`
	Type  string `json:"type"`

	Streams map[string]Stream `json:"streams"`

	URL string `json:"url"`
}

//EmptyList empty data list
var EmptyList = make([]Data, 0)

func (stream *Stream) calculateTotalSize() {
	size := 0
	for _, url := range stream.URLs {
		size += url.Size
	}
	stream.Size = size
}
