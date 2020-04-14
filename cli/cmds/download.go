package cmds

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/downloader"
	"github.com/mmzou/geektime-dl/service"
	"github.com/mmzou/geektime-dl/utils"
	"github.com/urfave/cli"
)

//NewDownloadCommand login command
func NewDownloadCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "",
			Usage:     "",
			UsageText: "",
			Action:    downloadAction,
			Before:    authorizationFunc,
		},
	}
}

func downloadAction(c *cli.Context) error {
	showInfo := c.Parent().Bool("info")

	args := c.Parent().Args()
	cid, err := strconv.Atoi(args.First())
	if err != nil {
		cli.ShowCommandHelp(c, "download")
		return errors.New("请输入课程ID")
	}

	//课程目录ID
	aid := 0
	if len(args) > 1 {
		aid, err = strconv.Atoi(args.Get(1))
		if err != nil {
			return errors.New("课程目录ID错误")
		}
	}

	course, articles, err := application.CourseWithArticles(cid)
	if err != nil {
		return err
	}

	downloadData := extractDownloadData(course, articles, aid, showInfo)
	// printExtractDownloadData(downloadData)

	if showInfo {
		downloadData.PrintInfo()
		return nil
	}

	errors := make([]error, 0)
	for _, datum := range downloadData.Data {
		if err := downloader.Download(datum); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors[0]
	}

	return nil
}

func extractDownloadData(course *service.Course, articles []*service.Article, aid int, showInfo bool) downloader.Data {
	downloadData := downloader.Data{
		Title: course.ColumnTitle,
	}
	data := downloader.EmptyData

	if course.IsColumn() {
		downloadData.Type = "专栏"
		key := "df"
		for _, article := range articles {
			if aid > 0 && article.ID != aid {
				continue
			}
			urls := []downloader.URL{
				{
					URL:  article.AudioDownloadURL,
					Size: article.AudioSize,
					Ext:  "mp3",
				},
			}

			streams := map[string]downloader.Stream{
				key: {
					URLs:    urls,
					Size:    article.AudioSize,
					Quality: key,
				},
			}

			data = append(data, downloader.Datum{
				ID:      article.ID,
				Title:   article.ArticleTitle,
				IsCanDL: article.IsCanPreview(),
				Streams: streams,
				Type:    "专栏",
			})
		}
	} else if course.IsVideo() {
		downloadData.Type = "视频"

		videoIds := map[int]string{}

		videoData := make([]*downloader.Datum, 0)

		for _, article := range articles {
			if aid > 0 && article.ID != aid {
				continue
			}

			videoIds[article.ID] = article.VideoID

			videoMediaMaps := &map[string]downloader.VideoMediaMap{}
			utils.UnmarshalJSON(article.VideoMediaMap, videoMediaMaps)

			urls := []downloader.URL{}

			streams := map[string]downloader.Stream{}
			for key, videoMediaMap := range *videoMediaMaps {
				streams[key] = downloader.Stream{
					URLs:    urls,
					Size:    videoMediaMap.Size,
					Quality: key,
				}
			}

			datum := &downloader.Datum{
				ID:      article.ID,
				Title:   article.ArticleTitle,
				IsCanDL: article.IsCanPreview(),
				Streams: streams,
				Type:    "视频",
			}

			videoData = append(videoData, datum)
		}

		if !showInfo {
			wgp := utils.NewWaitGroupPool(10)
			for _, datum := range videoData {
				wgp.Add()
				go func(datum *downloader.Datum, streams map[int]string) {
					defer func() {
						wgp.Done()
					}()
					if datum.IsCanDL {
						playInfo, _ := application.GetVideoPlayInfo(datum.ID, streams[datum.ID])
						for _, info := range playInfo.PlayInfoList.PlayInfo {
							if urls, err := utils.M3u8URLs(info.URL); err == nil {
								key := strings.ToLower(info.Definition)
								stream := datum.Streams[key]
								for _, url := range urls {
									stream.URLs = append(stream.URLs, downloader.URL{
										URL: url,
										Ext: "ts",
									})
								}
								datum.Streams[key] = stream
							}
						}
					}
				}(datum, videoIds)
			}
			wgp.Wait()
		}
		for _, d := range videoData {
			data = append(data, *d)
		}
	}

	downloadData.Data = data

	return downloadData
}

func printExtractDownloadData(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", jsonData)
	}
}
