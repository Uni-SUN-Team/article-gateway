package model

type Article struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	PublishedAt string `json:"publishedAt"`
	Locale      string `json:"locale"`
	Slug        string `json:"slug"`
}

type Articles struct {
	Data []Article          `json:"data"`
	Meta ArticlesPagination `json:"meta"`
}

type ArticlesPagination struct {
	Pagination ArticlesPaginationContent `json:"pagination"`
}

type ArticlesPaginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}
