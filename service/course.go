package service

import "fmt"

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
	UpdateFrequency   string `json:"update_frequency"`
}

//Columns 获取专栏
func (s *Service) Columns() ([]*Course, error) {
	return s.getCourses(1)
}

//Videos 获取专栏
func (s *Service) Videos() ([]*Course, error) {
	return s.getCourses(3)
}

//获取课程信息
func (s *Service) getCourses(courseType int) ([]*Course, error) {
	ids, err := s.courses(courseType)

	if err != nil {
		return nil, err
	}

	return s.showCourses(ids)
}

func (s *Service) courses(courseType int) ([]int, error) {
	body, err := s.requestCourses(courseType)

	if err != nil {
		return nil, err
	}

	defer body.Close()

	courses := new(CourseList)
	if err := handleJSONParse(body, &courses); err != nil {
		return nil, err
	}

	var ids []int
	for _, item := range courses.List {
		ids = append(ids, item.ID)
	}

	return ids, nil
}

func (s *Service) showCourses(ids []int) ([]*Course, error) {
	body, err := s.requestCourseDetail(ids)
	if err != nil {
		return nil, err
	}

	defer body.Close()

	var courses []*Course
	if err := handleJSONParse(body, &courses); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return courses, nil
}
