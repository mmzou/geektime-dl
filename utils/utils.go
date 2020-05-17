package utils

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mmzou/geektime-dl/requester"
)

// MAXLENGTH Maximum length of file name
const MAXLENGTH = 80

//FileName filter invalid string
func FileName(name string, ext string) string {
	rep := strings.NewReplacer("\n", " ", "/", " ", "|", "-", ": ", "：", ":", "：", "'", "’", "\t", " ")
	name = rep.Replace(name)

	if runtime.GOOS == "windows" {
		rep := strings.NewReplacer("\"", " ", "?", " ", "*", " ", "\\", " ", "<", " ", ">", " ", ":", " ", "：", " ")
		name = rep.Replace(name)
	}

	name = strings.TrimSpace(name)

	limitedName := LimitLength(name, MAXLENGTH)
	if ext != "" {
		return fmt.Sprintf("%s.%s", limitedName, ext)
	}
	return limitedName
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

// FilePath gen valid file path
func FilePath(name, ext string, escape bool) (string, error) {
	var outputPath string

	var fileName string
	if escape {
		fileName = FileName(name, ext)
	} else {
		fileName = fmt.Sprintf("%s.%s", name, ext)
	}
	outputPath = filepath.Join(fileName)
	return outputPath, nil
}

//Mkdir mkdir path
func Mkdir(elem ...string) (string, error) {
	path := filepath.Join(elem...)

	err := os.MkdirAll(path, os.ModePerm)

	return path, err
}

// FileSize return the file size of the specified path file
func FileSize(filePath string) (int, bool, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return int(file.Size()), true, nil
}

// M3u8URLs get all ts urls from m3u8 url
func M3u8URLs(uri string) ([]string, error) {
	if len(uri) == 0 {
		return nil, errors.New("M3u8地址为空")
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
