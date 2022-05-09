package model

type Courses struct {
	Id            int64           `json:"id"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	Content       string          `json:"content"`
	Short_content string          `json:"short_content"`
	Price         float64         `json:"price"`
	CreatedAt     string          `json:"createdAt"`
	UpdatedAt     string          `json:"updatedAt"`
	PublishedAt   string          `json:"publishedAt"`
	Locale        string          `json:"locale"`
	Slug          string          `json:"slug"`
	Thumnail      ArticleThumnail `json:"thumnail"`
}
