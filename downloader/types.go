package downloader

import (
	"fmt"

	"github.com/fatih/color"
)

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
	name    string
}

//Datum download infomation
type Datum struct {
	Title string `json:"title"`
	Type  string `json:"type"`

	Streams map[string]Stream `json:"streams"`

	URL string `json:"url"`
}

//Data 课程信息
type Data struct {
	Title string  `json:"title"`
	Data  []Datum `json:"articles"`
}

//EmptyList empty data list
var EmptyList = make([]Datum, 0)

func (data *Data) printInfo(stream string, isDownloading bool) {
	cyan := color.New(color.FgCyan)
	fmt.Println()
	cyan.Printf(" Title:     ")
	fmt.Println(data.Title)

	if isDownloading {
		for _, article := range data.Data {
			cyan.Printf("     Title:     ")
			fmt.Println(article.Title)
			cyan.Printf("     Streams:   ")
			article.Streams[stream].printStream()
		}
	} else {
		for _, article := range data.Data {
			cyan.Printf("     Title:     ")
			fmt.Println(article.Title)
			cyan.Printf("     Streams:   ")
			fmt.Println("     # All available quality")
			for _, s := range article.Streams {
				s.printStream()
			}
		}
	}
}

func (stream *Stream) calculateTotalSize() {
	size := 0
	for _, url := range stream.URLs {
		size += url.Size
	}
	stream.Size = size
}

func (stream Stream) printStream() {
	blue := color.New(color.FgBlue)
	cyan := color.New(color.FgCyan)

	blue.Println(fmt.Sprintf("          [%s]  -------------------", stream.Quality))
	if stream.Quality != "" {
		cyan.Printf("         Quality:         ")
		fmt.Println(stream.Quality)
	}
	cyan.Printf("         Size:            ")
	if stream.Size == 0 {
		stream.calculateTotalSize()
	}
	fmt.Printf("%.2f MiB (%d Bytes)\n", float64(stream.Size)/1024/1024, stream.Size)
}
