package cmds

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/downloader"
	"github.com/mmzou/geektime-dl/service"
	"github.com/mmzou/geektime-dl/utils"
	"github.com/urfave/cli"
)

//NewDownloadCommand login command
func NewDownloadCommand() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:      "download",
			Usage:     "下载",
			UsageText: appName + " download",
			Action:    downloadAction,
			Before:    authorizationFunc,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "info, i", Usage: "只输出视频信息"},
			},
		},
	}
}

func downloadAction(c *cli.Context) error {
	args := c.Parent().Args()
	if args.First() == "download" {
		args = c.Args()
	}

	cid, err := strconv.Atoi(args.First())
	if err != nil {
		cli.ShowCommandHelp(c, "download")
		return errors.New("请输入课程ID")
	}

	course, articles, err := application.CourseWithArticles(cid)
	if err != nil {
		return err
	}

	downloadData := extractDownloadData(course, articles)
	// printExtractDownloadData(downloadData)

	downloader.Download(downloadData)

	return nil
}

func extractDownloadData(course *service.Course, articles []*service.Article) downloader.Data {
	downloadData := downloader.Data{
		Title: course.ColumnTitle,
	}
	data := downloader.EmptyData

	if course.IsColumn() {
		key := "df"
		for _, article := range articles {
			if !article.IncludeAudio {
				//	continue
			}
			urls := []downloader.URL{
				{
					URL:  article.AudioDownloadURL,
					Size: article.AudioSize,
					Ext:  "mp3",
				},
			}

			streams := map[string]downloader.Stream{
				key: downloader.Stream{
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
		for _, article := range articles {
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

			data = append(data, downloader.Datum{
				ID:      article.ID,
				Title:   article.ArticleTitle,
				IsCanDL: article.IsCanPreview(),
				Streams: streams,
				Type:    "视频",
			})
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
