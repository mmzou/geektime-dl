package utils

import (
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"strings"

	"github.com/mmzou/geektime-dl/requester"
)

// MAXLENGTH Maximum length of file name
const MAXLENGTH = 80

//FileName filter invalid string
func FileName(name string) string {
	rep := strings.NewReplacer("\n", " ", "/", " ", "|", "-", ": ", "：", ":", "：", "'", "’")
	name = rep.Replace(name)

	if runtime.GOOS == "windows" {
		rep := strings.NewReplacer("\"", " ", "?", " ", "*", " ", "\\", " ", "<", " ", ">", " ")
		name = rep.Replace(name)
	}

	return LimitLength(name, MAXLENGTH)
}

//LimitLength cut string
func LimitLength(s string, length int) string {
	ellipses := "..."

	str := []rune(s)
	if len(str) > length {
		s = string(str[:length-len(ellipses)]) + ellipses
	}

	return s
}

// M3u8URLs get all ts urls from m3u8 url
func M3u8URLs(uri string) ([]string, error) {
	if len(uri) == 0 {
		return nil, errors.New("Url is null")
	}

	html, err := requester.HTTPGet(uri)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(html), "\n")
	var urls []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			if strings.HasPrefix(line, "http") {
				urls = append(urls, line)
			} else {
				base, err := url.Parse(uri)
				if err != nil {
					continue
				}
				u, err := url.Parse(line)
				if err != nil {
					continue
				}
				urls = append(urls, fmt.Sprintf("%s", base.ResolveReference(u)))
			}
		}
	}
	return urls, nil
}

// //FilePath get valid file path
// func FilePath(name, ext string, escape bool) (string, error) {
// 	var downloadPath string
// 	if config.Instance.DownloadPath != "" {
// 		if _, err := os.Stat(config.Instance.DownloadPath); err != nil {
// 			return downloadPath, err
// 		}
// 	}

// 	fileName := fmt.Sprintf("%s.%s", name, ext)
// 	if escape {
// 		fileName = FileName(fileName)
// 	}

// 	return filepath.Join(config.Instance.DownloadPath, fileName), nil
// }
