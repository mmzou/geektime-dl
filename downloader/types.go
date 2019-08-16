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

//Article download infomation
type Article struct {
	Title string `json:"title"`
	Type  string `json:"type"`

	Streams map[string]Stream `json:"streams"`

	URL string `json:"url"`
}

//Course 课程信息
type Course struct {
	Title    string    `json:"title"`
	Articles []Article `json:"articles"`
}

//EmptyList empty data list
var EmptyList = make([]Article, 0)

func (stream *Stream) calculateTotalSize() {
	size := 0
	for _, url := range stream.URLs {
		size += url.Size
	}
	stream.Size = size
}
