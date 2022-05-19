package model

type ResponseArticlesAll struct {
	Error      error                     `json:"error"`
	Result     []ArticlePayloads         `json:"result"`
	Pagination ArticlesPaginationContent `json:"pagination"`
}

type ResponseArticle struct {
	Error  error          `json:"error"`
	Result ArticlePayload `json:"result"`
}
