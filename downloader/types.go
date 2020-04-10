package downloader

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/olekukonko/tablewriter"
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
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Type    string `json:"type"`
	IsCanDL bool   `json:"is_can_dl"`

	Streams map[string]Stream `json:"streams"`

	URL string `json:"url"`
}

//Data 课程信息
type Data struct {
	Title string  `json:"title"`
	Data  []Datum `json:"articles"`
}

//EmptyData empty data list
var EmptyData = make([]Datum, 0)

func (data *Data) printInfo(s string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "类型", "名称", "大小", "是否能下载"})
	table.SetAutoWrapText(false)
	i := 0
	for _, p := range data.Data {
		stream := p.Streams[s]
		if stream.Size == 0 {
			stream.calculateTotalSize()
		}
		//计算大小
		size := fmt.Sprintf("%.2fM", float64(stream.Size)/1024/1024)

		reg, _ := regexp.Compile(" \\| ")
		title := reg.ReplaceAllString(p.Title, " ")

		isCanDL := ""
		if p.IsCanDL {
			isCanDL = "是"
		}

		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.ID), p.Type, title, size, isCanDL})
		i++
	}
	table.Render()
}

func (stream *Stream) calculateTotalSize() {
	size := 0
	for _, url := range stream.URLs {
		size += url.Size
	}
	stream.Size = size
}
