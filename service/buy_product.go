package service

// ProductAll all protuct
type ProductAll struct {
	Columns *Product
	Videos  *Product
}

//Product all product
type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Page  struct {
		More  bool `json:"more"`
		Count int  `json:"count"`
	} `json:"page"`
	List []struct {
		Title  string `json:"title"`
		Conver string `json:"cover"`
		Type   string `json:"type"`
		Extra  struct {
			LastAid        int    `json:"last_aid"`
			ColumnID       int    `json:"column_id"`
			ColumnTitle    string `json:"column_title"`
			ColumnSubtitle string `json:"column_subtitle"`
			AuthorName     string `json:"author_name"`
			AuthorIntro    string `json:"author_intro"`
			ColumnCover    string `json:"column_cover"`
			ColumnType     int    `json:"column_type"`
			ArticleCount   int    `json:"article_count"`
			IsIncludeAudio bool   `json:"is_include_audio"`
		} `json:"extra"`
	} `json:"list"`
}

//BuyProductAll 获取所有购买的课程信息
func (s *Service) BuyProductAll() (*ProductAll, error) {
	body, err := s.requestBuyAll()

	if err != nil {
		return nil, err
	}

	defer body.Close()

	var products []*Product

	if err := handleJSONParse(body, &products); err != nil {
		return nil, err
	}
	productAll := &ProductAll{
		Columns: products[0],
		Videos:  products[1],
	}

	return productAll, nil
}
