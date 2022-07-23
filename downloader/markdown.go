package downloader

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/service"
	"github.com/mmzou/geektime-dl/utils"
)

func PrintToMarkdown(v Datum, path string) error {
	name := utils.FileName(v.Title, "md")

	filename := filepath.Join(path, name)

	fmt.Printf("正在生成文件：【\033[37;1m%s\033[0m】 ", name)

	_, exist, err := utils.FileSize(filename)
	if err != nil {
		fmt.Printf("\033[31;1m%s, err=%v\033[0m\n", "失败1", err)
		return err
	}

	if exist {
		fmt.Printf("\033[33;1m%s\033[0m\n", "已存在")
		return nil
	}

	if err != nil {
		fmt.Printf("\033[31;1m%s, err=%v\033[0m\n", "失败2", err)
		return err
	}

	detail, err := application.ArticleDetail(v.ID)

	if err != nil {
		fmt.Printf("\033[31;1m%s, err=%v\033[0m\n", "失败3", err)
		return err
	}
	res := contentsToMarkdown(v.Title, detail.ArticleContent)

	commentList, err := application.ArticleComments(v.ID)

	if err != nil {
		fmt.Printf("\033[31;1m%s, err=%v\033[0m\n", "失败4", err)
		return err
	}
	res += articleCommentsToMarkdown(commentList)

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return err
	}
	_, err = f.WriteString(res)
	if err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return err
	}
	if err = f.Close(); err != nil {
		if err != nil {
			return err
		}
	}
	fmt.Printf("\033[32;1m%s\033[0m\n", "完成")

	return nil
}

func contentsToMarkdown(title, content string) (res string) {
	res += getMdHeader(1) + title + "\r\n\r\n"
	res += content + "\r\n\r\n"
	res += "---\r\n\r\n"
	return
}
func articleCommentsToMarkdown(content *service.CommentList) (res string) {
	res = getMdHeader(2) + "精选留言\r\n\r\n"
	for _, comment := range content.List {
		res += comment.UserName + ": " + comment.CommentContent + "\r\n\r\n"

		// 留言讨论
		if comment.CommentDiscussions != nil && len(comment.CommentDiscussions.List) > 0 {
			for _, reply := range comment.CommentDiscussions.List {
				if reply.ReplyAuthor.Nickname != "" {
					res += "> " + reply.Author.Nickname + " ▶ " + reply.ReplyAuthor.Nickname + ": "
				} else {
					res += "> " + reply.Author.Nickname + ": "
				}

				res += strings.Replace(reply.Discussion.DiscussionContent, "\n", "", -1) + "\r\n\r\n"

				if reply.ChildDiscussionNumber > 0 {
					for _, discussion := range reply.ChildDiscussions {
						if discussion.ReplyAuthor.Nickname != "" {
							res += ">> " + discussion.Author.Nickname + " ▶" + discussion.ReplyAuthor.Nickname + ": "
						} else {
							res += ">> " + discussion.Author.Nickname + ":"
						}

						res += strings.Replace(discussion.Discussion.DiscussionContent, "\n", "", -1) + "\r\n\r\n"
					}
				}
			}
		}
	}
	res += "---\r\n"
	return
}

func getMdHeader(level int) string {
	switch level {
	case 1:
		return "# "
	case 2:
		return "## "
	case 3:
		return "### "
	case 4:
		return "#### "
	case 5:
		return "##### "
	case 6:
		return "###### "
	default:
		return ""
	}
}
