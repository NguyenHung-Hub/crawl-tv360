package models

type Response struct {
	Data []struct {
		ID       interface{} `json:"id"`
		Name     string      `json:"name"`
		Type     string      `json:"type"`
		ItemType string      `json:"itemType,omitempty"`
		Content  []Content   `json:"content"`
	} `json:"data"`
}

type Content struct {
	ContentID     string `json:"contentId" gorm:"column:contentId;"`
	Name          string `json:"name" gorm:"column:name;"`
	Description   string `json:"description" gorm:"column:description;"`
	PublishedTime string `json:"publishedTime" gorm:"column:publishedTime;"`
	CoverImage    string `json:"coverImage" gorm:"column:coverImage;"`
	CoverImageH   string `json:"coverImageH" gorm:"column:coverImageH;"`
	Duration      int64  `json:"duration" gorm:"column:duration"`
	CategoryID    string `json:"-" gorm:"column:categoryId;"`
	Type          string `json:"-" gorm:"column:type;"`
	IsFree        int64  `json:"isFree" gorm:"column:isFree;"`
	IsDrm         int64  `json:"isDrm" gorm:"column:isDrm;"`
}

type Response2 struct {
	Data struct {
		ID      interface{} `json:"id"`
		Name    string      `json:"name"`
		Type    string      `json:"type"`
		Content []Content2  `json:"content"`
	} `json:"data"`
}

type Content2 struct {
	ContentID     string `json:"id" gorm:"column:contentId;"`
	Name          string `json:"name" gorm:"column:name;"`
	Description   string `json:"description" gorm:"column:description;"`
	PublishedTime string `json:"publishedTime" gorm:"column:publishedTime;"`
	CoverImage    string `json:"coverImage" gorm:"column:coverImage;"`
	CoverImageH   string `json:"coverImageH" gorm:"column:coverImageH;"`
	Duration      int64  `json:"duration" gorm:"column:duration"`
	CategoryID    string `json:"-" gorm:"column:categoryId;"`
	Type          string `json:"type" gorm:"column:type;"`
	IsFree        int64  `json:"isFree" gorm:"column:isFree;"`
	IsDrm         int64  `json:"isDrm" gorm:"column:isDrm;"`
}

func (Content2) TableName() string { return "contents" }

type ResponseTV struct {
	Data []struct {
		ID       interface{} `json:"id"`
		Name     string      `json:"name"`
		Type     string      `json:"type"`
		ItemType string      `json:"itemType,omitempty"`
		Content  []ContentTV `json:"content"`
	} `json:"data"`
}

type ContentTV struct {
	ContentID     int64  `json:"id" gorm:"column:contentId;"`
	Name          string `json:"name" gorm:"column:name;"`
	Description   string `json:"description" gorm:"column:description;"`
	PublishedTime string `json:"publishedTime" gorm:"column:publishedTime;"`
	CoverImage    string `json:"coverImage" gorm:"column:coverImage;"`
	CoverImageH   string `json:"coverImageH" gorm:"column:coverImageH;"`
	Type          string `json:"type" gorm:"column:type;"`
	IsFree        int64  `json:"isFree" gorm:"column:isFree;"`
	IsDrm         int64  `json:"isDrm" gorm:"column:isDrm;"`
}

func (ContentTV) TableName() string { return "contents_tv" }
