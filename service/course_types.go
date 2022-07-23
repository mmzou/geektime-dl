package service

import "encoding/json"

// CourseList 课程列表基础信息
type CourseList struct {
	List []struct {
		ID                int  `json:"id"`
		ColumnCtime       int  `json:"column_ctime"`
		ColumnGroupbuy    int  `json:"column_groupbuy"`
		ColumnPrice       int  `json:"column_price"`
		ColumnPriceMarket int  `json:"column_price_market"`
		ColumnSku         int  `json:"column_sku"`
		ColumnType        int  `json:"column_type"`
		HadSub            bool `json:"had_sub"`
		IsChannel         int  `json:"is_channel"`
		IsExperience      bool `json:"is_experience"`
		LastAid           int  `json:"last_aid"`
		LastChapterID     int  `json:"last_chapter_id"`
		PriceType         int  `json:"price_type"`
		SubCount          int  `json:"sub_count"`
	} `json:"list"`
}

// Course 课程信息
type Course struct {
	ID                int    `json:"id"`
	Authorintro       string `json:"author_intro"`
	AuthorName        string `json:"author_name"`
	ChannelBackAmount int    `json:"channel_back_amount"`
	ColumnBgcolor     string `json:"column_bgcolor"`
	ColumnCover       string `json:"column_cover"`
	ColumnCtime       int    `json:"column_ctime"`
	ColumnPrice       int    `json:"column_price"`
	ColumnPriceMarket int    `json:"column_price_market"`
	ColumnPriceSale   int    `json:"column_price_sale"`
	ColumnSku         int    `json:"column_sku"`
	ColumnSubtitle    string `json:"column_subtitle"`
	ColumnTitle       string `json:"column_title"`
	ColumnType        int    `json:"column_type"`
	ColumnTnit        string `json:"column_unit"`
	HadSub            bool   `json:"had_sub"`
	IsChannel         bool   `json:"is_channel"`
	IsExperience      bool   `json:"is_experience"`
	IsOnboard         bool   `json:"is_onboard"`
	PriceType         int    `json:"price_type"`
	SubCount          int    `json:"sub_count"`
	ShowChapter       bool   `json:"show_chapter"`
	UpdateFrequency   string `json:"update_frequency"`
}

// Article 课程文章信息
type Article struct {
	ID             int    `json:"id"`
	ArticleContent string `json:"article_content"`
	ArticleTitle   string `json:"article_title"`
	ArticleSummary string `json:"article_summary"`
	ArticleCover   string `json:"article_cover"`
	ArticleTime    int    `json:"article_ctime"`
	ChapterID      int    `json:"chapter_id string"`
	IncludeAudio   bool   `json:"include_audio"`
	// Is can preview
	ColumnHadSub        bool `json:"column_had_sub"`
	ArticleCouldPreview bool `json:"article_could_preview"`
	// Audio info
	AudioDownloadURL string `json:"audio_download_url"`
	AudioSize        int    `json:"audio_size"`
	// Viode info
	VideoMediaMap json.RawMessage `json:"video_media_map"`
	VideoID       string          `json:"video_id"`
	VideoCover    string          `json:"video_cover"`
}

// VideoPlayAuth 视频的播放授权信息
type VideoPlayAuth struct {
	PlayAuth string `json:"play_auth"`
}

// VideoPlayInfo 视频播放信息
type VideoPlayInfo struct {
	VideoBase struct {
		VideoID  string `json:"VideoId"`
		Title    string `json:"Title"`
		CoverURL string `josn:"CoverURL"`
	} `json:"VideoBase"`
	PlayInfoList struct {
		PlayInfo []struct {
			URL        string `json:"PlayURL"`
			Size       int64  `json:"Size"`
			Definition string `json:"Definition"`
		} `json:"PlayInfo"`
	} `json:"PlayInfoList"`
}

type articleResult struct {
	Articles []*Article `json:"list"`
	Page     struct {
		Count int  `json:"count"`
		More  bool `json:"more"`
	} `json:"page"`
}

// CommentList 文稿评论列表
type CommentList struct {
	List []Comment `json:"list"`
	Page struct {
		Count int  `json:"count"`
		More  bool `json:"more"`
	} `json:"page"`
}
type Comment struct {
	HadLiked           bool                `json:"had_liked"`
	Id                 int                 `json:"id"`
	UserName           string              `json:"user_name"`
	CanDelete          bool                `json:"can_delete"`
	ProductType        string              `json:"product_type"`
	Uid                int                 `json:"uid"`
	Ucode              string              `json:"ucode"`
	UserHeader         string              `json:"user_header"`
	CommentIsTop       bool                `json:"comment_is_top"`
	CommentCtime       int                 `json:"comment_ctime"`
	IsPvip             bool                `json:"is_pvip"`
	Replies            []CommentReplay     `json:"replies"`
	DiscussionCount    int                 `json:"discussion_count"`
	RaceMedal          int                 `json:"race_medal"`
	Score              string              `json:"score"`
	ProductId          int                 `json:"product_id"`
	CommentContent     string              `json:"comment_content"`
	LikeCount          int                 `json:"like_count"`
	CommentDiscussions *CommentDiscussions `json:"comment_discussions"`
}

type CommentReplay struct {
	Ctime        int    `json:"ctime"`
	UserNameReal string `json:"user_name_real"`
	Uid          string `json:"uid"`
	Id           string `json:"id"`
	Content      string `json:"content"`
	Utype        int    `json:"utype"`
	CommentId    int    `json:"comment_id"`
	UserName     string `json:"user_name"`
}

type Author struct {
	Id        int    `json:"id"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Note      string `json:"note"`
	Ucode     string `json:"ucode"`
	RaceMedal int    `json:"race_medal"`
	UserType  int    `json:"user_type"` // 1-读者, 2-作者
	IsPvip    bool   `json:"is_pvip"`
}

type Discussion struct {
	Id                int    `json:"id"`
	DiscussionContent string `json:"discussion_content"`
	LikesNumber       int    `json:"likes_number"`
	IsDelete          bool   `json:"is_delete"`
	IsHidden          bool   `json:"is_hidden"`
	Ctime             int    `json:"ctime"`
	IsLiked           bool   `json:"is_liked"`
	CanDelete         bool   `json:"can_delete"`
	IsComplain        bool   `json:"is_complain"`
	IsTop             bool   `json:"is_top"`
	ParentId          int    `json:"parent_id"`
}

type CommentDiscussions struct {
	List []struct {
		Author                Author     `json:"author"`
		ReplyAuthor           Author     `json:"reply_author"`
		Discussion            Discussion `json:"discussion"`
		Score                 int        `json:"score"`
		Extra                 string     `json:"extra"`
		ChildDiscussionNumber int        `json:"child_discussion_number"`
		ChildDiscussions      []struct {
			Author      Author     `json:"author"`
			ReplyAuthor Author     `json:"reply_author"`
			Discussion  Discussion `json:"discussion"`
			Score       int        `json:"score"`
			Extra       string     `json:"extra"`
		} `json:"child_discussions"`
	} `json:"list"`
	Page struct {
		Total   int  `json:"total"`
		More    bool `json:"more"`
		Current int  `json:"current"`
	} `json:"page"`
}

// IsColumn 是否专栏
func (course *Course) IsColumn() bool {
	return course.ColumnType == 1
}

// IsVideo 是否视频
func (course *Course) IsVideo() bool {
	return course.ColumnType == 3
}

// IsCanPreview 是否能看
func (article *Article) IsCanPreview() bool {
	return article.ColumnHadSub || article.ArticleCouldPreview
}
