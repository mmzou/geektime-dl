package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mmzou/geektime-dl/config"
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

//FilePath get valid file path
func FilePath(name, ext string, escape bool) (string, error) {
	var downloadPath string
	if config.Instance.DownloadPath != "" {
		if _, err := os.Stat(config.Instance.DownloadPath); err != nil {
			return downloadPath, err
		}
	}

	fileName := fmt.Sprintf("%s.%s", name, ext)
	if escape {
		fileName = FileName(fileName)
	}

	return filepath.Join(config.Instance.DownloadPath, fileName), nil
}
