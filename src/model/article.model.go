package model

type Articles struct {
	Data []ArticlePayloads  `json:"data"`
	Meta ArticlesPagination `json:"meta"`
}

type ArticlesSlug struct {
	Data []ArticlePayload   `json:"data"`
	Meta ArticlesPagination `json:"meta"`
}

type Article struct {
	Data ArticlePayload     `json:"data"`
	Meta ArticlesPagination `json:"meta"`
}

type ArticlePayload struct {
	Id                    int64                   `json:"id"`
	Title                 string                  `json:"title"`
	Content               string                  `json:"content"`
	Description           string                  `json:"description"`
	CreatedAt             string                  `json:"createdAt"`
	UpdatedAt             string                  `json:"updatedAt"`
	PublishedAt           string                  `json:"publishedAt"`
	Locale                string                  `json:"locale"`
	Slug                  string                  `json:"slug"`
	Thumnail              ArticleThumnail         `json:"thumnail"`
	Categories            []Categories            `json:"categories"`
	UsersPermissionsUsers []UsersPermissionsUsers `json:"users_permissions_users"`
	Advisors              []Advisors              `json:"advisors"`
	Courses               []Courses               `json:"courses"`
}

type ArticlePayloads struct {
	Id          int64        `json:"id"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	Description string       `json:"description"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
	PublishedAt string       `json:"publishedAt"`
	Locale      string       `json:"locale"`
	Slug        string       `json:"slug"`
	Categories  []Categories `json:"categories"`
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
