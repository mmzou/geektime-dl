package service

//CourseList 课程列表基础信息
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

//Course 课程信息
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

//Article 课程文章信息
type Article struct {
	ID             int    `json:"id"`
	ArticleTitle   string `json:"article_title"`
	ArticleSummary string `json:"article_summary"`
	ArticleCover   string `json:"article_cover"`
	ArticleTime    int    `json:"article_ctime"`
	ChapterID      int    `json:"chapter_id string"`
	IncludeAudio   bool   `json:"include_audio"`
	//Is can preview
	ColumnHadSub        bool `json:"column_had_sub"`
	ArticleCouldPreview bool `json:"article_could_preview"`
	//Audio info
	AudioDownloadURL string `json:"audio_download_url"`
	AudioSize        int    `json:"audio_size"`
}

type articleResult struct {
	Articles []*Article `json:"list"`
	Page     struct {
		Count int  `json:"count"`
		More  bool `json:"more"`
	} `json:"page"`
}

//IsColumn 是否专栏
func (course *Course) IsColumn() bool {
	return course.ColumnType == 1
}

//IsVideo 是否视频
func (course *Course) IsVideo() bool {
	return course.ColumnType == 3
}

//IsCanPreview 是否能看
func (article *Article) IsCanPreview() bool {
	return article.ColumnHadSub || article.ArticleCouldPreview
}
